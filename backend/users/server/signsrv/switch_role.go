package signsrv

import (
	"context"
	"database/sql"
	"time"

	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/internal/userjwt"
	"github.com/athomecomar/athome/backend/users/server"
	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/athome/pb/pbutil"
	"github.com/athomecomar/xerrors"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"
)

func (s *Server) SwitchRole(ctx context.Context, in *pbusers.SwitchRoleRequest) (*pbusers.SignResponse, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}

	db, err := server.ConnDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	auth, authCloser, err := pbutil.ConnAuth(ctx)
	if err != nil {
		return nil, err
	}
	defer authCloser()

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	return s.switchRole(ctx, db, auth, in)
}

func (s *Server) switchRole(ctx context.Context, db *sqlx.DB, c pbauth.AuthClient, in *pbusers.SwitchRoleRequest) (*pbusers.SignResponse, error) {
	oldUserId, err := pbutil.GetUserFromAccessToken(ctx, c, in.GetAccessToken())
	if err != nil {
		return nil, err
	}
	oldUser, err := server.FindUser(ctx, db, oldUserId)
	if err != nil {
		return nil, err
	}
	row := db.QueryRowxContext(ctx, `SELECT * FROM users WHERE email=$1 AND role=$2`, oldUser.Email, in.GetRole())
	user := &ent.User{}
	err = row.StructScan(user)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, status.Errorf(xerrors.NotFound, "user StructScan: %v", err)
	}
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "user StructScan: %v", err)
	}

	if user.PasswordHash != oldUser.PasswordHash {
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(in.GetPassword()))
		if err != nil {
			return nil, status.Errorf(xerrors.Unauthenticated, "bcrypt.CompareHashAndPassword: %v", err)
		}
	}

	signToken, err := userjwt.CreateSignToken(user.Id)
	if err != nil {
		return nil, status.Errorf(xerrors.Internal, "createSignToken: %v", err)
	}

	return s.sign(ctx, c, &pbusers.SignRequest{SignToken: signToken})
}

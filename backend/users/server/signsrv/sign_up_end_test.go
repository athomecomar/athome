package signsrv

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/athomecomar/athome/backend/users/ent"
	"github.com/athomecomar/athome/backend/users/ent/field"
	"github.com/athomecomar/athome/backend/users/internal/userjwt"
	"github.com/athomecomar/athome/backend/users/internal/usertest"
	"github.com/athomecomar/athome/backend/users/userconf"
	"github.com/athomecomar/athome/pb/pbusers"
	"github.com/athomecomar/storeql"
	"github.com/athomecomar/storeql/test/sqlassist"
	"github.com/athomecomar/storeql/test/sqlhelp"
	"github.com/athomecomar/xerrors"
	"github.com/google/go-cmp/cmp"
	_ "github.com/lib/pq"
	"google.golang.org/grpc/status"
)

const fooPwd = "foopassword3"

func TestServer_signUpEnd(t *testing.T) {
	type args struct {
		ctx      context.Context
		in       *pbusers.SignUpEndRequest
		previous *ent.Onboarding
	}
	tests := []struct {
		name       string
		args       args
		queryStubs []*sqlassist.QueryStubber
		execStubs  []*sqlassist.ExecStubber
		want       *pbusers.SignUpEndResponse
		wantStatus xerrors.Code
	}{
		{
			name: "oor consumer",
			args: args{
				ctx:      context.Background(),
				in:       usertest.OnboardingToSignUpEndRequest(gOnboardings.Consumers.Foo, fooPwd),
				previous: usertest.SetStage(t, gOnboardings.Consumers.Foo, field.SelectCategory),
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "INSERT INTO users", Rows: sqlmock.NewRows([]string{"id"}).AddRow(gUsers.Consumers.Foo.Id),
				},
			},
			execStubs: []*sqlassist.ExecStubber{
				{
					Expect: "DELETE FROM onboardings", Result: sqlmock.NewResult(0, 1),
				},
			},
			want:       &pbusers.SignUpEndResponse{User: usertest.UserToSignInUser(t, gUsers.Consumers.Foo)},
			wantStatus: xerrors.OutOfRange,
		},
		{
			name: "basic consumer",
			args: args{
				ctx:      context.Background(),
				in:       usertest.OnboardingToSignUpEndRequest(gOnboardings.Consumers.Foo, fooPwd),
				previous: usertest.SetStage(t, gOnboardings.Consumers.Foo, field.Shared),
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "INSERT INTO users", Rows: sqlmock.NewRows([]string{"id"}).AddRow(gUsers.Consumers.Foo.Id),
				},
			},
			execStubs: []*sqlassist.ExecStubber{
				{
					Expect: "DELETE FROM onboardings", Result: sqlmock.NewResult(0, 1),
				},
			},
			want:       &pbusers.SignUpEndResponse{User: usertest.UserToSignInUser(t, gUsers.Consumers.Foo)},
			wantStatus: xerrors.OK,
		},
		{
			name: "basic service-provider",
			args: args{
				ctx:      context.Background(),
				in:       usertest.OnboardingToSignUpEndRequest(gOnboardings.ServiceProviders.Medic.Foo, fooPwd),
				previous: usertest.SetStage(t, gOnboardings.ServiceProviders.Medic.Foo, field.Identification),
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "INSERT INTO users", Rows: sqlmock.NewRows([]string{"id"}).AddRow(gUsers.ServiceProviders.Medic.Foo.Id),
				},
				{
					Expect: "SELECT * FROM onboarding_identifications",
					Rows: sqlmock.
						NewRows(storeql.SQLColumns(gOnboardingIdentifications.ServiceProviders.Medic.Foo)).
						AddRow(storeql.
							SQLValues(gOnboardingIdentifications.ServiceProviders.Medic.Foo)...,
						),
				},
				{
					Expect: "INSERT INTO identifications", Rows: sqlmock.NewRows([]string{"id"}).AddRow(gOnboardingIdentifications.ServiceProviders.Medic.Foo.Id),
				},
			},
			execStubs: []*sqlassist.ExecStubber{
				{
					Expect: "DELETE FROM onboarding_identifications", Result: sqlmock.NewResult(0, 1),
				},
				{
					Expect: "DELETE FROM onboardings", Result: sqlmock.NewResult(0, 1),
				},
			},
			want:       &pbusers.SignUpEndResponse{User: usertest.UserToSignInUser(t, gUsers.ServiceProviders.Medic.Foo)},
			wantStatus: xerrors.OK,
		},
		{
			name: "oor service-provider",
			args: args{
				ctx:      context.Background(),
				in:       usertest.OnboardingToSignUpEndRequest(gOnboardings.ServiceProviders.Medic.Foo, fooPwd),
				previous: usertest.SetStage(t, gOnboardings.ServiceProviders.Medic.Foo, field.Shared),
			},
			queryStubs: []*sqlassist.QueryStubber{
				{
					Expect: "INSERT INTO users", Rows: sqlmock.NewRows([]string{"id"}).AddRow(gUsers.ServiceProviders.Medic.Foo.Id),
				},
			},
			execStubs: []*sqlassist.ExecStubber{
				{
					Expect: "DELETE FROM onboardings", Result: sqlmock.NewResult(0, 1),
				},
			},
			wantStatus: xerrors.OutOfRange,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			db, mock := sqlhelp.MockDB(t)
			for _, stub := range tt.queryStubs {
				stub.Stub(mock)
			}
			for _, stub := range tt.execStubs {
				stub.Stub(mock)
			}
			s := &Server{}
			got, err := s.signUpEnd(tt.args.ctx, db, tt.args.in, usertest.CopyOnboarding(t, tt.args.previous))
			if status.Code(err) != tt.wantStatus {
				t.Fatalf("Server.signUpEnd() error = %v, status: %v;  wantStatus %v", err, status.Code(err), tt.wantStatus)
			}
			if status.Code(err) != xerrors.OK {
				return
			}
			tt.want.User.SignToken = got.User.GetSignToken()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Server.signUpEnd() errored mismatch (-want +got): %s", diff)
			}

			if status.Code(err) != xerrors.OK {
				return
			}

			claims, err := userjwt.ClaimJwt(got.User.GetSignToken(), userconf.GetSIGN_JWT_SECRET)
			if err != nil {
				t.Fatalf("Server.signUpEnd() errored parsing generated sign token: %v", err)
			}
			userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
			if err != nil {
				t.Fatalf("Server.signUpEnd() errored parsing generated sign token: %v", err)
			}
			if userId != tt.want.User.Id {
				t.Errorf("Server.signUpEnd() errored mismatch user id assigned to sign token. Got %v, want %v", userId, tt.want.User.Id)
			}
			if err := claims.Valid(); err != nil {
				t.Errorf("Server.signUpEnd() given invalid sign token: %v", err)
			}

		})
	}
}

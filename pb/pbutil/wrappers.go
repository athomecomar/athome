package pbutil

import (
	"context"

	"github.com/athomecomar/athome/pb/pbagreement"
	"github.com/athomecomar/athome/pb/pbauth"
	"github.com/athomecomar/athome/pb/pbshared"
	"github.com/athomecomar/storeql"
)

func GetUserFromAccessToken(ctx context.Context, c pbauth.AuthClient, access string) (uint64, error) {
	resp, err := c.RetrieveAuthentication(ctx, &pbauth.RetrieveAuthenticationRequest{AccessToken: access})
	if err != nil {
		return 0, err
	}

	return resp.GetUserId(), nil
}

func Agree(ctx context.Context, uid uint64, token uint64) error {
	agree, agreeCloser, err := ConnAgreement(ctx)
	if err != nil {
		return err
	}
	_, err = agree.Verify(ctx, &pbagreement.VerifyRequest{AgreedUserId: uid, AgreementToken: token})
	if err != nil {
		return err
	}
	return agreeCloser()
}

func ToPbEntity(s storeql.Storable) *pbshared.Entity {
	return &pbshared.Entity{
		EntityId:    s.GetId(),
		EntityTable: s.SQLTable(),
	}
}

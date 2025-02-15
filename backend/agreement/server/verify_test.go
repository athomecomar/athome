package server

import (
	"context"
	"testing"
	"time"

	"github.com/athomecomar/athome/pb/pbagreement"
	"github.com/athomecomar/xerrors"
	"google.golang.org/grpc/status"
)

func TestServer_verify(t *testing.T) {
	type args struct {
		ctx context.Context
		in  *pbagreement.VerifyRequest
	}
	tests := []struct {
		name           string
		previousRecord *record
		args           args
		wantErr        xerrors.Code
	}{
		{
			name: "basic",
			previousRecord: &record{
				UserId:         1,
				AgreementToken: 89438,
			},
			args: args{
				ctx: context.Background(),
				in:  &pbagreement.VerifyRequest{AgreedUserId: 1, AgreementToken: 89438},
			},
		},
		{
			name: "mismatch token",
			previousRecord: &record{
				UserId:         1,
				AgreementToken: 3423,
			},
			args: args{
				ctx: context.Background(),
				in:  &pbagreement.VerifyRequest{AgreedUserId: 1, AgreementToken: 32934},
			},
			wantErr: xerrors.PermissionDenied,
		},
		{
			name: "mismatch uid",
			previousRecord: &record{
				UserId:         1,
				AgreementToken: 2342,
			},
			args: args{
				ctx: context.Background(),
				in:  &pbagreement.VerifyRequest{AgreedUserId: 2, AgreementToken: 2342},
			},
			wantErr: xerrors.PermissionDenied,
		},
		{
			name: "no records",
			args: args{
				ctx: context.Background(),
				in:  &pbagreement.VerifyRequest{AgreedUserId: 2, AgreementToken: 213212},
			},
			wantErr: xerrors.PermissionDenied,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			miniredis := connMiniredis(t)
			redis := redisCli(miniredis)
			if tt.previousRecord != nil {
				tt.previousRecord.save(t, redis, 30*time.Second)
			}
			s := &Server{Redis: redis}

			_, err := s.verify(tt.args.ctx, tt.args.in)
			if status.Code(err) != tt.wantErr {
				t.Errorf("Server.verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

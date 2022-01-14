package auth

import (
	"context"
	"coolcar/shared/auth/token"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type tokenVerifier interface {
	Verifier(token string) (string, error)
}

// interceptor i
type interceptor struct {
	verifier tokenVerifier
}

func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {

	f, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("cannot open public key file: %v", err)
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("cannot read public key: %v", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("cannot parse public key: %v ", err)
	}

	i := &interceptor{
		verifier: &token.JWTTokenVerifier{
			PublicKey: pubKey,
		},
	}

	return i.HandleReq, nil
}

// interceptor i
func (i *interceptor) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	tkn, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "")
	}
	aid, err := i.verifier.Verifier(tkn)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "token not valid: %v", err)
	}
	return handler(ctx, req)
}

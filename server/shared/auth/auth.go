package auth

import (
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
)

// interceptor i
type interceptor struct {
	publicKey *rsa.PublicKey
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
		publicKey: pubKey,
	}

	return i.HandleReq, nil
}

// interceptor i
func (i *interceptor) HandleReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	return handler(ctx, req)
}

package cos

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type Service struct {
	client *cos.Client
	secID  string
	secKey string
}

// "https://coolcar-1253590403.cos.ap-shanghai.myqcloud.com"

// "AKIDykGVUDfqkNjocCVXIl4B4h4B1atvgANK"
// "i9JDrEMHwMWQ5UdWAXeV7TaR2LRHrCH6"
func NewService(addr, secID, secKey string) (*Service, error) {
	u, err := url.Parse(addr)
	if err != nil {
		return nil, fmt.Errorf("cannot parse addr: %v", err)
	}
	b := &cos.BaseURL{
		BucketURL: u,
	}

	return &Service{
		client: cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  secID,
				SecretKey: secKey,
			},
		}),
		secID:  secID,
		secKey: secKey,
	}, nil
}

func (s *Service) SignURL(c context.Context, method, path string, timeoutSec time.Duration) (string, error) {

}

func (s *Service) Get(c context.Context, path string) (io.ReadCloser, error) {

}

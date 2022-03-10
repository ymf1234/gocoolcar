package cos

import (
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

type Service struct {
	client *cos.Client
	secID  string
	secKey string
}

func NewService() (*Service, error) {
	u, err := url.Parse("https://coolcar-1253590403.cos.ap-shanghai.myqcloud.com")
	if err != nil {
		panic(err)
	}
	b := &cos.BaseURL{
		BucketURL: u,
	}
	secID := ""
	secKey := ""

	cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secID,
			SecretKey: secKey,
		},
	})
}

package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const privteKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEA00nNFq/mHMj5q89D+dRUxMF7oovJFg4uklQpdfpvVeYTHT+c
KvmDQzv08W3UtOVx/amRV1gYpuoU4fz7gMtzSPbCEiIlYU6JfoDjNDvj6fOKJ5LJ
1BBo82fTnjovGIFr7Y4UI/eqczLYbJIlRPa9v/bcKVhgOdbypGywMfEfgHBJKdYE
a8OwQIYA+q/O6ztwfmWO5r7o37rHlTo3lISA75YEAkYU0PG54sAWsfQyzUZ86Zk9
EhxBTSxyY8IFSRTF7aehSWNy7eTMygrKG0zZmwtFw9IZaMgfe7pSKlpAufj1hIaF
JkAz7A3E3jf8W9yBLfyMUKbEAOr08SiGQwOSiQIDAQABAoIBADGF7dUwe24pySN1
HEi1dPz9IL5zpzlNL6tKoMEvZuMqXzaOz0rfWt3qcroX9TBFS1eaZB+l3XlaCnUB
t+1zQl1KwShA6qQJJX8KNLbshroyRCKMbFQRWJgcsOQ4XQuQYjy2/Y+EoGYJ9+JS
yvuEHeudr/Lc7QkBkbPqac3651hSC85DVjTzSnNs3Szb87Inh+fbIhvH+1I2wIsu
Uz+CjrvTvSOW2QNHEY7buG1ueBl9GA3dYGlIV4jCQLWubAvejX1jRHnhbpBpo3iG
4MvgpXDYT/g6V7eN7GTNYKjTJGS5l5h5zn88JARnR4aE84KahTkj2xJHd1d2FtRw
Z5XfdYECgYEA/kqFbeNHtNsp1FEG6yccDQBQ9CrI/qK4+efZzYtghP0PoAam+lvC
VXzHxoCc+tsGo3yw5i59jSasfGdGIoI5om7pokmmuFhjnhFip89HcOW/F0gFRD0U
bUC9ZsxwBCPSJNlHLUPrdRf3+zSI/IoIBbj/a5j9jQ8IH+EiPuFPZzkCgYEA1LVM
ae5nF+YYr33+aIF03Yt2WUGvCNmHaln79q3mVYj9mZt2GJKMbBUWVwGlQS4CS1Xa
FfL01aE+0vCMltw2F+nbHPYSnae3ex0T56ekIHAjuHt2LC71LRf6Nj3v4fYYejTu
iEe/RIK8RzD7xvKLRfMQBeZ3NSEovVJu1pndtdECgYBxAiM8C6aBXGjBf2Touk/A
A2T2rdD3QM+d2Ai2TPuO/nukLXClkIPTxAlwURXyccYIf9CJ6x1GzmoXOzoy+dat
NkdEJSEficBzrdw/IWHpbWC38vzML6zVRNEnTzF01Yp9b8Yu+qE4P3eoPtemC523
FVkFs07zEjwVFtNGSJEmGQKBgEG/8OjlupRWmdbVWWiwa5F35JNejF5t7z/epNXx
lztlZw6uUWjzE9jtA8TNS2nEcQ0ccpdKTIgkmajWln6jf00RCpv/LG3zQ/IYYOjA
lyk3wLLr5vFtangP61YM+zObwKE1723Mz8C7sLtI1ur+WjK7GX7ZaIvRvigXbQ+E
7FWhAoGABDwuEKVWmbYxlck8P/VkpplXln9x3avSoHSkYARz7UwIybdAa8eGwlnb
5xS1wtDeFmtaSlVE5jlKrtTMCOCJ1E7YfzRnktbyQeQkvApEQ8QEAWXSfgPM92Ai
M1diwVt9ZEbia0POfslv7bG6EgaiYFO5l5y+03zbyVzRY82Eao0=
-----END RSA PRIVATE KEY-----`

func TestGenerateToken(t *testing.T) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privteKey))
	if err != nil {
		t.Fatalf("cannot parse private key: %v", err)
	}
	g := newJWTTokenGen("coolcar/auth", key)
	g.nowFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}
	tkn, err := g.GeneratrToken("5f7c3168e2283aa722e351a3", 2*time.Hour)
	if err != nil {
		t.Errorf("33 %v", err)
	}

	want := "aaa"

	if want != tkn {
		t.Errorf("1: %q -- 2:%q ", want, tkn)
	}
}

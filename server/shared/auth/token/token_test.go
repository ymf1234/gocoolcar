package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const publicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDEnzRVRX8SyyDgXEuAxaSa4Eqw
xj00Fj4ZXehj97f9rSKvmKg0ckHkg6PlVHuXczbxEqjCu+T5uE4qW0CLsG1XgmN2
H8L93RxPr+OBgr02VnCmNwacboIe6jn1f+YPjnobc7cG5dXvm8O2kzWdTQ/6248K
ZKaZ7B6UM35VAZgj1QIDAQAB
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		t.Fatalf("cannot parse public key: %v", err)
	}

	v := &JWTTokenVerifier{
		PublicKey: pubKey,
	}

	tkn := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NDE2Mzc0ODUsImlhdCI6MTY0MTYzMDI4NSwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjFkNzA2ZjczMjdjYjhiOWQ2NGNlNDVkIn0.Y5LdfI5HZ5XT-1dZmARQ5GDeSHURLzDzTN8De4CU9sUXQv241QjGU-U03RYp5UM-areEPxqnWR_zTw_-BRENkJgt6zqD6bB3Z1WipKe1mmd3G8XMPYrg8LfCIpYzX29e0azHvIa-I5F1XmPNrnq1ehXNJGowkOMbVVi2u1sFW3rKeUdeXgvkrNy_mbKu2ZKxAZhV35weXbE99ZAIyzI99oDAMNgGmlZPO0-ut6wled_AOKIgscsZXqRDkqJO8rgHlj00de8Q--tRCciUNq6bjugzi4VoqxyPuCmGsbhi9TG9O9xMWrFkuNT0kT5w3SQv1_P6pGpToKFe9O10gSHqtg"

	jwt.TimeFunc = func() time.Time {
		return time.Unix(1111111111, 0)
	}
	accountID, err := v.Verifier(tkn)

	if err != nil {
		t.Errorf("verification failed: %v", err)
	}

	want := "oQ5lw5IDKTU-PJ7qz0RLlNAGgr58"
	if accountID != want {
		t.Errorf("want: %v  -- got : %v", want, accountID)
	}
}

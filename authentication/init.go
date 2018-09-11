package authentication

import (
	"crypto/rsa"
	joseRsa "github.com/dvsekhvalnov/jose2go/keys/rsa"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	errorMessages "github.com/quinlanmorake/verisart-go/types/core/errorMessages"

	config "github.com/quinlanmorake/verisart-go/config"
)

var tokenIssuer coreTypes.String

// NOTE : May well eventually move the below into other files or so on.
// Added here as use of ioutil has already been shown
func Init(config config.Config) coreTypes.Result {
	tokenIssuer = config.Jwt.Issuer

	if _privateKey, readPrivateKeyError := joseRsa.ReadPrivate([]byte(privateKeyString)); readPrivateKeyError != nil {
		return coreTypes.Result{
			Code:    errorCodes.JWT_ERROR_PARSING_THE_PRIVATE_KEY,
			Message: errorMessages.ErrorMessage(readPrivateKeyError.Error()),
		}
	} else {
		privateKey = _privateKey
	}

	if _publicKey, readPublicKeyError := joseRsa.ReadPublic([]byte(publicKeyString)); readPublicKeyError != nil {
		return coreTypes.Result{
			Code:    errorCodes.JWT_ERROR_PARSING_THE_PUBLIC_KEY,
			Message: errorMessages.ErrorMessage(readPublicKeyError.Error()),
		}
	} else {
		publicKey = _publicKey
	}

	return coreTypes.NewSuccessResult()
}

const privateKeyString = `-----BEGIN RSA PRIVATE KEY-----
MIIEpQIBAAKCAQEAy2CSgMl19a9ayokp5aU1qAYl11y3WPRHy56npFkkwj6WQp8O
0lfOj4YZ2O9aEZyWppiJHRjPfwAQHr9v4ADuqcbbyaYy2psAGlHN1TnyB+FsatwA
36QQMdwSAsDIeNlfV1abeQI4vpcPa2QuL8qvNuTlRi+m0UZjgwHMzMKJpNDm1OKF
XPSiCCLhPPdPyR9WMj4gBDr9BGE4DwmYlTB3RqgFwywXQ56D6TZO3POzL+JgBEJY
5JZTd031Qdd9YV3iess4iroyp04nthx0hMoHvrO2WQmTDi+bXN7OQT6Eh5cpCkUz
NlhoDl1tsvLOu2Hsjj678ubxBMhXjRxGBEGV/wIDAQABAoIBAF0maVOlcmrKGzqC
4na1WcMQtcpTRALxN+USUNRcpBJ32hi/g2liIFDqafJLhqRkZTf+y2ZRU9Bmdfuv
UQ13P9jxJP/SKJ6pub1UVINiG/HyFNdI1vmXyncbxgAOzoNThHCbUXB2TQk2ZTCZ
3PG00tXAtyaTcsONisyZSXVLcQHUfu6G181xzr269nzmegz8jnHXhccYpO0UIwlu
zOsmu2oNgXCtlRQvkTkxalOkAAzZ6cx+qFlpLHny8oqxVaMzUL+uRR7cb5s44Sps
JvLtRtM04DWp1Srsvr1lpFsugrsPZ1buctLGpziI7RpXmJAvHDheJT12dVrKpJlh
aPMicJECgYEA+G8g/NmrUU9kXuLGiqGXFsqRj22nKnhgRXqI2Q+p9iKv3plDzJIU
IjiTCIYFmw+Tmu3i57EdPD1ZLg/6qOgPN9KjPevupEvQxjRHiaF7EYKODP8D5Gs9
gn4vy026kx4FIFnGPv1H3cKg4Lll5Zuh0w3nE9u6YxrkYpANrJgWgdkCgYEA0ZIq
crP4a8CeVnudL1JfzUXvhEsYuvwpVmsbYkendCl2Xt7L9CGXTH18pz1rJEoDKu5n
hR+A2kWrs6RZqtd4isB0JhJnGnmbuNPw/6vUIMMwoClvWgtcYdansjjRsrbqDZ21
nYDRwOBr0IBNpaPO+EmLLj0YJa9IWeXD/2kEl5cCgYEA0jkvf+ctiiuVcpmRikP8
10No8ybF2zFrvb9Xx93yY56slb+52IUQVnjrKr5GRhVjQSnl5UtBwvTi3xCUepOM
NR+gFUjtcsfZDa+1jWhA/OsuCx9MiuCYdzESfoXyyIURr3NoR2sKgkQs4Jzh9e7B
fBf6nxpDWZOrCSB/Abrc6NkCgYEAgbkfa2b5lAFkQZY84wjzYkYzD8nIYh4qnGKq
Tbia7+2mZu00hEayt9dJbOA4zPq/YTi3fZDVmsbblNRb6MN9yy17+AEWLy4tdUhD
+o7rPOh11f4v/iXgJnPP1CMVsrFEye7gd8FAlUqVkjeJWMnAGLK1Y4bPxqvsdjEi
lJv2CrsCgYEAnT4c40rr0SHRROzu6OZC7wD/YnpUh69uHZm99+fXtp3lxNPL5Rq0
b8HEc5I412pMTLrymVNzhVRX1GVO13p0rLp0ZF8ey1cwDrM+fnZ8gL+fsWM/9tZO
GFzkS+A0nwWtV3OgIenro84QX7IsBwRleFzVrGMbH6e1kZf3mYrCuNk=
-----END RSA PRIVATE KEY-----`

var privateKey *rsa.PrivateKey

const publicKeyString = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAy2CSgMl19a9ayokp5aU1
qAYl11y3WPRHy56npFkkwj6WQp8O0lfOj4YZ2O9aEZyWppiJHRjPfwAQHr9v4ADu
qcbbyaYy2psAGlHN1TnyB+FsatwA36QQMdwSAsDIeNlfV1abeQI4vpcPa2QuL8qv
NuTlRi+m0UZjgwHMzMKJpNDm1OKFXPSiCCLhPPdPyR9WMj4gBDr9BGE4DwmYlTB3
RqgFwywXQ56D6TZO3POzL+JgBEJY5JZTd031Qdd9YV3iess4iroyp04nthx0hMoH
vrO2WQmTDi+bXN7OQT6Eh5cpCkUzNlhoDl1tsvLOu2Hsjj678ubxBMhXjRxGBEGV
/wIDAQAB
-----END PUBLIC KEY-----`

var publicKey *rsa.PublicKey

package authentication

import (
	"time"
	
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	businessTypes "github.com/quinlanmorake/verisart-go/types"

	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"	
)

type TokenData struct {
	Algorithm coreTypes.String `json:"alg"`

	ExpiresAt    int64  `json:"exp"`
	
	Issuer    coreTypes.String `json:"iss"`
	IssueTime int64  `json:"iat"`   // timestamp - seconds since epoch

	Subject   coreTypes.String `json:"sub"`

	UniqueId  coreTypes.String `json:"jti"`   // unique id for this token, only care if we have one time use tokens
	User  businessTypes.User `json:"user"`
}

func (t TokenData) IsValid() coreTypes.Result {
	if t.Issuer.ToLowercaseString() != tokenIssuer.ToLowercaseString() { 
		return coreTypes.NewResultFromErrorCode(errorCodes.JWT_TOKEN_HAS_INVALID_ISSUER)
	}

	// Check that it the token has not expired
	currentTimestamp := time.Now().UTC().Unix()
	expiryTime := int64(t.ExpiresAt)

	if (currentTimestamp > expiryTime) {
		return coreTypes.NewResultFromErrorCode(errorCodes.JWT_TOKEN_HAS_EXPIRED)
	}
	
	return coreTypes.NewSuccessResult()	
}
package authentication

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	jose "github.com/dvsekhvalnov/jose2go"

	businessTypes "github.com/quinlanmorake/verisart-go/types"
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"

	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	errorMessages "github.com/quinlanmorake/verisart-go/types/core/errorMessages"
)

func GenerateToken(user businessTypes.User, tokenLifetime time.Duration) (coreTypes.String, coreTypes.Result) {
	currentTimestamp := time.Now()

	rand.Seed(time.Now().UnixNano())
	randomId := fmt.Sprintf("%d%05d", currentTimestamp, rand.Intn(65535))

	tokenData := TokenData{
		ExpiresAt: currentTimestamp.Add(tokenLifetime).UTC().Unix(),
		Issuer:    tokenIssuer, // Global variable defined in init.go
		IssueTime: currentTimestamp.UTC().Unix(),
		UniqueId:  coreTypes.String(randomId),
		User:      user,
	}

	payloadByteArray, _ := json.Marshal(tokenData)
	payload := string(payloadByteArray)

	if token, signingError := jose.Sign(payload, jose.RS256, privateKey); signingError != nil {
		return coreTypes.String(""), coreTypes.Result{
			Code:    errorCodes.JWT_TOKEN_SIGNING_ERROR,
			Message: errorMessages.ErrorMessage(signingError.Error()),
		}
	} else {
		return coreTypes.String(token), coreTypes.NewSuccessResult()
	}
}

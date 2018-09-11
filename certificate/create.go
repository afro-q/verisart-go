package certificate

import (
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"	
	businessTypes "github.com/quinlanmorake/verisart-go/types"
	
	database "github.com/quinlanmorake/verisart-go/database"

	user "github.com/quinlanmorake/verisart-go/user"
)

func Create(userId coreTypes.UserId, certificate businessTypes.Certificate) (businessTypes.Certificate, coreTypes.Result) {
	// Make sure user exists
	if _, getUserByIdResult := user.GetUserById(userId); getUserByIdResult.IsNotOk() {
		return businessTypes.NewEmptyCertificate(), getUserByIdResult
	}

	newCertificate := businessTypes.Certificate{
		CreatedAt: coreTypes.NewDateStringTimestamp(),
		OwnerId: userId,
		Note: certificate.Note,
		Title: certificate.Title,
		Year: certificate.Year,		
	}
	
	return newCertificate, database.Add(&newCertificate).Result
}
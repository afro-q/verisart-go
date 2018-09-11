package http

import (
	"net/http"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	
	httpHelpers "github.com/quinlanmorake/verisart-go/helpers/http"
	parameterNames "github.com/quinlanmorake/verisart-go/helpers/http/parameterNames"
	
	certificate "github.com/quinlanmorake/verisart-go/certificate"	
)

func LoadCertificatesForUser(w http.ResponseWriter, r *http.Request) {
	response := loadCertificatesForUserResponse{}
	defer func() {
		httpHelpers.WriteResponse(w, response)
	}()

	if userIdAsRequestParameter, getUserIdResult := httpHelpers.GetParameterFromRequest(r, parameterNames.USER_ID, errorCodes.USER_ID_IS_INVALID); getUserIdResult.IsNotOk() {
		response.Error = getUserIdResult		
	} else {
		response.Certificates, response.Error = certificate.LoadAllCertificatesForUser(coreTypes.UserId(userIdAsRequestParameter))
	}	
}

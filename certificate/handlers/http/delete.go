package http

import (
	"net/http"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"	
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	
	httpHelpers "github.com/quinlanmorake/verisart-go/helpers/http"
	parameterNames "github.com/quinlanmorake/verisart-go/helpers/http/parameterNames"
	
	authentication "github.com/quinlanmorake/verisart-go/authentication"
	
	certificate "github.com/quinlanmorake/verisart-go/certificate"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	response := deleteResponse{}
	defer func() {
		httpHelpers.WriteResponse(w, response)
	}()

	// The middleware assures us that we have a token if we've come this far
	user, _ := authentication.GetUserFromTokenInHeaders(r)

	if certificateIdAsRequestParameter, getCertificateIdResult := httpHelpers.GetParameterFromRequest(r, parameterNames.CERTIFICATE_ID, errorCodes.CERTIFICATE_ID_IS_INVALID); getCertificateIdResult.IsNotOk() {
		response.Error = getCertificateIdResult
	} else {
		response.Error = certificate.Delete(user.Id, coreTypes.String(certificateIdAsRequestParameter))
	}
}
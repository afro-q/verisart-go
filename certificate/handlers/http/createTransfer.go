package http

import (
	"net/http"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"	
	errorCodes "github.com/quinlanmorake/verisart-go/types/core/errorCodes"
	businessTypes "github.com/quinlanmorake/verisart-go/types"
	
	httpHelpers "github.com/quinlanmorake/verisart-go/helpers/http"
	parameterNames "github.com/quinlanmorake/verisart-go/helpers/http/parameterNames"
	
	authentication "github.com/quinlanmorake/verisart-go/authentication"
	
	certificate "github.com/quinlanmorake/verisart-go/certificate"
)

func CreateTransfer(w http.ResponseWriter, r *http.Request) {
	response := createTransferResponse{}
	defer func() {
		httpHelpers.WriteResponse(w, response)
	}()

	request := createTransferRequest{}
	if parseRequestResult := httpHelpers.ParseJsonBody(r.Body, &request); parseRequestResult.IsNotOk() {
		response.Error = parseRequestResult
		return
	}
	
	// The middleware assures us that we have a token if we've come this far
	user, _ := authentication.GetUserFromTokenInHeaders(r)

	if certificateIdAsRequestParameter, getCertificateIdResult := httpHelpers.GetParameterFromRequest(r, parameterNames.CERTIFICATE_ID, errorCodes.CERTIFICATE_ID_IS_INVALID); getCertificateIdResult.IsNotOk() {
		response.Error = getCertificateIdResult
	} else {
		var transfer businessTypes.Transfer
		
		transfer, response.Error = certificate.CreateTransfer(user, coreTypes.String(certificateIdAsRequestParameter), request.To)
		
		response.Transfer = Transfer{
			To: transfer.To,
			Status: transfer.Status,
		}
	}
}

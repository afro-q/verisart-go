package http

import (
	"net/http"

	httpHelpers "github.com/quinlanmorake/verisart-go/helpers/http"

	authentication "github.com/quinlanmorake/verisart-go/authentication"
	certificate "github.com/quinlanmorake/verisart-go/certificate"
)

func Create(w http.ResponseWriter, r *http.Request) {
	response := createResponse{}
	defer func() {
		httpHelpers.WriteResponse(w, response)
	}()

	request := createRequest{}
	if parseRequestResult := httpHelpers.ParseJsonBody(r.Body, &request); parseRequestResult.IsNotOk() {
		response.Error = parseRequestResult
		return
	}

	// The middleware assures us that we have a token if we've come this far
	user, _ := authentication.GetUserFromTokenInHeaders(r)
	
	response.Certificate, response.Error = certificate.Create(user.Id, request.ToCertificate())
}

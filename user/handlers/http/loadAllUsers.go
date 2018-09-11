package http

import (
	"net/http"

	httpHelpers "github.com/quinlanmorake/verisart-go/helpers/http"

	user "github.com/quinlanmorake/verisart-go/user"
)

/*
 This just returns all users in the database
 NOTE: No filters or sorting or pagination functions added
*/
func LoadAllUsers(w http.ResponseWriter, r *http.Request) {
	response := loadAllUsersResponse{}
	defer func() {
		httpHelpers.WriteResponse(w, response)
	}()

	response.Users, response.Error = user.LoadAllUsers()
}

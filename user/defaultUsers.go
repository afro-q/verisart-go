package user

import (
	businessTypes "github.com/quinlanmorake/verisart-go/types"
	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

var defaultUsers []businessTypes.User = []businessTypes.User{
	businessTypes.User{
		Email: coreTypes.EmailAddress("john.s@mailinator.com"),
		Name:  coreTypes.String("John Smith"),
	},
	businessTypes.User{
		Email: coreTypes.EmailAddress("graham.b@mailinator.com"),
		Name:  coreTypes.String("Graham Bell"),
	},
	businessTypes.User{
		Email: coreTypes.EmailAddress("clerk_m@mailinator.com"),
		Name:  coreTypes.String("Clerk Maxwell"),
	},
	businessTypes.User{
		Email: coreTypes.EmailAddress("conan_d@mailinator.com"),
		Name:  coreTypes.String("Conan Doyle"),
	},
}

package config

import (
	"fmt"

	coreTypes "github.com/quinlanmorake/verisart-go/types/core"
)

type server struct {
	Address coreTypes.String `json:"address"`
	Port    int    `json:"port"`
}

func (s server) GetListenAddress() string {
	return fmt.Sprintf("%v:%v", s.Address, s.Port)
}

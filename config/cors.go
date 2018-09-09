package config

type cors struct {
	AllowMethods string `json:"allowMethods"`
	AllowOrigin  string `json:"allowOrigin"`
	AllowHeaders string `json:"allowHeaders"`
}

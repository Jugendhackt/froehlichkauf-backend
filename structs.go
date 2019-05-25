package main

type datapoint struct {
}

type barcode struct {
	Code string `json:"code"`
}

type securitySetting struct {
	Header string `json:"header"`
	Option string `json:"option"`
}

type serverConfig struct {
	Port           string `json:"port"`
	Cert           string `json:"cert"`
	Key            string `json:"key"`
	SecurityConfig string `json:"securityConfig"`
}

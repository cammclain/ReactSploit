package models

// MsfAuthResponse represents the response structure for authentication requests to msfrpc.
type MsfAuthResponse struct {
	Token string `json:"token"`
	Error bool   `json:"error"`
	Info  string `json:"info,omitempty"`
}

// MsfGenericResponse can be used to decode general responses from msfrpc that do not carry specific structured data.
type MsfGenericResponse struct {
	Result string `json:"result"`
	Error  bool   `json:"error"`
	Info   string `json:"info,omitempty"`
}

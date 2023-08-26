package web

type WebResponse struct {
	Error   bool   `json:"error"`
	ReffID  string `json:"reff_id"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

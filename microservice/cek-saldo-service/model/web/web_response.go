package web

type WebResponse struct {
	Error   bool   `json:"error"`
	ReffID  string `json:"reff_id,omitempty"`
	Message any    `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

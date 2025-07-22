package response

type ErrorResponse struct {
	Error   string `json:"error" example:"Invalid request"`
	Message string `json:"message,omitempty" example:"Detailed error message"`
}

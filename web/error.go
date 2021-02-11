package web

type HTTPError struct {
	Message string      `json:"message" example:"invalid format"`
	Error   string      `json:"error,omitempty" example:"<internal error message>"`
	Data    interface{} `json:"data,omitempty" swaggerignore:"true"`
}

const (
	APIErrorInvalidRequestFormat = "Invalid Request Format"
	APIErrorInternalDatabase     = "Internal Database Error"
	APIErrorNotFound             = "Not found"
)

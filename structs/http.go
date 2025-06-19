package structs

type HttpResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error interface{} `json:"error,omitempty"`
}
type HttpErrorResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

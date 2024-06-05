package utils

type HttpErrorResponse struct {
	Method        string                 `json:"method"`
	Path          string                 `json:"path"`
	Status        int                    `json:"status"`
	CorrelationId string                 `json:"correlationId"`
	Error         HttpErrorResponseError `json:"error"`
}

type HttpErrorResponseError struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Payload map[string]interface{} `json:"payload"`
}

package data

type MiddlewareResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"message"`
	Data    []interface{} `json:"data"`
}

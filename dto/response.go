package dto

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseWithData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

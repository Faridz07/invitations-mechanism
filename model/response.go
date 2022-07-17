package model

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseWithArrayData struct {
	Status  string        `json:"status"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data"`
}

type ResponseWithSingleData struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

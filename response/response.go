package response

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Resp    interface{} `json:"respone"`
}

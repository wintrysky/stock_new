package global

type Response struct {
	Success    bool         `json:"success"`
	Data    interface{} `json:"data"`
	Total	int64 `json:"total"`
	Message string      `json:"message"`
}

const LogicY = "Y"
const LogicN = "N"
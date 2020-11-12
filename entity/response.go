// 统一响应结构体

package entity

const (
	SUCCESS = 200
	FAILURE = 400
	ERROR   = 500
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

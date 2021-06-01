package tool

import (
	"encoding/json"
	"io"
)

type JsonParse struct {

}

//解析请求，取出POST请求中的数据，赋值给对应结构体
func Decode(io io.ReadCloser, v interface{}) error{
	return json.NewDecoder(io).Decode(v)
}


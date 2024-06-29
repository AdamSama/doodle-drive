package util

import (
	"encoding/json"
	"fmt"
	"log"
)

// RespMsg: interface for sending response msg
type RespMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// NewRespMsg: generate response object
func NewRespMsg(code int, msg string, data interface{}) *RespMsg {
	return &RespMsg{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// JSONBytes: convert response object to binary format
func (resp *RespMsg) JSONBytes() []byte {
	r, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	return r
}

// JSONString: convert object to json
func (resp *RespMsg) JSONString() string {
	r, err := json.Marshal(resp)
	if err != nil {
		log.Println(err)
	}
	return string(r)
}

// GenSimpleRespStream: only contain simple response (code, msg)([]byte)
func GenSimpleRespStream(code int, msg string) []byte {
	return []byte(fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg))
}

// GenSimpleRespString: only contain simple response (code, msg)(string)
func GenSimpleRespString(code int, msg string) string {
	return fmt.Sprintf(`{"code":%d,"msg":"%s"}`, code, msg)
}

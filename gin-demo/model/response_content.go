package model

type ResponseContent struct {
    Code int `json:"code"`
    Msg string `json:"msg"`
}

func NewSuccessResponse() ResponseContent {
    return ResponseContent{
        Code: 0,
        Msg:  "",
    }
}
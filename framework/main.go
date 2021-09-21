package main

import (
	"fmt"

	web "example.com/framework/pkg"
)

type signUpReq struct {
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	ConfirmPassword string `json:"ConfirmPassword"`
}

func SignUp(ctx *web.Context) {
	req := &signUpReq{}

	err := ctx.ReadJSON(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}

	resp := &commonResponse{
		Data: 123,
	}

	err = ctx.OkJson(resp)
	if err != nil {
		fmt.Printf("err: %v", err)
	}
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func main() {
	server := web.NewSdkHttpServer("my-test-server")
	server.Route("POST", "/signup", SignUp)
	server.Start(":8080")
}

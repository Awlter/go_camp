package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	web "example.com/framework/pkg"
)

type signUpReq struct {
	Email           string `json:"Email"`
	Password        string `json:"Password"`
	ConfirmPassword string `json:"ConfirmPassword"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	req := &signUpReq{}

	ctx := &web.Context{
		W: w,
		R: r,
	}

	err := ctx.ReadJSON(req)
	if err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}

	resp := &commonResponse{
		Data: 123,
	}
	respJSON, err := json.Marshal(resp)

	if err != nil {
		fmt.Fprintf(w, "deserialized failed: %v", err)
		return
	}

	fmt.Fprintf(w, "%s", string(respJSON))
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func main() {
	server := web.NewSdkHttpServer("my-test-server")
	server.Route("/signup", SignUp)
	server.Start(":8080")
}

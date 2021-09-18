package main

import (
	"encoding/json"
	"fmt"
	"io"
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

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "read body failed: %v", err)
		return
	}

	err = json.Unmarshal(body, req)
	if err != nil {
		fmt.Fprintf(w, "deserialized failed: %v", err)
		return
	}

	ctx := &commonResponse{
		Data: 123,
	}
	respJSON, err := json.Marshal(ctx)

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

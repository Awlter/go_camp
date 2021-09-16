package main

import (
	"fmt"
	"net/http"

	web "example.com/framework/pkg"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是主页")
}

func main() {
	server := web.NewSdkHttpServer("my-test-server")
	server.Route("/", home)
	server.Start(":8080")
}

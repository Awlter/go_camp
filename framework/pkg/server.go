package web

import (
	"fmt"
	"net/http"
)

type Server interface {
	Route(method string, pattern string, handlerFunc func(ctx *Context))
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler *MapHandler
}

func (s *sdkHttpServer) Route(method string, pattern string, handlerFunc func(ctx *Context)) {
	fmt.Println(s.handler.handlers)
	key := s.handler.key(method, pattern)
	fmt.Println(key)
	// s.handler.handlers[key] = handlerFunc
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, s.handler)
}

func NewSdkHttpServer(name string) Server {
	return &sdkHttpServer{Name: name}
}

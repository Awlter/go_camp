package web

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

//interface{} top level hierarchy
func (c *Context) ReadJSON(req interface{}) error {

	r := c.R
	body, err := io.ReadAll(r.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(body, req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Context) WriteJSON(code int, resp interface{}) error {
	respJSON, err := json.Marshal(resp)

	if err != nil {
		return err
	}

	c.W.WriteHeader(code)

	_, err = c.W.Write(respJSON)

	return err
}

func (c *Context) OkJson(resp interface{}) error {
	return c.WriteJSON(http.StatusOK, resp)
}

func (c *Context) BadRequestJson(err interface{}) error {
	return c.WriteJSON(http.StatusBadRequest, err)
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		W: w,
		R: r,
	}
}

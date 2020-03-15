package controllers

import (
	"github.com/savsgio/atreugo/v10"
)

type baseResponse struct {
	Code    int    `json:"code"`
	Data    string `json:"data,omitempty"`
	Message string `json:"msg,omitempty"`
}

//GET ...
func GET(ctx *atreugo.RequestCtx) error {
	response := baseResponse{200, "data", "message"}
	return ctx.JSONResponse(response)
}

//POST ...
func POST(ctx *atreugo.RequestCtx) error {
	response := baseResponse{200, "data", "message"}
	return ctx.JSONResponse(response)
}

//PUT ...
func PUT(ctx *atreugo.RequestCtx) error {
	response := baseResponse{200, "data", "message"}
	return ctx.JSONResponse(response)
}

//DELETE ...
func DELETE(ctx *atreugo.RequestCtx) error {
	response := baseResponse{200, "data", "message"}
	return ctx.JSONResponse(response)
}

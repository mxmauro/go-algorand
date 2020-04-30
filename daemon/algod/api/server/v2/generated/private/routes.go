// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (POST /v2/register-participation-keys/{address})
	RegisterParticipationKeys(ctx echo.Context, address string, params RegisterParticipationKeysParams) error

	// (POST /v2/shutdown)
	ShutdownNode(ctx echo.Context, params ShutdownNodeParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// RegisterParticipationKeys converts echo context to params.
func (w *ServerInterfaceWrapper) RegisterParticipationKeys(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params RegisterParticipationKeysParams
	// ------------- Optional query parameter "fee" -------------
	if paramValue := ctx.QueryParam("fee"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "fee", ctx.QueryParams(), &params.Fee)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter fee: %s", err))
	}

	// ------------- Optional query parameter "key-dilution" -------------
	if paramValue := ctx.QueryParam("key-dilution"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "key-dilution", ctx.QueryParams(), &params.KeyDilution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter key-dilution: %s", err))
	}

	// ------------- Optional query parameter "round-last-valid" -------------
	if paramValue := ctx.QueryParam("round-last-valid"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "round-last-valid", ctx.QueryParams(), &params.RoundLastValid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round-last-valid: %s", err))
	}

	// ------------- Optional query parameter "no-wait" -------------
	if paramValue := ctx.QueryParam("no-wait"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "no-wait", ctx.QueryParams(), &params.NoWait)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter no-wait: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RegisterParticipationKeys(ctx, address, params)
	return err
}

// ShutdownNode converts echo context to params.
func (w *ServerInterfaceWrapper) ShutdownNode(ctx echo.Context) error {
	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params ShutdownNodeParams
	// ------------- Optional query parameter "timeout" -------------
	if paramValue := ctx.QueryParam("timeout"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "timeout", ctx.QueryParams(), &params.Timeout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter timeout: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ShutdownNode(ctx, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST("/v2/register-participation-keys/:address", wrapper.RegisterParticipationKeys)
	router.POST("/v2/shutdown", wrapper.ShutdownNode)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
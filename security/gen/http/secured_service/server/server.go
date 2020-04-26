// Code generated by goa v3.1.2, DO NOT EDIT.
//
// secured_service HTTP server
//
// Command:
// $ goa gen goa.design/examples/security/design -o
// $(GOPATH)/src/goa.design/examples/security

package server

import (
	"context"
	"net/http"

	securedservice "goa.design/examples/security/gen/secured_service"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the secured_service service endpoint HTTP handlers.
type Server struct {
	Mounts           []*MountPoint
	Signin           http.Handler
	Secure           http.Handler
	DoublySecure     http.Handler
	AlsoDoublySecure http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the secured_service service endpoints
// using the provided encoder and decoder. The handlers are mounted on the
// given mux using the HTTP verb and path defined in the design. errhandler is
// called whenever a response fails to be encoded. formatter is used to format
// errors returned by the service methods prior to encoding. Both errhandler
// and formatter are optional and can be nil.
func New(
	e *securedservice.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Signin", "POST", "/signin"},
			{"Secure", "GET", "/secure"},
			{"DoublySecure", "PUT", "/secure"},
			{"AlsoDoublySecure", "POST", "/secure"},
		},
		Signin:           NewSigninHandler(e.Signin, mux, decoder, encoder, errhandler, formatter),
		Secure:           NewSecureHandler(e.Secure, mux, decoder, encoder, errhandler, formatter),
		DoublySecure:     NewDoublySecureHandler(e.DoublySecure, mux, decoder, encoder, errhandler, formatter),
		AlsoDoublySecure: NewAlsoDoublySecureHandler(e.AlsoDoublySecure, mux, decoder, encoder, errhandler, formatter),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "secured_service" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Signin = m(s.Signin)
	s.Secure = m(s.Secure)
	s.DoublySecure = m(s.DoublySecure)
	s.AlsoDoublySecure = m(s.AlsoDoublySecure)
}

// Mount configures the mux to serve the secured_service endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountSigninHandler(mux, h.Signin)
	MountSecureHandler(mux, h.Secure)
	MountDoublySecureHandler(mux, h.DoublySecure)
	MountAlsoDoublySecureHandler(mux, h.AlsoDoublySecure)
}

// MountSigninHandler configures the mux to serve the "secured_service" service
// "signin" endpoint.
func MountSigninHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/signin", f)
}

// NewSigninHandler creates a HTTP handler which loads the HTTP request and
// calls the "secured_service" service "signin" endpoint.
func NewSigninHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeSigninRequest(mux, decoder)
		encodeResponse = EncodeSigninResponse(encoder)
		encodeError    = EncodeSigninError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "signin")
		ctx = context.WithValue(ctx, goa.ServiceKey, "secured_service")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountSecureHandler configures the mux to serve the "secured_service" service
// "secure" endpoint.
func MountSecureHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/secure", f)
}

// NewSecureHandler creates a HTTP handler which loads the HTTP request and
// calls the "secured_service" service "secure" endpoint.
func NewSecureHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeSecureRequest(mux, decoder)
		encodeResponse = EncodeSecureResponse(encoder)
		encodeError    = EncodeSecureError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "secure")
		ctx = context.WithValue(ctx, goa.ServiceKey, "secured_service")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDoublySecureHandler configures the mux to serve the "secured_service"
// service "doubly_secure" endpoint.
func MountDoublySecureHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/secure", f)
}

// NewDoublySecureHandler creates a HTTP handler which loads the HTTP request
// and calls the "secured_service" service "doubly_secure" endpoint.
func NewDoublySecureHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDoublySecureRequest(mux, decoder)
		encodeResponse = EncodeDoublySecureResponse(encoder)
		encodeError    = EncodeDoublySecureError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "doubly_secure")
		ctx = context.WithValue(ctx, goa.ServiceKey, "secured_service")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountAlsoDoublySecureHandler configures the mux to serve the
// "secured_service" service "also_doubly_secure" endpoint.
func MountAlsoDoublySecureHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/secure", f)
}

// NewAlsoDoublySecureHandler creates a HTTP handler which loads the HTTP
// request and calls the "secured_service" service "also_doubly_secure"
// endpoint.
func NewAlsoDoublySecureHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeAlsoDoublySecureRequest(mux, decoder)
		encodeResponse = EncodeAlsoDoublySecureResponse(encoder)
		encodeError    = EncodeAlsoDoublySecureError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "also_doubly_secure")
		ctx = context.WithValue(ctx, goa.ServiceKey, "secured_service")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

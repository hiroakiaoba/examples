// Code generated by goa v3.1.2, DO NOT EDIT.
//
// divider gRPC server
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package server

import (
	"context"

	divider "goa.design/examples/error/gen/divider"
	dividerpb "goa.design/examples/error/gen/grpc/divider/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the dividerpb.DividerServer interface.
type Server struct {
	IntegerDivideH goagrpc.UnaryHandler
	DivideH        goagrpc.UnaryHandler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the divider service endpoints.
func New(e *divider.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		IntegerDivideH: NewIntegerDivideHandler(e.IntegerDivide, uh),
		DivideH:        NewDivideHandler(e.Divide, uh),
	}
}

// NewIntegerDivideHandler creates a gRPC handler which serves the "divider"
// service "integer_divide" endpoint.
func NewIntegerDivideHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeIntegerDivideRequest, EncodeIntegerDivideResponse)
	}
	return h
}

// IntegerDivide implements the "IntegerDivide" method in
// dividerpb.DividerServer interface.
func (s *Server) IntegerDivide(ctx context.Context, message *dividerpb.IntegerDivideRequest) (*dividerpb.IntegerDivideResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "integer_divide")
	ctx = context.WithValue(ctx, goa.ServiceKey, "divider")
	resp, err := s.IntegerDivideH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "has_remainder":
				return nil, goagrpc.NewStatusError(codes.Unknown, err, goagrpc.NewErrorResponse(err))
			case "div_by_zero":
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err, goagrpc.NewErrorResponse(err))
			case "timeout":
				return nil, goagrpc.NewStatusError(codes.DeadlineExceeded, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*dividerpb.IntegerDivideResponse), nil
}

// NewDivideHandler creates a gRPC handler which serves the "divider" service
// "divide" endpoint.
func NewDivideHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDivideRequest, EncodeDivideResponse)
	}
	return h
}

// Divide implements the "Divide" method in dividerpb.DividerServer interface.
func (s *Server) Divide(ctx context.Context, message *dividerpb.DivideRequest) (*dividerpb.DivideResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "divide")
	ctx = context.WithValue(ctx, goa.ServiceKey, "divider")
	resp, err := s.DivideH.Handle(ctx, message)
	if err != nil {
		if en, ok := err.(ErrorNamer); ok {
			switch en.ErrorName() {
			case "div_by_zero":
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err, goagrpc.NewErrorResponse(err))
			case "timeout":
				return nil, goagrpc.NewStatusError(codes.DeadlineExceeded, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*dividerpb.DivideResponse), nil
}

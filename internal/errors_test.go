package internal

import (
	"errors"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestNewLogicError(t *testing.T) {
	defaultError := errors.New("simple error")
	customError := GenerateError(defaultError, "localized text")
	if err, ok := customError.(*LogicError); ok {
		fmt.Printf("localized message is: %s\r\n", err.localizedMessage)
		statusErr, _ := status.New(codes.FailedPrecondition,
			"default error text").
			WithDetails(&errdetails.LocalizedMessage{
				Message: err.localizedMessage,
			})
		fmt.Printf("%v", statusErr.Err())
	}
}

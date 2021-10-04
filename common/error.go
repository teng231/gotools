package common

import (
	"runtime/debug"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Err(code codes.Code, msg string) error {
	st := status.New(code, msg)
	badreq := &errdetails.BadRequest_FieldViolation{
		Description: string(debug.Stack()),
	}
	det, err := st.WithDetails(badreq)
	if err != nil {
		return err
	}
	return det.Err()
}

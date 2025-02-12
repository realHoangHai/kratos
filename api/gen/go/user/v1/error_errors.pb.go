// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package userv1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsNotLoggedIn(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_NOT_LOGGED_IN.String() && e.Code == 401
}

func ErrorNotLoggedIn(format string, args ...interface{}) *errors.Error {
	return errors.New(401, UserErrorReason_NOT_LOGGED_IN.String(), fmt.Sprintf(format, args...))
}

func IsAccessForbidden(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_ACCESS_FORBIDDEN.String() && e.Code == 403
}

func ErrorAccessForbidden(format string, args ...interface{}) *errors.Error {
	return errors.New(403, UserErrorReason_ACCESS_FORBIDDEN.String(), fmt.Sprintf(format, args...))
}

func IsResourceNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_RESOURCE_NOT_FOUND.String() && e.Code == 404
}

func ErrorResourceNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, UserErrorReason_RESOURCE_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsMethodNotAllowed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_METHOD_NOT_ALLOWED.String() && e.Code == 405
}

func ErrorMethodNotAllowed(format string, args ...interface{}) *errors.Error {
	return errors.New(405, UserErrorReason_METHOD_NOT_ALLOWED.String(), fmt.Sprintf(format, args...))
}

func IsRequestTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_REQUEST_TIMEOUT.String() && e.Code == 408
}

func ErrorRequestTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(408, UserErrorReason_REQUEST_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

func IsInternalServerError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_INTERNAL_SERVER_ERROR.String() && e.Code == 500
}

func ErrorInternalServerError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, UserErrorReason_INTERNAL_SERVER_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsNotImplemented(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_NOT_IMPLEMENTED.String() && e.Code == 501
}

func ErrorNotImplemented(format string, args ...interface{}) *errors.Error {
	return errors.New(501, UserErrorReason_NOT_IMPLEMENTED.String(), fmt.Sprintf(format, args...))
}

func IsNetworkError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_NETWORK_ERROR.String() && e.Code == 502
}

func ErrorNetworkError(format string, args ...interface{}) *errors.Error {
	return errors.New(502, UserErrorReason_NETWORK_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsServiceUnavailable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_SERVICE_UNAVAILABLE.String() && e.Code == 503
}

func ErrorServiceUnavailable(format string, args ...interface{}) *errors.Error {
	return errors.New(503, UserErrorReason_SERVICE_UNAVAILABLE.String(), fmt.Sprintf(format, args...))
}

func IsNetworkTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_NETWORK_TIMEOUT.String() && e.Code == 504
}

func ErrorNetworkTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(504, UserErrorReason_NETWORK_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

func IsRequestNotSupport(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_REQUEST_NOT_SUPPORT.String() && e.Code == 505
}

func ErrorRequestNotSupport(format string, args ...interface{}) *errors.Error {
	return errors.New(505, UserErrorReason_REQUEST_NOT_SUPPORT.String(), fmt.Sprintf(format, args...))
}

func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_USER_NOT_FOUND.String() && e.Code == 600
}

func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(600, UserErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsIncorrectPassword(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_INCORRECT_PASSWORD.String() && e.Code == 599
}

func ErrorIncorrectPassword(format string, args ...interface{}) *errors.Error {
	return errors.New(599, UserErrorReason_INCORRECT_PASSWORD.String(), fmt.Sprintf(format, args...))
}

func IsUserFreeze(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_USER_FREEZE.String() && e.Code == 598
}

func ErrorUserFreeze(format string, args ...interface{}) *errors.Error {
	return errors.New(598, UserErrorReason_USER_FREEZE.String(), fmt.Sprintf(format, args...))
}

func IsInvalidUserid(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_INVALID_USERID.String() && e.Code == 101
}

func ErrorInvalidUserid(format string, args ...interface{}) *errors.Error {
	return errors.New(101, UserErrorReason_INVALID_USERID.String(), fmt.Sprintf(format, args...))
}

func IsInvalidPassword(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_INVALID_PASSWORD.String() && e.Code == 102
}

func ErrorInvalidPassword(format string, args ...interface{}) *errors.Error {
	return errors.New(102, UserErrorReason_INVALID_PASSWORD.String(), fmt.Sprintf(format, args...))
}

func IsTokenExpired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_TOKEN_EXPIRED.String() && e.Code == 103
}

func ErrorTokenExpired(format string, args ...interface{}) *errors.Error {
	return errors.New(103, UserErrorReason_TOKEN_EXPIRED.String(), fmt.Sprintf(format, args...))
}

func IsInvalidToken(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_INVALID_TOKEN.String() && e.Code == 104
}

func ErrorInvalidToken(format string, args ...interface{}) *errors.Error {
	return errors.New(104, UserErrorReason_INVALID_TOKEN.String(), fmt.Sprintf(format, args...))
}

func IsTokenNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_TOKEN_NOT_EXIST.String() && e.Code == 105
}

func ErrorTokenNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(105, UserErrorReason_TOKEN_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

func IsUserNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == UserErrorReason_USER_NOT_EXIST.String() && e.Code == 106
}

func ErrorUserNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(106, UserErrorReason_USER_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package clientv1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 400
func IsBadRequest(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_BAD_REQUEST.String() && e.Code == 400
}

// 400
func ErrorBadRequest(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ClientErrorReason_BAD_REQUEST.String(), fmt.Sprintf(format, args...))
}

// 401
func IsNotLoggedIn(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_NOT_LOGGED_IN.String() && e.Code == 401
}

// 401
func ErrorNotLoggedIn(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ClientErrorReason_NOT_LOGGED_IN.String(), fmt.Sprintf(format, args...))
}

// 403
func IsAccessForbidden(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_ACCESS_FORBIDDEN.String() && e.Code == 403
}

// 403
func ErrorAccessForbidden(format string, args ...interface{}) *errors.Error {
	return errors.New(403, ClientErrorReason_ACCESS_FORBIDDEN.String(), fmt.Sprintf(format, args...))
}

// 404
func IsResourceNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_RESOURCE_NOT_FOUND.String() && e.Code == 404
}

// 404
func ErrorResourceNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ClientErrorReason_RESOURCE_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 405
func IsMethodNotAllowed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_METHOD_NOT_ALLOWED.String() && e.Code == 405
}

// 405
func ErrorMethodNotAllowed(format string, args ...interface{}) *errors.Error {
	return errors.New(405, ClientErrorReason_METHOD_NOT_ALLOWED.String(), fmt.Sprintf(format, args...))
}

// 408
func IsRequestTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_REQUEST_TIMEOUT.String() && e.Code == 408
}

// 408
func ErrorRequestTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(408, ClientErrorReason_REQUEST_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

// 500
func IsInternalServerError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_INTERNAL_SERVER_ERROR.String() && e.Code == 500
}

// 500
func ErrorInternalServerError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ClientErrorReason_INTERNAL_SERVER_ERROR.String(), fmt.Sprintf(format, args...))
}

// 501
func IsNotImplemented(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_NOT_IMPLEMENTED.String() && e.Code == 501
}

// 501
func ErrorNotImplemented(format string, args ...interface{}) *errors.Error {
	return errors.New(501, ClientErrorReason_NOT_IMPLEMENTED.String(), fmt.Sprintf(format, args...))
}

// 502
func IsNetworkError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_NETWORK_ERROR.String() && e.Code == 502
}

// 502
func ErrorNetworkError(format string, args ...interface{}) *errors.Error {
	return errors.New(502, ClientErrorReason_NETWORK_ERROR.String(), fmt.Sprintf(format, args...))
}

// 503
func IsServiceUnavailable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_SERVICE_UNAVAILABLE.String() && e.Code == 503
}

// 503
func ErrorServiceUnavailable(format string, args ...interface{}) *errors.Error {
	return errors.New(503, ClientErrorReason_SERVICE_UNAVAILABLE.String(), fmt.Sprintf(format, args...))
}

// 504
func IsNetworkTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_NETWORK_TIMEOUT.String() && e.Code == 504
}

// 504
func ErrorNetworkTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(504, ClientErrorReason_NETWORK_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

// 505
func IsRequestNotSupport(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_REQUEST_NOT_SUPPORT.String() && e.Code == 505
}

// 505
func ErrorRequestNotSupport(format string, args ...interface{}) *errors.Error {
	return errors.New(505, ClientErrorReason_REQUEST_NOT_SUPPORT.String(), fmt.Sprintf(format, args...))
}

// 用户ID无效
func IsInvalidUserid(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_INVALID_USERID.String() && e.Code == 101
}

// 用户ID无效
func ErrorInvalidUserid(format string, args ...interface{}) *errors.Error {
	return errors.New(101, ClientErrorReason_INVALID_USERID.String(), fmt.Sprintf(format, args...))
}

// 密码无效
func IsInvalidPassword(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_INVALID_PASSWORD.String() && e.Code == 102
}

// 密码无效
func ErrorInvalidPassword(format string, args ...interface{}) *errors.Error {
	return errors.New(102, ClientErrorReason_INVALID_PASSWORD.String(), fmt.Sprintf(format, args...))
}

// token过期
func IsTokenExpired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_TOKEN_EXPIRED.String() && e.Code == 103
}

// token过期
func ErrorTokenExpired(format string, args ...interface{}) *errors.Error {
	return errors.New(103, ClientErrorReason_TOKEN_EXPIRED.String(), fmt.Sprintf(format, args...))
}

// token无效
func IsInvalidToken(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_INVALID_TOKEN.String() && e.Code == 104
}

// token无效
func ErrorInvalidToken(format string, args ...interface{}) *errors.Error {
	return errors.New(104, ClientErrorReason_INVALID_TOKEN.String(), fmt.Sprintf(format, args...))
}

// token不存在
func IsTokenNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_TOKEN_NOT_EXIST.String() && e.Code == 105
}

// token不存在
func ErrorTokenNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(105, ClientErrorReason_TOKEN_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

// 用户不存在
func IsUserNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_USER_NOT_EXIST.String() && e.Code == 106
}

// 用户不存在
func ErrorUserNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(106, ClientErrorReason_USER_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

// 账号冻结
func IsUserFreeze(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ClientErrorReason_USER_FREEZE.String() && e.Code == 107
}

// 账号冻结
func ErrorUserFreeze(format string, args ...interface{}) *errors.Error {
	return errors.New(107, ClientErrorReason_USER_FREEZE.String(), fmt.Sprintf(format, args...))
}

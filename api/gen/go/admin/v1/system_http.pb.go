// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             (unknown)
// source: admin/v1/system.proto

package adminv1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	v1 "github.com/realHoangHai/kratos/api/gen/go/common/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSystemServiceCreateSystem = "/admin.v1.SystemService/CreateSystem"
const OperationSystemServiceDeleteSystem = "/admin.v1.SystemService/DeleteSystem"
const OperationSystemServiceGetSystem = "/admin.v1.SystemService/GetSystem"
const OperationSystemServiceListSystem = "/admin.v1.SystemService/ListSystem"
const OperationSystemServiceUpdateSystem = "/admin.v1.SystemService/UpdateSystem"

type SystemServiceHTTPServer interface {
	// CreateSystem Create system settings
	CreateSystem(context.Context, *CreateSystemRequest) (*System, error)
	// DeleteSystem Delete system settings
	DeleteSystem(context.Context, *DeleteSystemRequest) (*emptypb.Empty, error)
	// GetSystem Get system setting data
	GetSystem(context.Context, *GetSystemRequest) (*System, error)
	// ListSystem Get a list of system settings
	ListSystem(context.Context, *v1.PagingRequest) (*ListSystemResponse, error)
	// UpdateSystem Update system settings
	UpdateSystem(context.Context, *UpdateSystemRequest) (*System, error)
}

func RegisterSystemServiceHTTPServer(s *http.Server, srv SystemServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/system", _SystemService_ListSystem0_HTTP_Handler(srv))
	r.GET("/admin/v1/system/{id}", _SystemService_GetSystem0_HTTP_Handler(srv))
	r.POST("/admin/v1/system", _SystemService_CreateSystem0_HTTP_Handler(srv))
	r.PUT("/admin/v1/system/{id}", _SystemService_UpdateSystem0_HTTP_Handler(srv))
	r.DELETE("/admin/v1/system/{id}", _SystemService_DeleteSystem0_HTTP_Handler(srv))
}

func _SystemService_ListSystem0_HTTP_Handler(srv SystemServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemServiceListSystem)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListSystem(ctx, req.(*v1.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSystemResponse)
		return ctx.Result(200, reply)
	}
}

func _SystemService_GetSystem0_HTTP_Handler(srv SystemServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSystemRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemServiceGetSystem)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSystem(ctx, req.(*GetSystemRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*System)
		return ctx.Result(200, reply)
	}
}

func _SystemService_CreateSystem0_HTTP_Handler(srv SystemServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateSystemRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemServiceCreateSystem)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSystem(ctx, req.(*CreateSystemRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*System)
		return ctx.Result(200, reply)
	}
}

func _SystemService_UpdateSystem0_HTTP_Handler(srv SystemServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateSystemRequest
		if err := ctx.Bind(&in.System); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemServiceUpdateSystem)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateSystem(ctx, req.(*UpdateSystemRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*System)
		return ctx.Result(200, reply)
	}
}

func _SystemService_DeleteSystem0_HTTP_Handler(srv SystemServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteSystemRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSystemServiceDeleteSystem)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteSystem(ctx, req.(*DeleteSystemRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type SystemServiceHTTPClient interface {
	CreateSystem(ctx context.Context, req *CreateSystemRequest, opts ...http.CallOption) (rsp *System, err error)
	DeleteSystem(ctx context.Context, req *DeleteSystemRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	GetSystem(ctx context.Context, req *GetSystemRequest, opts ...http.CallOption) (rsp *System, err error)
	ListSystem(ctx context.Context, req *v1.PagingRequest, opts ...http.CallOption) (rsp *ListSystemResponse, err error)
	UpdateSystem(ctx context.Context, req *UpdateSystemRequest, opts ...http.CallOption) (rsp *System, err error)
}

type SystemServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewSystemServiceHTTPClient(client *http.Client) SystemServiceHTTPClient {
	return &SystemServiceHTTPClientImpl{client}
}

func (c *SystemServiceHTTPClientImpl) CreateSystem(ctx context.Context, in *CreateSystemRequest, opts ...http.CallOption) (*System, error) {
	var out System
	pattern := "/admin/v1/system"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemServiceCreateSystem))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SystemServiceHTTPClientImpl) DeleteSystem(ctx context.Context, in *DeleteSystemRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/system/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemServiceDeleteSystem))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SystemServiceHTTPClientImpl) GetSystem(ctx context.Context, in *GetSystemRequest, opts ...http.CallOption) (*System, error) {
	var out System
	pattern := "/admin/v1/system/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemServiceGetSystem))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SystemServiceHTTPClientImpl) ListSystem(ctx context.Context, in *v1.PagingRequest, opts ...http.CallOption) (*ListSystemResponse, error) {
	var out ListSystemResponse
	pattern := "/admin/v1/system"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSystemServiceListSystem))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SystemServiceHTTPClientImpl) UpdateSystem(ctx context.Context, in *UpdateSystemRequest, opts ...http.CallOption) (*System, error) {
	var out System
	pattern := "/admin/v1/system/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSystemServiceUpdateSystem))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in.System, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

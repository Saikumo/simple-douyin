// Code generated by Kitex v0.3.2. DO NOT EDIT.

package usersrv

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"google.golang.org/protobuf/proto"
	"saikumo.org/simple-douyin/cmd/relation/kitex_gen/user"
)

func serviceInfo() *kitex.ServiceInfo {
	return userSrvServiceInfo
}

var userSrvServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserSrv"
	handlerType := (*user.UserSrv)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register":    kitex.NewMethodInfo(registerHandler, newRegisterArgs, newRegisterResult, false),
		"Login":       kitex.NewMethodInfo(loginHandler, newLoginArgs, newLoginResult, false),
		"GetUserById": kitex.NewMethodInfo(getUserByIdHandler, newGetUserByIdArgs, newGetUserByIdResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "user",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.3.2",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DouyinUserRegisterRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserSrv).Register(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *RegisterArgs:
		success, err := handler.(user.UserSrv).Register(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*RegisterResult)
		realResult.Success = success
	}
	return nil
}
func newRegisterArgs() interface{} {
	return &RegisterArgs{}
}

func newRegisterResult() interface{} {
	return &RegisterResult{}
}

type RegisterArgs struct {
	Req *user.DouyinUserRegisterRequest
}

func (p *RegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in RegisterArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *RegisterArgs) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserRegisterRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var RegisterArgs_Req_DEFAULT *user.DouyinUserRegisterRequest

func (p *RegisterArgs) GetReq() *user.DouyinUserRegisterRequest {
	if !p.IsSetReq() {
		return RegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *RegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

type RegisterResult struct {
	Success *user.DouyinUserRegisterResponse
}

var RegisterResult_Success_DEFAULT *user.DouyinUserRegisterResponse

func (p *RegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in RegisterResult")
	}
	return proto.Marshal(p.Success)
}

func (p *RegisterResult) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserRegisterResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *RegisterResult) GetSuccess() *user.DouyinUserRegisterResponse {
	if !p.IsSetSuccess() {
		return RegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *RegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DouyinUserRegisterResponse)
}

func (p *RegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DouyinUserRegisterRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserSrv).Login(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *LoginArgs:
		success, err := handler.(user.UserSrv).Login(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*LoginResult)
		realResult.Success = success
	}
	return nil
}
func newLoginArgs() interface{} {
	return &LoginArgs{}
}

func newLoginResult() interface{} {
	return &LoginResult{}
}

type LoginArgs struct {
	Req *user.DouyinUserRegisterRequest
}

func (p *LoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in LoginArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *LoginArgs) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserRegisterRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var LoginArgs_Req_DEFAULT *user.DouyinUserRegisterRequest

func (p *LoginArgs) GetReq() *user.DouyinUserRegisterRequest {
	if !p.IsSetReq() {
		return LoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *LoginArgs) IsSetReq() bool {
	return p.Req != nil
}

type LoginResult struct {
	Success *user.DouyinUserRegisterResponse
}

var LoginResult_Success_DEFAULT *user.DouyinUserRegisterResponse

func (p *LoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in LoginResult")
	}
	return proto.Marshal(p.Success)
}

func (p *LoginResult) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserRegisterResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *LoginResult) GetSuccess() *user.DouyinUserRegisterResponse {
	if !p.IsSetSuccess() {
		return LoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *LoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DouyinUserRegisterResponse)
}

func (p *LoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func getUserByIdHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DouyinUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserSrv).GetUserById(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetUserByIdArgs:
		success, err := handler.(user.UserSrv).GetUserById(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetUserByIdResult)
		realResult.Success = success
	}
	return nil
}
func newGetUserByIdArgs() interface{} {
	return &GetUserByIdArgs{}
}

func newGetUserByIdResult() interface{} {
	return &GetUserByIdResult{}
}

type GetUserByIdArgs struct {
	Req *user.DouyinUserRequest
}

func (p *GetUserByIdArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetUserByIdArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetUserByIdArgs) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetUserByIdArgs_Req_DEFAULT *user.DouyinUserRequest

func (p *GetUserByIdArgs) GetReq() *user.DouyinUserRequest {
	if !p.IsSetReq() {
		return GetUserByIdArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetUserByIdArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetUserByIdResult struct {
	Success *user.DouyinUserResponse
}

var GetUserByIdResult_Success_DEFAULT *user.DouyinUserResponse

func (p *GetUserByIdResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetUserByIdResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetUserByIdResult) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetUserByIdResult) GetSuccess() *user.DouyinUserResponse {
	if !p.IsSetSuccess() {
		return GetUserByIdResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetUserByIdResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DouyinUserResponse)
}

func (p *GetUserByIdResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, Req *user.DouyinUserRegisterRequest) (r *user.DouyinUserRegisterResponse, err error) {
	var _args RegisterArgs
	_args.Req = Req
	var _result RegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, Req *user.DouyinUserRegisterRequest) (r *user.DouyinUserRegisterResponse, err error) {
	var _args LoginArgs
	_args.Req = Req
	var _result LoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) GetUserById(ctx context.Context, Req *user.DouyinUserRequest) (r *user.DouyinUserResponse, err error) {
	var _args GetUserByIdArgs
	_args.Req = Req
	var _result GetUserByIdResult
	if err = p.c.Call(ctx, "GetUserById", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

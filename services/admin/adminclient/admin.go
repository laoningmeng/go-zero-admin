// Code generated by goctl. DO NOT EDIT.
// Source: admin.proto

package adminclient

import (
	"context"

	"github.com/laoningmeng/go-zero-admin/services/admin/admin"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Admin interface {
		Login(ctx context.Context, in *admin.LoginReq, opts ...grpc.CallOption) (*admin.LoginReply, error)
		Logout(ctx context.Context, in *admin.LogoutReq, opts ...grpc.CallOption) (*admin.LogoutReply, error)
		UserAdd(ctx context.Context, in *admin.UserAddReq, opts ...grpc.CallOption) (*admin.UserAddReply, error)
		UserUpdate(ctx context.Context, in *admin.UserUpdateReq, opts ...grpc.CallOption) (*admin.UserUpdateReply, error)
		UserQuery(ctx context.Context, in *admin.UserQueryReq, opts ...grpc.CallOption) (*admin.UserQueryReply, error)
		UserList(ctx context.Context, in *admin.UserListReq, opts ...grpc.CallOption) (*admin.UserListReply, error)
		UserDelete(ctx context.Context, in *admin.UserDeleteReq, opts ...grpc.CallOption) (*admin.UserDeleteReply, error)
		//角色管理
		RoleAdd(ctx context.Context, in *admin.RoleAddReq, opts ...grpc.CallOption) (*admin.RoleAddReply, error)
		RoleUpdate(ctx context.Context, in *admin.RoleUpdateReq, opts ...grpc.CallOption) (*admin.RoleUpdateReply, error)
		RoleQuery(ctx context.Context, in *admin.RoleQueryReq, opts ...grpc.CallOption) (*admin.RoleQueryReply, error)
		RoleList(ctx context.Context, in *admin.RoleListReq, opts ...grpc.CallOption) (*admin.RoleListReply, error)
		RoleDelete(ctx context.Context, in *admin.RoleDeleteReq, opts ...grpc.CallOption) (*admin.RoleDeleteReply, error)

		//权限管理
		RuleAdd(ctx context.Context, in *admin.RuleAddReq, opts ...grpc.CallOption) (*admin.RuleAddReply, error)
		RuleUpdate(ctx context.Context, in *admin.RuleUpdateReq, opts ...grpc.CallOption) (*admin.RuleUpdateReply, error)
		RuleQuery(ctx context.Context, in *admin.RuleQueryReq, opts ...grpc.CallOption) (*admin.RuleQueryReply, error)
		RuleList(ctx context.Context, in *admin.RuleListReq, opts ...grpc.CallOption) (*admin.RuleListReply, error)
		RuleDelete(ctx context.Context, in *admin.RuleDeleteReq, opts ...grpc.CallOption) (*admin.RuleDeleteReply, error)
	}

	defaultAdmin struct {
		cli zrpc.Client
	}
)

func NewAdmin(cli zrpc.Client) Admin {
	return &defaultAdmin{
		cli: cli,
	}
}

func (m *defaultAdmin) Login(ctx context.Context, in *admin.LoginReq, opts ...grpc.CallOption) (*admin.LoginReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultAdmin) Logout(ctx context.Context, in *admin.LogoutReq, opts ...grpc.CallOption) (*admin.LogoutReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.Logout(ctx, in, opts...)
}

func (m *defaultAdmin) UserAdd(ctx context.Context, in *admin.UserAddReq, opts ...grpc.CallOption) (*admin.UserAddReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.UserAdd(ctx, in, opts...)
}

func (m *defaultAdmin) UserUpdate(ctx context.Context, in *admin.UserUpdateReq, opts ...grpc.CallOption) (*admin.UserUpdateReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.UserUpdate(ctx, in, opts...)
}

func (m *defaultAdmin) UserQuery(ctx context.Context, in *admin.UserQueryReq, opts ...grpc.CallOption) (*admin.UserQueryReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.UserQuery(ctx, in, opts...)
}

func (m *defaultAdmin) UserList(ctx context.Context, in *admin.UserListReq, opts ...grpc.CallOption) (*admin.UserListReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.UserList(ctx, in, opts...)
}

func (m *defaultAdmin) UserDelete(ctx context.Context, in *admin.UserDeleteReq, opts ...grpc.CallOption) (*admin.UserDeleteReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.UserDelete(ctx, in, opts...)
}

func (m *defaultAdmin) RoleAdd(ctx context.Context, in *admin.RoleAddReq, opts ...grpc.CallOption) (*admin.RoleAddReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.RoleAdd(ctx, in, opts...)
}

func (m *defaultAdmin) RoleUpdate(ctx context.Context, in *admin.RoleUpdateReq, opts ...grpc.CallOption) (*admin.RoleUpdateReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.RoleUpdate(ctx, in, opts...)
}

func (m *defaultAdmin) RoleQuery(ctx context.Context, in *admin.RoleQueryReq, opts ...grpc.CallOption) (*admin.RoleQueryReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.RoleQuery(ctx, in, opts...)
}

func (m *defaultAdmin) RoleList(ctx context.Context, in *admin.RoleListReq, opts ...grpc.CallOption) (*admin.RoleListReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.RoleList(ctx, in, opts...)
}

func (m *defaultAdmin) RoleDelete(ctx context.Context, in *admin.RoleDeleteReq, opts ...grpc.CallOption) (*admin.RoleDeleteReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.RoleDelete(ctx, in, opts...)
}

func (m *defaultAdmin) RuleAdd(ctx context.Context, in *admin.RuleAddReq, opts ...grpc.CallOption) (*admin.RuleAddReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.RuleAdd(ctx, in, opts...)
}

func (m *defaultAdmin) RuleUpdate(ctx context.Context, in *admin.RuleUpdateReq, opts ...grpc.CallOption) (*admin.RuleUpdateReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.RuleUpdate(ctx, in, opts...)
}

func (m *defaultAdmin) RuleQuery(ctx context.Context, in *admin.RuleQueryReq, opts ...grpc.CallOption) (*admin.RuleQueryReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.RuleQuery(ctx, in, opts...)
}

func (m *defaultAdmin) RuleList(ctx context.Context, in *admin.RuleListReq, opts ...grpc.CallOption) (*admin.RuleListReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.RuleList(ctx, in, opts...)
}

func (m *defaultAdmin) RuleDelete(ctx context.Context, in *admin.RuleDeleteReq, opts ...grpc.CallOption) (*admin.RuleDeleteReply, error) {
	client := admin.NewAdminClient(m.cli.Conn())
	return client.RuleDelete(ctx, in, opts...)
}

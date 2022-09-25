package auth

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Possible roles for a user.
// New values may be added in the future.
type Access_Role int32

const (
	// Default value. This value is unused.
	Access_ROLE_UNSPECIFIED Access_Role = 0
	// Full access to everything.
	Access_ROLE_ADMIN Access_Role = 1
	// Full access only to user's own data.
	// Read-only access to user's Access.
	Access_ROLE_USER Access_Role = 2
)

// Enum value maps for Access_Role.
var (
	Access_Role_name = map[int32]string{
		0: "ROLE_UNSPECIFIED",
		1: "ROLE_ADMIN",
		2: "ROLE_USER",
	}
	Access_Role_value = map[string]int32{
		"ROLE_UNSPECIFIED": 0,
		"ROLE_ADMIN":       1,
		"ROLE_USER":        2,
	}
)

func (x Access_Role) Enum() *Access_Role {
	p := new(Access_Role)
	*p = x
	return p
}

func (x Access_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Access_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_powerman_example_auth_service_proto_enumTypes[0].Descriptor()
}

// Request.
type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The account to create.
	Account *Account `protobuf: "bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	// The ID to use for the account.
	// This value should be 4-63 characters [a-z0-9-].
	AccountId string `protobuf:"bytes,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_powerman_example_auth_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_powerman_example_auth_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Request.
type SigninIdentityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Different ways to authenticate.
	//
	// Types that are assignable to Auth:
	//	*SigninIdentityRequest_Account
	//	*SigninIdentityRequest_Email
	Auth isSigninIdentityRequest_Auth `protobuf_oneof:"auth"`
}

func (m *SigninIdentityRequest) GetAuth() isSigninIdentityRequest_Auth {
	if m != nil {
		return m.Auth
	}

	return nil
}

type isSigninIdentityRequest_Auth interface {
	isSigninIdentityRequest_Auth()
}

var file_powerman_example_auth_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_powerman_example_auth_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_powerman_example_auth_service_proto_goTypes = []interface{}{
	(Access_Role)(0),              // 0: crnchy.example.auth.Access.Role
	(*CreateAccountRequest)(nil),  // 1
	(*SigninIdentityRequest)(nil), // 2

}

// Account contains data needed for authentication.
type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Format: "accounts/{account_id}".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Default identity connected to the account.
	User *User `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Format: "users/{user_uid}".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// By default set to {account_id}.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Permissions.
	Access *Access `protobuf:"bytes,3,opt,name=access,proto3" json:"access,omitempty"`
}

// Access describes identity's permissions.
type Access struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User's role.
	Role Access_Role `protobuf:"varint,1,opt,name=role,proto3,enum=powerman.example.auth.Access_Role" json:"role,omitempty"`
}

func (x *Access) Reset() {
	*x = Access{}
	if protoimpl.UnsafeEnabled {
		mi := &file_powerman_example_auth_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

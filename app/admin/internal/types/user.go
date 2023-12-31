package types

type UserAddReq struct {
	Username     string `json:"username"`
	RoleId       int32  `json:"role_id"`
	DepartmentId int32  `json:"department_id"`
}

type UserAddResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type UserListReq struct {
	Page         int32  `form:"page"`
	PageSize     int32  `form:"page_size"`
	Username     string `form:"username,optional"`
	RoleId       int32  `form:"role_id,optional"`
	DepartmentId int32  `form:"department_id,optional"`
	Status       int32  `form:"status,optional"`
}

type User struct {
	Id           int64  `json:"id"`
	Username     string `json:"username"`
	RoleName     string `json:"role_name"`
	Avatar       string `json:"avatar"`
	RoleId       int32  `json:"role_id"`
	DepartmentId int32  `json:"department_id"`
	Department   string `json:"department"`
}

type UserListResp struct {
	Code  int32  `json:"code"`
	Total int64  `json:"total"`
	Items []User `json:"items"`
}

type UserDelReq struct {
	Id int64 `json:"id"`
}

type UserDelResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type UserUpdateReq struct {
	Id           int64 `json:"id"`
	RoleId       int32 `json:"role_id"`
	DepartmentId int32 `json:"department_id"`
}

type UserUpdateResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type UserInfo struct {
	User *User    `json:"user"`
	Menu []string `json:"menu"`
	Btn  []string `json:"btn"`
}

type UserInfoResp struct {
	Code    int32     `json:"code"`
	Message string    `json:"message"`
	Data    *UserInfo `json:"data"`
}

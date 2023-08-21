package types

import "time"

type RoleAddReq struct {
	Name   string `json:"name"`
	Title  string `json:"title"`
	Status int32  `json:"status"`
}

type RoleAddResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type RoleListReq struct {
	Page     int32  `json:"page"`
	PageSize int32  `json:"page_size"`
	Name     string `json:"name,optional"`
	Title    string `json:"title,optional"`
	Status   int32  `json:"status,optional"`
}

type Role struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Status    int32     `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RoleListResp struct {
	Code  int32  `json:"code"`
	Total int32  `json:"total"`
	Items []Role `json:"items"`
}

type RoleDelReq struct {
	Id int32 `json:"id"`
}

type RoleDelResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type RoleUpdateReq struct {
	Id     int32  `json:"id"`
	Name   string `json:"name"`
	Title  string `json:"title"`
	Status int32  `json:"status"`
}

type RoleUpdateResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

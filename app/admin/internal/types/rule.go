package types

import "time"

type RuleAddReq struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Type  int32  `json:"type"`
}

type RuleAddResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type RuleListReq struct {
	Page     int32  `json:"page"`
	PageSize int32  `json:"page_size"`
	Name     string `json:"name,optional"`
	Title    string `json:"title,optional"`
	Type     int32  `json:"type,optional"`
}

type Rule struct {
	Id        int32     `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Type      int32     `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RuleListResp struct {
	Code  int32  `json:"code"`
	Total int32  `json:"total"`
	Items []Rule `json:"items"`
}

type RuleDelReq struct {
	Id int32 `json:"id"`
}

type RuleDelResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

type RuleUpdateReq struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Title string `json:"title"`
	Type  int32  `json:"type"`
}

type RuleUpdateResp struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

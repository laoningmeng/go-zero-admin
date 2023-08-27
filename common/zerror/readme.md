# errcode 

## 1. errcode 设计
5位错误码，前两位定位模块， 后面定位具体的业务

0 表示正常，没有异常

1xxxx-4xxxx 为api接口你异常错误码

5xxxx-9xxxx 为后端server异常码

## 2. 码表

### 1. api





| 错误码 | 模块 | 模块码 | 业务码 | 表示                            | 描述             |
| ------ | ---- | ------ | ------ | ------------------------------- | ---------------- |
| 10000  | 通用 | 10     | 000    | API_COMMON_ERROR                | Api异常          |
| 10001  |      | 10     | 001    | API_COMMON_TOKEN_EXPIRE_ERROR   | token失效        |
| 10002  |      | 10     | 002    | API_COMMON_TOKEN_ILLEGAL_ERROR  | 非法token        |
| 10003  |      | 10     | 003    | API_COMMON_TOEKN_GENERATE_ERROR | 生成token异常    |
| 10004  |      | 10     | 004    | API_COMMON_TOKEN_MISS_ERROR     | 缺少token        |
| 11000  | 登录 | 11     | 000    | API_LOGIN_PARAM_ERROR           | 登录账户密码异常 |
|        |      | 11     | 001    |                                 |                  |
|        |      |        |        |                                 |                  |
|        |      |        |        |                                 |                  |
|        |      |        |        |                                 |                  |
|        |      |        |        |                                 |                  |











### 2.server

| 错误码 | 模块 | 模块码 | 业务码 | 表示                | 描述        |
| ------ | ---- | ------ | ------ | ------------------- | ----------- |
|        | 通用 | 50     | 000    | SERVER_COMMON_ERROR | Server 异常 |
|        |      |        |        |                     |             |
|        |      |        |        |                     |             |
|        |      |        |        |                     |             |
|        |      |        |        |                     |             |




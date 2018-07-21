package def

type ErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	Success                     = ErrorResponse{Code: 201, Msg: "成功啦！"}
	ErrorRequestBodyParseFailed = ErrorResponse{Code: 400, Msg: "请求不对(⊙o⊙)…"}
	ErrorNotAuthUser            = ErrorResponse{Code: 401, Msg: "用户没有权限！"}
	ErrorDB                     = ErrorResponse{Code: 500, Msg: "数据库错误~"}
)

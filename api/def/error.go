package def 

type Err struct {
	Msg string `json:"error"`

}

type ErrorResponse struct {
	Code int
	Error Err
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{Code:400,Error:Err{Msg:"request body is not correct"}}
	ErrorNotAuthUser = ErrorResponse{Code:401,Error:Err{Msg:"user is not right"}}
	ErrorDB = ErrorResponse{Code:500,Error:Err{Msg:"DB failed"}}
)
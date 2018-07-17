package def 

type Err struct {
	Msg string `json:"error"`

}

type ErroResponse struct {
	Code int
	Error Err
}

var (
	ErroRequestBodyParseFailed = ErroResponse{Code:0,Error:Err{Msg:"request body is not correct"}}
	ErrorNotAuthUser = ErroResponse{Code:0,Error:Err{Msg:"user is not authed"}}
)
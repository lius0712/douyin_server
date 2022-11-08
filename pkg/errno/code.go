package errno

const (
	SuccessCode             = 200
	ServiceErrCode          = 10001
	ParamErrCode            = 10002
	LoginErrCode            = 10003
	ErrBind                 = 400
	UserNotExistErrCode     = 10004
	UserAlreadyExistErrCode = 10005
)

var (
	Success             = NewErrNo(SuccessCode, "Success")
	ServiceErr          = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr            = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	LoginErr            = NewErrNo(LoginErrCode, "Wrong username or password")
	UserNotExistErr     = NewErrNo(UserNotExistErrCode, "User does not exists")
	UserAlreadyExistErr = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	ErrHttpBind         = NewErrNo(ErrBind, "Error occurred while binding the request body to the struct")
)

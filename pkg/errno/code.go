package errno

const (
	SuccessCode             = 0
	ServiceErrCode          = 10001
	ParamErrCode            = 10002
	LoginErrCode            = 10003
	ErrBindCode             = 400
	ErrDecodingFailed       = 500
	UserNotExistErrCode     = 10004
	UserAlreadyExistErrCode = 10005
)

// HttpServer code
const (
	ErrSuccess = 0
)
const (
	// ErrDatabase - 500: Database base error.
	ErrDatabaseCode int32 = iota + 100101
	// ErrRecordNotFound - 500: record not found error
	ErrRecordNotFound
	// ErrInvalidTransaction - 500: invalid transaction when you are trying to `Commit` or `Rollback`
	ErrInvalidTransaction
	// ErrNotImplemented - 500: not implemented
	ErrNotImplemented
	// ErrMissingWhereClause - 500: missing where clause
	ErrMissingWhereClause
)

// JWT code
const (
	ValidationErrorMalformed        uint32 = 1 << iota // Token is malformed
	ValidationErrorUnverifiable                        // Token could not be verified because of signing problems
	ValidationErrorSignatureInvalid                    // Signature validation failed

	// Standard Claim validation errors
	ValidationErrorAudience      // AUD validation failed
	ValidationErrorExpired       // EXP validation failed
	ValidationErrorIssuedAt      // IAT validation failed
	ValidationErrorIssuer        // ISS validation failed
	ValidationErrorNotValidYet   // NBF validation failed
	ValidationErrorId            // JTI validation failed
	ValidationErrorClaimsInvalid // Generic claims validation error
	ValidationErrorTokenInvalid  // Generic claims validation error
)

// server error
var (
	Success             = NewErrNo(SuccessCode, "Success")
	ServiceErr          = NewErrNo(ServiceErrCode, "Service is unable to start successfully")
	ParamErr            = NewErrNo(ParamErrCode, "Wrong Parameter has been given")
	DecodingFailed      = NewErrNo(ErrDecodingFailed, "Decoding failed due to an error with the data")
	LoginErr            = NewErrNo(LoginErrCode, "Wrong username or password")
	UserNotExistErr     = NewErrNo(UserNotExistErrCode, "User does not exists")
	UserAlreadyExistErr = NewErrNo(UserAlreadyExistErrCode, "User already exists")
	ErrHttpBind         = NewErrNo(ErrBindCode, "Error occurred while binding the request body to the struct")
	ErrDatabase         = NewErrNo(ErrDatabaseCode, "Database error")
	ErrBind             = NewErrNo(ErrBindCode, "Error occurred while binding the request body to the struct")
)

var (
	HttpSuccess = NewHttpErr(ErrSuccess, 200, "OK")
)

var (
	ErrorMalformed        = NewTokenErr(int32(ValidationErrorMalformed), "Token is malformed")
	ErrorUnverifiable     = NewTokenErr(int32(ValidationErrorUnverifiable), "Token could not be verified because of signing problems")
	ErrorSignatureInvalid = NewTokenErr(int32(ValidationErrorSignatureInvalid), "Signature validation failed")
	ErrorAudience         = NewTokenErr(int32(ValidationErrorAudience), "AUD validation failed")                // AUD validation failed
	ErrorExpired          = NewTokenErr(int32(ValidationErrorExpired), "EXP validation failed")                 // EXP validation failed
	ErrorIssuedAt         = NewTokenErr(int32(ValidationErrorIssuedAt), "IAT validation failed")                // IAT validation failed
	ErrorIssuer           = NewTokenErr(int32(ValidationErrorIssuer), "ISS validation failed")                  // ISS validation failed
	ErrorNotValidYet      = NewTokenErr(int32(ValidationErrorNotValidYet), "NBF validation failed")             // NBF validation failed
	ErrorId               = NewTokenErr(int32(ValidationErrorId), "JTI validation failed")                      // JTI validation failed
	ErrorClaimsInvalid    = NewTokenErr(int32(ValidationErrorClaimsInvalid), "Generic claims validation error") // Generic claims validation error
	ErrorTokenInvalid     = NewTokenErr(int32(ValidationErrorTokenInvalid), "Couldn't handle this token")
)

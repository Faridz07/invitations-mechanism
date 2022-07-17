package constant

var (
	ErrInvalidRequest          = "invalid request"
	ErrInvalidEmail            = "invalid email"
	ErrPasswordDoesntMatch     = "password doesn't match"
	ErrUserDoesntExist         = "user doesn't exist!"
	ErrLoginFailed             = "email or password doesn't match!"
	ErrHashPasswordDoesntMatch = "user exist but hashpassword doesn't match!"
	ErrToHashPassword          = "failed to hash password"
	ErrToCompareHashPassword   = "failed to compare password"
	ErrUserLoginUnAuthorized   = "unauthorized! you not allowed to login in this page!"
	ErrSomethingWhenWrong      = "something when wrong, please try again!"
)

var (
	EmailRegex = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

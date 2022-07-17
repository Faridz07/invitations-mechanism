package constant

var (
	InvalidRequest      = "invalid request"
	EmailRegex          = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	InvalidEmail        = "invalid email"
	PasswordDoesntMatch = "password doesn't match"
	UserDoesntExist     = "user doesn't exist!"
	LoginFailed         = "email or password doesn't match!"
)

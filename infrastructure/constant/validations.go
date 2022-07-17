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
	ErrInvalidSigningMethod    = "Signing method invalid"
	ErrInvalidToken            = "invalid jwt token"
	ErrRoleName                = "invalid role name"
	ErrMapClaimsNotFound       = "map jwt claims not found!"
	ErrGenerateInvitation      = "failed to generate invitation, try again!"
	ErrInvalidCode             = "invalid invitation code, try again!"
	ErrGetInvitation           = "failed to get invitation, try again!"
	ErrCodeEmpty               = "invitation code empty"
	ErrInvalidDeviceId         = "invalid deviceId, please reinstall apps!"
	ErrRetryInvalidCode        = "invalid invitation code, you have %d chances left!"
	ErrToManyAttempts          = "too many failed login attempts, please try again in %v!"
)

var (
	EmailRegex = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
	Charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

package constants

// JWT and Credential Constants
const (
	AuthCookieMaxAge = OneDayInHour * OneHourInMinute * OneMinuteInSecond
	AuthCookieName   = "accessToken"
	JwtSubClaim      = "sub"
	JwtExpClaim      = "exp"
	JwtUsernameClaim = "username"
)

// Path and Param Constant
const (
	UsernameParam  = "username"
	ProjectIdParam = "projectId"
)

var skipValidatePath_POST = []string{"/api/user/", "/api/user/login"}

func GetSkipValidatePath() []string {
	return skipValidatePath_POST
}

// Locals (Context) Constant
const (
	UsernameContext = "username"
	UserIdContext   = "userId"
)

// Date and Time Constant
const (
	OneDayInHour      = 24
	OneHourInMinute   = 60
	OneMinuteInSecond = 60
)

package constants

// JWT and Credential Constants
const (
	JwtMaxAge     = 24 * 60 * 60
	JwtCookieName = "accessToken"
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

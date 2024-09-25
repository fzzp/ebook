package formdto

type CreateUserRequest struct {
	Email    string `form:"email"`
	Password string `form:"password"`
	Username string `form:"username"`
}

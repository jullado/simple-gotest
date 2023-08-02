package login

var (
	userAdmin = "admin"
	passAdmin = "admin01"
)

func Login(username, password string) string {
	// check username not found
	if username == "" {
		return "login fail : username not found"
	}

	// check password not found
	if password == "" {
		return "login fail : password not found"
	}

	// check password must be between 6-16 characters
	if len(password) < 6 || len(password) > 16 {
		return "login fail : password must be between 6-16 characters"
	}

	// login success
	if username == userAdmin && password == passAdmin {
		return "login success"
	}

	return "login fail : unauthorized"
}

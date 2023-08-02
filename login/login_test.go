package login

import "testing"

func TestLoginSuccess(t *testing.T) {
	username := "admin"
	password := "admin01"

	expected := "login success"
	actual := Login(username, password)

	if expected != actual {
		t.Errorf("success 1 error Login() actual = %v, expected = %v", actual, expected)
	}
}

func TestLoginFail(t *testing.T) {
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		// TODO: Add test cases.
		{
			name: "error1",
			args: args{
				username: "",
				password: "asdadakjd",
			},
			expected: "login fail : username not found",
		},
		{
			name: "error2",
			args: args{
				username: "asdasdad",
				password: "",
			},
			expected: "login fail : password not found",
		},
		{
			name: "error3",
			args: args{
				username: "asdasdad",
				password: "123",
			},
			expected: "login fail : password must be between 6-16 characters",
		},
		{
			name: "error4",
			args: args{
				username: "asdasdad",
				password: "123456789123456789",
			},
			expected: "login fail : password must be between 6-16 characters",
		},
		{
			name: "error5",
			args: args{
				username: "asdasdad",
				password: "asdasdasd",
			},
			expected: "login fail : unauthorized",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if actual := Login(tt.args.username, tt.args.password); actual != tt.expected {
				t.Errorf("Login() actual = %v, expected = %v", actual, tt.expected)
			}
		})
	}
}

func BenchmarkLogin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Login("admin", "admin01")
	}
}

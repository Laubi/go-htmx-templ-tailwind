package persistence

import (
	"testing"
)

func TestUser_GravatarImage(t *testing.T) {

	tests := []struct {
		email string
		want  string
	}{
		{
			"MyEmailAddress@example.com",
			"https://gravatar.com/avatar/84059b07d4be67b806386c0aad8070a23f18836bbaae342275dc0a83414c32ee",
		},
	}
	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			u := User{
				Email: tt.email,
			}
			if got := u.GravatarImage(); got != tt.want {
				t.Errorf("GravatarImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

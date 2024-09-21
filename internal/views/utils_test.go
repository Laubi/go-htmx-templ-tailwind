package views

import "testing"

func TestGravatarImage(t *testing.T) {
	// taken from https://docs.gravatar.com/api/avatars/hash/
	email := "MyEmailAddress@example.com "

	url := gravatarImage(email)

	expected := "https://gravatar.com/avatar/84059b07d4be67b806386c0aad8070a23f18836bbaae342275dc0a83414c32ee?d=identicon"
	if url != expected {
		t.Fatalf(`gravatarImage("%s") = "%s", want "%s"`, email, url, expected)
	}
}

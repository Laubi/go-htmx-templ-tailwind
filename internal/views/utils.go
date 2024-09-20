package views

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

func gravatarImage(email string) string {
	// https://docs.gravatar.com/api/avatars/images/

	sanitizedEmail := strings.ToLower(strings.TrimSpace(email))
	hash := sha256.Sum256([]byte(sanitizedEmail))
	prettyHash := hex.EncodeToString(hash[:])

	return fmt.Sprintf("https://gravatar.com/avatar/%s?d=identicon", prettyHash)
}

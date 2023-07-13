package cache

import "fmt"

const (
	ProductKeyFormat = "product:%s"
	UserKeyFormat    = "user:%s"
)

func ProductKey(id string) string {
	return fmt.Sprintf(ProductKeyFormat, id)
}

func UserKey(email string) string {
	return fmt.Sprintf(UserKeyFormat, email)
}

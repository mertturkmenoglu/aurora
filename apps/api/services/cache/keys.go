package cache

import "fmt"

const (
	ProductKeyFormat = "product:%s"
	UserKeyFormat    = "user:%s"
	BrandKeyFormat   = "brand:%s"
)

func ProductKey(id string) string {
	return fmt.Sprintf(ProductKeyFormat, id)
}

func UserKey(email string) string {
	return fmt.Sprintf(UserKeyFormat, email)
}

func BrandKey(id string) string {
	return fmt.Sprintf(BrandKeyFormat, id)
}

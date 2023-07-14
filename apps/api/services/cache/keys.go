package cache

import "fmt"

type Key string

const (
	AuthKeyFormat    Key = "auth:%s"
	ProductKeyFormat Key = "product:%s"
	UserKeyFormat    Key = "user:%s"
	BrandKeyFormat   Key = "brand:%s"
)

func GetFormattedKey(t Key, key string) string {
	switch t {
	case ProductKeyFormat:
		return fmt.Sprintf(string(ProductKeyFormat), key)
	case UserKeyFormat:
		return fmt.Sprintf(string(UserKeyFormat), key)
	case BrandKeyFormat:
		return fmt.Sprintf(string(BrandKeyFormat), key)
	default:
		return ""
	}
}

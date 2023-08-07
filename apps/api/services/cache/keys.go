package cache

import "fmt"

type Key string

const (
	AuthKeyFormat           Key = "auth:%s"
	ForgotPasswordKeyFormat Key = "forgot-password:%s"
	ProductKeyFormat        Key = "product:%s"
	UserKeyFormat           Key = "user:%s"
	BrandKeyFormat          Key = "brand:%s"
)

const (
	HomeAggregationKey = "home"
)

func GetFormattedKey(t Key, key string) string {
	switch t {
	case ProductKeyFormat:
		return fmt.Sprintf(string(ProductKeyFormat), key)
	case UserKeyFormat:
		return fmt.Sprintf(string(UserKeyFormat), key)
	case BrandKeyFormat:
		return fmt.Sprintf(string(BrandKeyFormat), key)
	case ForgotPasswordKeyFormat:
		return fmt.Sprintf(string(ForgotPasswordKeyFormat), key)
	default:
		return ""
	}
}

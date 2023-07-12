package cache

import "fmt"

const (
	ProductKeyFormat = "product:%s"
)

func ProductKey(id string) string {
	return fmt.Sprintf(ProductKeyFormat, id)
}

package auth

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type LoginDto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (dto LoginDto) GetKey() map[string]types.AttributeValue {
	email, err := attributevalue.Marshal(dto.Email)

	if err != nil {
		panic(err)
	}

	return map[string]types.AttributeValue{
		"email": email,
	}
}

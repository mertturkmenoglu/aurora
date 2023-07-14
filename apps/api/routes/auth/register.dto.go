package auth

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type RegisterDto struct {
	FullName                 string `json:"fullName" binding:"required"`
	Email                    string `json:"email" binding:"required,email"`
	Password                 string `json:"password" binding:"required,min=8,max=64"`
	VerifyPassword           string `json:"verifyPassword" binding:"required,eqfield=Password"`
	HasAcceptedEmailCampaign bool   `json:"hasAcceptedEmailCampaign" validate:"exists"`
}

func (dto RegisterDto) GetKey() map[string]types.AttributeValue {
	email, err := attributevalue.Marshal(dto.Email)

	if err != nil {
		panic(err)
	}

	return map[string]types.AttributeValue{
		"email": email,
	}
}

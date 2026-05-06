package utils

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateRequest(c *gin.Context, req interface{}) bool {
	contentType := c.GetHeader("Content-Type")

	var err error
	if strings.Contains(contentType, "application/json") {
		err = c.ShouldBindJSON(req)
	} else {
		err = c.ShouldBind(req)
	}

	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			errorsMap := make(map[string]string)
			for _, fe := range ve {
				field := toSnakeCase(fe.Field())
				errorsMap[field] = messageFromTag(fe)
			}
			c.JSON(400, gin.H{"errors": errorsMap})
			return false
		}
		c.JSON(400, gin.H{"errors": "invalid request"})
		return false
	}
	return true
}

func messageFromTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "required"
	case "numeric":
		return "must be numeric"
	case "email":
		return "invalid email format"
	case "min":
		return "too short"
	case "max":
		return "to long"
	default:
		return "invalid"
	}
}

func toSnakeCase(str string) string {
	var result []rune
	for i, r := range str {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result = append(result, '_')
		}
		result = append(result, r)
	}
	return strings.ToLower(string(result))
}

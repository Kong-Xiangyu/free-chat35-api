package utils

import (
	"github.com/gin-gonic/gin"
	"math/rand"
)

func MappingModel(model string) string {
	var modelMapping = map[string]string{
		"gpt-3.5-turbo":          "text-davinci-002-render-sha",
		"gpt-3.5-turbo-16k":      "text-davinci-002-render-sha",
		"gpt-3.5-turbo-16k-0613": "text-davinci-002-render-sha",
		"gpt-3.5-turbo-0301":     "text-davinci-002-render-sha",
		"gpt-3.5-turbo-0613":     "text-davinci-002-render-sha",
		"gpt-3.5-turbo-1106":     "text-davinci-002-render-sha",
	}
	if model == "" {
		return "text-davinci-002-render-sha"
	}
	if v, ok := modelMapping[model]; ok {
		return v
	}
	return "text-davinci-002-render-sha"
}

func GenerateID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	id := "chatcmpl-"
	for i := 0; i < length; i++ {
		id += string(charset[rand.Intn(len(charset))])
	}
	return id
}

func ErrorResp(c *gin.Context, code int, msg interface{}, err interface{}) {
	c.AbortWithStatusJSON(code, gin.H{
		"code":  code,
		"msg":   msg,
		"error": err,
	})
	return
}

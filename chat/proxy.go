package chat

import (
	"github.com/gin-gonic/gin"
	"github.com/zbronya/free-chat-to-api/httpclient"
	"github.com/zbronya/free-chat-to-api/utils"
	"io"
	"net/url"
)

var targetURL, _ = url.Parse("https://chat.openai.com")

var client = httpclient.NewReqClient()

func ReverseProxy(c *gin.Context) {
	targetURL.Path = c.Request.URL.Path
	targetURL.RawQuery = c.Request.URL.RawQuery

	headers := make(map[string]string)
	for k, v := range c.Request.Header {
		headers[k] = v[0]
	}

	headers["Host"] = targetURL.Host

	body, _ := io.ReadAll(c.Request.Body)
	resp, err := client.Post(targetURL.String(), headers, body)

	if err != nil {
		utils.ErrorResp(c, 500, "fail to completions", err)
		return
	}

	for k, v := range resp.Header {
		c.Writer.Header()[k] = v
	}

	c.Writer.WriteHeader(resp.StatusCode)
	io.Copy(c.Writer, resp.Body)

}
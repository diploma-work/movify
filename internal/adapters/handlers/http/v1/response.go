package v1

import (
	"github.com/gin-gonic/gin"
	"time"
)

type response struct {
	Result    interface{} `json:"result"`
	Error     interface{} `json:"error"`
	Timestamp time.Time   `json:"timestamp"`
}

func (h v1Handler) respond(c *gin.Context, code int, res interface{}, err interface{}) {
	c.AbortWithStatusJSON(code, response{
		Result:    res,
		Error:     err,
		Timestamp: time.Now(),
	})
}

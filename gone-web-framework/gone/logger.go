package gone

import (
	"log"
	"time"
)

// log 中间件
func Logger() HandleFunc {
	return func(c *Context) {
		t := time.Now()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
		c.Next()
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

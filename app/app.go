package app

import (
	"github.com/gin-gonic/gin"
)

var (
	r = gin.Default()
)

// StartApplication func
func StartApplication() {
	Routes()
	r.Run()
}

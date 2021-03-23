package api

import (
	"github.com/gin-gonic/gin"
	"maxtv_middleware/pkg/pythonReporter"
	"os"
)

func ApiProcessing() {

	r := gin.Default()

	r.GET("/python-reporter/is-paid-maxtv-building", pythonReporter.IsPaidMaxtvBuilding)

	r.Run(":" + os.Getenv("PORT"))
}

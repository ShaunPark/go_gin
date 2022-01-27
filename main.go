package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func testfunc(ctx *gin.Context) {
	body := ctx.Request.Body
	value, _ := ioutil.ReadAll(body)
	fmt.Printf("%v\n", string(value))
}

func main() {
	router := gin.Default()

	router.POST("/", testfunc)
	router.Run("0.0.0.0:8080")
}

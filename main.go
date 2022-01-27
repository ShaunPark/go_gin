package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func testfunc(ctx *gin.Context) {
	body := ctx.Request.Body
	fmt.Printf("%v\n", ctx.Request)

	value, _ := ioutil.ReadAll(body)
	fmt.Printf("%v\n", string(value))

	payload := ctx.PostFormMap("payload")
	fmt.Printf("%v\n", payload)
	for k, v := range payload {
		fmt.Printf("---%v %s\n", k, v)
	}

}

func main() {
	router := gin.Default()

	router.POST("/", testfunc)
	router.Run("0.0.0.0:8080")
}

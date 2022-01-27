package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func testfunc(ctx *gin.Context) {
	body := ctx.Request.Body
	value, _ := ioutil.ReadAll(body)
	var data map[string]interface{}

	json.Unmarshal([]byte(value), &data)
	for k, v := range data {
		fmt.Printf("%s %s\n", k, v)
	}

	payload := ctx.QueryMap("payload")
	fmt.Printf("%v", payload)
}

func main() {
	router := gin.Default()

	router.POST("/", testfunc)
	router.Run("0.0.0.0:8080")
}

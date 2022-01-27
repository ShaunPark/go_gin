package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// func testfunc(ctx *gin.Context) {
// 	body := ctx.Request.Body
// 	value, _ := ioutil.ReadAll(body)
// 	fmt.Printf("%v\n", string(value))
// }

// func main() {
// 	router := gin.Default()

// 	router.POST("/", testfunc)
// 	router.Run("0.0.0.0:8080")
// }

type Payload struct {
	Payload string `json:"payload"`
}

func main() {
	http.HandleFunc("/", func(wr http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost: // 등록
			data := Payload{}
			json.NewDecoder(r.Body).Decode(&data) // 디코딩

			fmt.Printf("%v\n", data)
			fmt.Printf("%v\n", r.Body)

			r.ParseForm() // Parses the request body
			for k, v := range r.Form {
				fmt.Printf("%s = %s", k, v)
			}
		}
	})
	http.ListenAndServe("0.0.0.0:8080", nil)
}

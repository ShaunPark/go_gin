package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ShaunPark/go_gin/types"
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
		case http.MethodPost:
			r.ParseForm() // Parses the request body
			payload := types.InteractiveMessage{}
			json.Unmarshal([]byte(r.Form.Get("payload")), &payload)

			fmt.Printf("User Name : %s", payload.User.UserName)
			for _, a := range payload.Actions {
				fmt.Printf("Action value  : %s", a.Value)
			}
		}
	})
	http.ListenAndServe("0.0.0.0:8080", nil)
}

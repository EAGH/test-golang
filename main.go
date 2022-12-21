package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type ResponseHTTP struct {
	Status  int      `json:"status" bson:"status"`
	Content ArrayObj `json:"content" bson:"content"`
}
type ArrayObj struct {
	Obj []string `json:"obj" bson:"obj"`
}
type ResponseRandom struct {
	//Categories []string `json:"categories" bson:"categories"`
	//CreatedAt string `json:"created_at" bson:"created_at"`
	//IcoUrl string `json:"icon_url" bson:"icon_url"`
	ID string `json:"id" bson:"id"`
	//UpdatedAt string `json:"updated_at" bson:"updated_at"`
	//Url string `json:"url" bson:"url"`
	//Value string `json:"value" bson:"value"`
}

func main() {
	fmt.Println("Serving at 8080")
	r := chi.NewRouter()
	// unique endpoint routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		arr := ArrayObj{}
		var m ResponseRandom
		for i := 0; i < 25; i++ {
			resp, _ := http.Get("https://api.chucknorris.io/jokes/random")
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body) // response body is []byte
			json.Unmarshal([]byte(string(body)), &m)
			fmt.Println(string(body))
			arr.Obj = append(arr.Obj, m.ID)
			fmt.Println(resp)
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		response := ResponseHTTP{
			Status:  200,
			Content: arr}
		json.NewEncoder(w).Encode(response)
	})
	log.Fatal(http.ListenAndServe(":8080", r))
}

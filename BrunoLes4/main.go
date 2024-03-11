package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func main() {
//	http.HandleFunc("/", rootHandler)
//	http.HandleFunc("/hello", helloHandler)
//	http.HandleFunc("/about", aboutHandler)
//
//	fmt.Println("Сервер запущен на http://loclahost:8080")
//	http.ListenAndServe(":8080", nil)
//}
//
//func aboutHandler(writer http.ResponseWriter, request *http.Request) {
//	fmt.Fprintf(writer, "Давай поговорим о нас!")
//}
//
//func helloHandler(writer http.ResponseWriter, request *http.Request) {
//	fmt.Fprintf(writer, "Привет Мэн!")
//}
//
//func rootHandler(writer http.ResponseWriter, request *http.Request) {
//	number := 213
//	wr, _ := json.Marshal(number)
//	writer.Header().Set("Content-Type", "application/json")
//	writer.WriteHeader(http.StatusOK)
//	writer.Write(wr)
//}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

//import "github.com/beego/beego/v2/server/web"
//
//func main() {
//	web.Run()
//}

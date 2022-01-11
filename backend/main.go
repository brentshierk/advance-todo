package main

import (
	"backend/router"
	"fmt"
	"log"
	"net/http"


)

func main()  {
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}



//func checkError(err error) {
//	if err != nil {
//		log.Panic(err)
//	}
//}

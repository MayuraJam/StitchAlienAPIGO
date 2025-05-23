package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MayuraJam/StitchAlienAPIGO/webservice/database"
	route "github.com/MayuraJam/StitchAlienAPIGO/webservice/router"
)

const path = "stitch"
const basePath = "/api"

func main() {
	fmt.Println("Hello my new project")
	database.SetUpDB()
	route.SetupRoutes(basePath, path)
	log.Fatal(http.ListenAndServe(":8074", nil))
}

package main

import (
	"fmt"
	"log"
	"nylas/nylas"

	"github.com/joho/godotenv"
)

const apiKey string = "Bearer nyk_v0_YEZVfhzQRYWIjpKLLJ74JkEDqLCl7RIPqFVOILxItlhjN2nPbYgyzOAaCKNFwKRy"
const grantId string = "c0f2f708-567b-4a10-9b28-efad7d156283"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	threadResponse := nylas.GetThread("18df121c8f5b46df", grantId)

	fmt.Println(threadResponse.Data.MessageIDs)
}

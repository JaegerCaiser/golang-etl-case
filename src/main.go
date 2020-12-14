package main

import (
	"etl-neoway-challenge/src/router"
	"log"
)

func main() {
	var r = router.StartRouter()
	var err = r.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

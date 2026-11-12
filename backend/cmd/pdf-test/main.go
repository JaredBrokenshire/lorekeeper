package main

import (
	"fmt"
	"log"
	"lorekeeper/services"
)

func main() {
	docService := services.NewDocumentService()

	pages, err := docService.Extract("assets/test.pdf")
	if err != nil {
		log.Fatal(err)
	}

	for _, page := range pages {
		fmt.Printf("Page %d\n", page.Number)
		fmt.Println(page.Text)
		fmt.Println("-----------------")
	}
}

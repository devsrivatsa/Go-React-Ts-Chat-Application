package cmd

import (
	"log"

	"github.com/devsrivatsa/chat_app_go-ts-react/db"
)

func main() {
	database, err := db.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()
}

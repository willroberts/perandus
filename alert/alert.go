package alert

import (
	"log"

	"github.com/willroberts/perandus/models"
)

func Alert(i models.Item, s models.Stash) {
	log.Println("ALERT! Found matching item:")
	log.Println("\tName:", i.Name)
	log.Println("\tPrice:", i.Note)
	log.Println("\tSeller:", s.LastCharacterName)
	log.Println("\tID:", i.ID)
}
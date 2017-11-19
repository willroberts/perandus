package alert

import (
	"log"

	"github.com/willroberts/perandus/models"
)

// ConsoleLogAlert writes alert information to the terminal window.
func ConsoleLogAlert(i models.Item, s models.Stash) {
	log.Println("ALERT! Found matching item:")
	log.Println("\tName:", i.Name)
	log.Println("\tPrice:", i.Note)
	log.Println("\tSeller:", s.LastCharacterName)
	log.Println("\tID:", i.ID)
}

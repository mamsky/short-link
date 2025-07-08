package service

import (
	"log"
	"short-link/model"
	"short-link/repository"
)

func CreateLink(d *model.ShortLink) error {
	query := `INSERT INTO short_link (url, short_url) VALUES ($1, $2)`
	
	_, err := repository.DB.Exec(query, d.URL, d.ShortURL)
	if err != nil{
		log.Println("Error inserting short link:", err)
		return err;
	}
	log.Printf("✅ Short link inserted successfully: %s → %s\n", d.URL, d.ShortURL)
	return nil
}

func FindByID() {

}
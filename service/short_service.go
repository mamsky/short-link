package service

import (
	"database/sql"
	"fmt"
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

func FindByCode(code string) ( error,string) {
	var url string
	fmt.Println(url, code)
	query := `SELECT url FROM short_link WHERE short_url = $1`

	err := repository.DB.QueryRow(query, code).Scan(&url)
	if err != nil {
		if err == sql.ErrNoRows {
			return  fmt.Errorf("short code not found"),""
		}
		return  err,""
	}
	return  nil ,url
}

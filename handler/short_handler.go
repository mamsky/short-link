package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"short-link/model"
	"short-link/service"
	"short-link/utils"
	"strings"

	"github.com/joho/godotenv"
)

func HandleLongUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost{
		body, _ := io.ReadAll(r.Body)
		fmt.Println("RAW Request body:", string(body)) // 👈 log isi body
		r.Body = io.NopCloser(bytes.NewBuffer(body)) // reset body
		var url model.ShortLink
		err := json.NewDecoder(r.Body).Decode(&url)
		if err != nil{
			fmt.Println("Decode error:", err) // 👈 log error detail
			http.Error(w, "Invalid Request Body", http.StatusBadRequest)
			return;
		}

		urlCode := utils.GenerateCode(4)

		fieldData := &model.ShortLink{
			URL: url.URL,
			ShortURL: urlCode,
		}
		
		
		err = godotenv.Load()
		if err != nil{
			fmt.Println("Error loading .env file", err)
			return
		}

		baseUrl := os.Getenv("BASE_URL")
		fieldUrl := baseUrl + urlCode
		
		err = service.CreateLink(fieldData)

		if err != nil {
			http.Error(w, "Failed to save URL", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
			"url":fieldUrl,
		})
	}
}

func HandleRedirect(w http.ResponseWriter, r *http.Request){
	if r.Method == http.MethodGet{
		code := strings.TrimPrefix(r.URL.Path, "/")

		err, url := service.FindByCode(code)
		if err != nil {
			http.Error(w, "Short link not found", http.StatusNotFound)
			return
		}
		
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	}
}
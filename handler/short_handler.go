package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"short-link/model"
	"short-link/service"
	"short-link/utils"

	"github.com/joho/godotenv"
)

func HandleLongUrl(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost{
		var url model.ShortLink
		err := json.NewDecoder(r.Body).Decode(&url)
		if err != nil{
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
		response :=map[string]interface{}{
			"url": fieldUrl,
		}

		err = service.CreateLink(fieldData)
		
		if err != nil {
			http.Error(w, "Failed to save URL", http.StatusInternalServerError)
			return
		}


		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

func HandleRedirect(w http.ResponseWriter, r *http.Request){

}
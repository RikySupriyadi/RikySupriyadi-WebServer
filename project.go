package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Struct untuk response pertama
type judulAnimeRomanceTerbaikTerbaruResponse struct {
	Message string `json:"message"`
}

// Struct untuk response kedua
type seputarInfoResponse struct {
	Data []string `json:"data"`
}

func main() {
	// Menentukan handler untuk setiap endpoint
	http.HandleFunc("/judul-anime", judulAnimeRomanceTerbaikTerbaruHandler)
	http.HandleFunc("/info-anime", seputarInfoHandler)

	// Menjalankan server pada port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler untuk endpoint "/judulAnimeRomanceTerbaruterbaikyangSayaTonton"
func judulAnimeRomanceTerbaikTerbaruHandler(w http.ResponseWriter, r *http.Request) {
	// Membuat response
	response := judulAnimeRomanceTerbaikTerbaruResponse{
		Message: "THE ANGEL NEXT DOOR SPOILS ME ROTTEN",
	}

	// Mengubah response menjadi JSON dan mengirimkannya
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handler untuk endpoint "/seputarInfoTerkaitAnimenya"
func seputarInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Membuat response
	response := seputarInfoResponse{
		Data: []string{
			"Judul Lain: Otonari no Tenshi-sama ni Itsunomanika Dame Ningen ni Sareteita Ken",
			"Episode: 13 episode (S1)",
			"Durasi: 23 menit",
			"Tayang: Crunchyroll, Aniplus, Prime Video",
		},
	}

	// Mengubah response menjadi JSON dan mengirimkannya
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

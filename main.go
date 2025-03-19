package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Product repräsentiert ein Produkt mit den erforderlichen Informationen
type Product struct {
	Produktname   string  `json:"produktname"`
	ProduktID     int     `json:"produkt_id"`
	Preis         float64 `json:"preis"`
	Verfugbarkeit string  `json:"verfugbarkeit"`
	Versanddauer  string  `json:"versanddauer"`
}

// Handler für den Endpunkt /product/{id}/stocklevel
func productStockLevelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Ungültige Produkt-ID", http.StatusBadRequest)
		return
	}

	// Simuliere das Abrufen von Daten
	log.Println("Fetching stock level from ERP....")
	time.Sleep(1 * time.Second)
	log.Println("Estimate shipping time with ERP RPC....")

	// Erstelle ein Demo-Produkt
	product := Product{
		Produktname:   "Demo Produkt",
		ProduktID:     id,
		Preis:         99.99,
		Verfugbarkeit: "Auf Lager",
		Versanddauer:  "3-5 Tage",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// Handler für den Endpunkt /account/{id}/checkLimit
func accountCheckLimitHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["id"]
	amount := r.URL.Query().Get("amount") // Beispiel: ?amount=209,99

	// Log-Ausgaben zur Simulation der Datenabfrage
	log.Println("Getting account information..")
	log.Printf("Getting unpaid orders for account %s\n", accountId)
	log.Println("Checking limit for new order...")

	w.Header().Set("Content-Type", "application/json")
	amountInt, _ := strconv.Atoi(amount)
	if amountInt > 999 {
		// Rückgabe einer Response mit Status 200 und einer Fehlermeldung im Body
		response := map[string]string{
			"error":  "Order exceeds account limit",
			"amount": amount, // Optional: Ausgabe des angefragten Betrags
		}
		json.NewEncoder(w).Encode(response)
	} else {
		// Rückgabe einer Response mit Status 200 und einer Fehlermeldung im Body
		response := map[string]string{
			"success": "Order within account limit",
			"amount":  amount, // Optional: Ausgabe des angefragten Betrags
		}
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	// Router initialisieren
	r := mux.NewRouter()
	r.HandleFunc("/product/{id}/stocklevel", productStockLevelHandler).Methods("GET")
	r.HandleFunc("/account/{id}/checkLimit", accountCheckLimitHandler).Methods("GET")

	// Server starten
	fmt.Println("Server startet auf Port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

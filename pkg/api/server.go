package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AmbRew2606/server_monitor/pkg/monitor"
)

func StartServer() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		metrics, err := monitor.GetMetrics()
		if err != nil {
			http.Error(w, "Failed to get metrics", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(metrics)
	})

	log.Println("✅ API server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

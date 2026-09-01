package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/stanleyHayes/ghanaschools/internal/catalog"
)

type server struct {
	dataset       *catalog.Dataset
	allowedOrigin string
}

func main() {
	dataPath := os.Getenv("SCHOOLS_DATA_PATH")
	if dataPath == "" {
		dataPath = "data/schools.json"
	}
	raw, err := os.ReadFile(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	dataset, err := catalog.Parse(raw)
	if err != nil {
		log.Fatal(err)
	}
	origin := os.Getenv("ALLOWED_ORIGIN")
	if origin == "" {
		origin = "https://schools.digitalghana.dev"
	}
	s := &server{dataset: dataset, allowedOrigin: origin}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", s.health)
	mux.HandleFunc("GET /v1/schools", s.schools)
	mux.HandleFunc("GET /v1/schools/{id}", s.school)
	mux.HandleFunc("POST /graphql", s.graphql)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("GhanaSchools API listening on :%s dataset=%s", port, dataset.DatasetVersion)
	log.Fatal(http.ListenAndServe(":"+port, s.middleware(mux)))
}

func (s *server) middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "no-referrer")
		w.Header().Set("Cache-Control", "public, max-age=60")
		if origin := r.Header.Get("Origin"); origin != "" && origin == s.allowedOrigin {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
		}
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
func (s *server) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, 200, map[string]any{"status": "ok", "datasetVersion": s.dataset.DatasetVersion, "coverage": s.dataset.Coverage, "schools": len(s.dataset.Schools), "checkedAt": time.Now().UTC().Format(time.RFC3339)})
}
func (s *server) schools(w http.ResponseWriter, r *http.Request) {
	results := s.dataset.Search(catalog.Query{Search: r.URL.Query().Get("q"), Region: r.URL.Query().Get("region"), Type: r.URL.Query().Get("type")})
	writeJSON(w, 200, map[string]any{"data": results, "count": len(results), "datasetVersion": s.dataset.DatasetVersion, "coverage": s.dataset.Coverage})
}
func (s *server) school(w http.ResponseWriter, r *http.Request) {
	record, ok := s.dataset.Get(r.PathValue("id"))
	if !ok {
		writeJSON(w, 404, map[string]string{"code": "SCHOOL_NOT_FOUND", "message": "No school exists for this stable ID in the current subset."})
		return
	}
	writeJSON(w, 200, map[string]any{"data": record, "datasetVersion": s.dataset.DatasetVersion})
}

func (s *server) graphql(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Query     string         `json:"query"`
		Variables map[string]any `json:"variables"`
	}
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 32<<10)).Decode(&body); err != nil {
		writeJSON(w, 400, map[string]any{"errors": []map[string]string{{"message": "Invalid JSON request"}}})
		return
	}
	if strings.Contains(body.Query, "schools") {
		q, _ := body.Variables["q"].(string)
		region, _ := body.Variables["region"].(string)
		results := s.dataset.Search(catalog.Query{Search: q, Region: region})
		writeJSON(w, 200, map[string]any{"data": map[string]any{"schools": results}, "extensions": map[string]string{"datasetVersion": s.dataset.DatasetVersion}})
		return
	}
	writeJSON(w, 400, map[string]any{"errors": []map[string]string{{"message": "Only the documented schools query is supported in beta."}}})
}

package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stanleyHayes/ghanaschools/internal/catalog"
)

func testServer(t *testing.T) *server {
	t.Helper()
	raw, err := os.ReadFile("../../data/schools.json")
	if err != nil {
		t.Fatal(err)
	}
	dataset, err := catalog.Parse(raw)
	if err != nil {
		t.Fatal(err)
	}
	return &server{dataset: dataset, allowedOrigin: "https://schools.digitalghana.dev"}
}

func TestRESTSearchAndCORS(t *testing.T) {
	s := testServer(t)
	request := httptest.NewRequest(http.MethodGet, "/v1/schools?q=KNUST", nil)
	request.Header.Set("Origin", "https://schools.digitalghana.dev")
	response := httptest.NewRecorder()
	s.middleware(http.HandlerFunc(s.schools)).ServeHTTP(response, request)
	if response.Code != 200 || !strings.Contains(response.Body.String(), `"id":"knust"`) {
		t.Fatalf("status=%d body=%s", response.Code, response.Body.String())
	}
	if got := response.Header().Get("Access-Control-Allow-Origin"); got != "https://schools.digitalghana.dev" {
		t.Fatalf("cors=%q", got)
	}
}

func TestGraphQLParity(t *testing.T) {
	s := testServer(t)
	request := httptest.NewRequest(http.MethodPost, "/graphql", strings.NewReader(`{"query":"query($q:String){ schools(q:$q){ id name } }","variables":{"q":"KNUST"}}`))
	response := httptest.NewRecorder()
	s.graphql(response, request)
	if response.Code != 200 || !strings.Contains(response.Body.String(), `"id":"knust"`) {
		t.Fatalf("status=%d body=%s", response.Code, response.Body.String())
	}
}

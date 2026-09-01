package catalog

import (
	"os"
	"testing"
)

func fixture(t *testing.T) *Dataset {
	t.Helper()
	raw, err := os.ReadFile("../../data/schools.json")
	if err != nil {
		t.Fatal(err)
	}
	dataset, err := Parse(raw)
	if err != nil {
		t.Fatal(err)
	}
	return dataset
}

func TestDatasetInvariants(t *testing.T) {
	dataset := fixture(t)
	if got := len(dataset.Schools); got != 16 {
		t.Fatalf("schools=%d want 16", got)
	}
	for _, school := range dataset.Schools {
		if school.Type != "PUBLIC_UNIVERSITY" || school.Ownership != "PUBLIC" {
			t.Fatalf("unexpected beta scope record: %+v", school)
		}
		if len(school.Sources) != 2 {
			t.Fatalf("%s must expose field provenance", school.ID)
		}
	}
}

func TestAliasesAndFilters(t *testing.T) {
	dataset := fixture(t)
	if got := dataset.Search(Query{Search: "KNUST"}); len(got) != 1 || got[0].ID != "knust" {
		t.Fatalf("alias result=%+v", got)
	}
	if got := dataset.Search(Query{Search: "Ghana Institute of Journalism"}); len(got) != 1 || got[0].ID != "unimac" {
		t.Fatalf("historical alias result=%+v", got)
	}
	if got := dataset.Search(Query{Region: "Central"}); len(got) != 2 {
		t.Fatalf("central count=%d want 2", len(got))
	}
}

func TestStableIDsDoNotDependOnNames(t *testing.T) {
	dataset := fixture(t)
	school, ok := dataset.Get("ug")
	if !ok {
		t.Fatal("ug missing")
	}
	school.Name = "Renamed display value"
	if school.ID != "ug" {
		t.Fatal("stable id changed with name")
	}
}

func TestDuplicateNamesRemainDistinct(t *testing.T) {
	raw := []byte(`{"datasetVersion":"x","coverage":"x","reviewedAt":"2026-09-01","ghanaGeoDatasetVersion":"x","schools":[{"id":"one","name":"Same","aliases":[],"type":"PUBLIC_UNIVERSITY","ownership":"PUBLIC","locality":"Accra","region":"Greater Accra","ghanaGeoPlaceId":"a","sources":["s"]},{"id":"two","name":"Same","aliases":[],"type":"PUBLIC_UNIVERSITY","ownership":"PUBLIC","locality":"Accra","region":"Greater Accra","ghanaGeoPlaceId":"a","sources":["s"]}]}`)
	dataset, err := Parse(raw)
	if err != nil {
		t.Fatal(err)
	}
	if got := dataset.Search(Query{Search: "same"}); len(got) != 2 {
		t.Fatalf("duplicates collapsed: %+v", got)
	}
}

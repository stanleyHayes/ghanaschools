package catalog

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type School struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Aliases         []string `json:"aliases"`
	Type            string   `json:"type"`
	Ownership       string   `json:"ownership"`
	Locality        string   `json:"locality"`
	Region          string   `json:"region"`
	GhanaGeoPlaceID string   `json:"ghanaGeoPlaceId"`
	LocationNote    string   `json:"locationNote,omitempty"`
	Sources         []string `json:"sources"`
}

type Dataset struct {
	DatasetVersion         string   `json:"datasetVersion"`
	Coverage               string   `json:"coverage"`
	ReviewedAt             string   `json:"reviewedAt"`
	GhanaGeoDatasetVersion string   `json:"ghanaGeoDatasetVersion"`
	Schools                []School `json:"schools"`
}

type Query struct{ Search, Region, Type string }

func Parse(raw []byte) (*Dataset, error) {
	var dataset Dataset
	if err := json.Unmarshal(raw, &dataset); err != nil {
		return nil, fmt.Errorf("decode dataset: %w", err)
	}
	if dataset.DatasetVersion == "" || dataset.Coverage == "" || dataset.ReviewedAt == "" || dataset.GhanaGeoDatasetVersion == "" {
		return nil, errors.New("dataset metadata is incomplete")
	}
	if len(dataset.Schools) == 0 {
		return nil, errors.New("dataset contains no schools")
	}
	seen := map[string]bool{}
	for _, school := range dataset.Schools {
		if school.ID == "" || school.Name == "" || school.Type == "" || school.Ownership == "" || school.Locality == "" || school.Region == "" || school.GhanaGeoPlaceID == "" || len(school.Sources) == 0 {
			return nil, fmt.Errorf("school %q has incomplete required fields", school.ID)
		}
		if seen[school.ID] {
			return nil, fmt.Errorf("duplicate school id %q", school.ID)
		}
		seen[school.ID] = true
	}
	return &dataset, nil
}

func normalize(value string) string {
	return strings.Join(strings.FieldsFunc(strings.ToLower(value), func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) }), " ")
}

func (d *Dataset) Search(query Query) []School {
	needle, region, kind := normalize(query.Search), normalize(query.Region), normalize(query.Type)
	result := make([]School, 0)
	for _, school := range d.Schools {
		if region != "" && normalize(school.Region) != region {
			continue
		}
		if kind != "" && normalize(school.Type) != kind {
			continue
		}
		if needle != "" {
			matched := strings.Contains(normalize(school.Name), needle) || strings.Contains(normalize(school.ID), needle)
			for _, alias := range school.Aliases {
				matched = matched || strings.Contains(normalize(alias), needle)
			}
			if !matched {
				continue
			}
		}
		result = append(result, school)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })
	return result
}

func (d *Dataset) Get(id string) (School, bool) {
	for _, school := range d.Schools {
		if school.ID == id {
			return school, true
		}
	}
	return School{}, false
}

package graph

import (
    "context"
    "encoding/json"
    "io/ioutil"
    "strings"
	"log"
	"fmt"
    "study-gator-backend/graph/model"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

func (r *queryResolver) SearchByClassTitle(ctx context.Context, term string) ([]*model.Resource, error) {
    titleLower := strings.ToLower(term)
	
	log.Printf("SearchByClassTitle query received: %s", term)

	// Read from the JSON file 
	bytes, err := ioutil.ReadFile("graph/resources.json") 
	if err != nil {
		log.Printf("Error reading resources.json file: %s", err)
		return nil, fmt.Errorf("error reading from resources file: %w", err)
	}

	// Deserialize the JSON data into the slice of Resource
	var resources []*model.Resource
	if err := json.Unmarshal(bytes, &resources); err != nil {
		log.Printf("Error unmarshaling resources.json: %s", err)
		return nil, fmt.Errorf("error unmarshaling resources data: %w", err)
	}

	// Filter resources by the search title
	var results []*model.Resource
	for _, resource := range resources {
		if resource.ClassTitle != nil && strings.Contains(strings.ToLower(*resource.ClassTitle), titleLower) {
			results = append(results, resource)
		}
	}

	// If no resources found, log this as well
	if len(results) == 0 {
		log.Printf("No resources found for the title: %s", term)
	} else {
		// Log the number of results found
		log.Printf("Found %d resources for the title: %s", len(results), term)
	}

	return results, nil
}

func stringPointer(s string) *string {
    return &s
}
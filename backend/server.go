package main

import (
	"net/http"
	"os"
	"study-gator-backend/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	// Add CORS middleware around every request
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173", "http://localhost:8080", "https://studygator.chasemacdonnell.net"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	// Serve static files from the images directory
	fs := http.FileServer(http.Dir("./graph/images"))
	router.Handle("/images/*", http.StripPrefix("/images/", fs))

	// Setup GraphQL server
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// Add GraphQL playground
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))

	// Add GraphQL endpoint
	router.Handle("/query", srv)

	// Start the server
	http.ListenAndServe(":"+port, router)
}
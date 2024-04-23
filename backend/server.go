package main

import (
	"net/http"
	"os"
	"study-gator-backend/graph"
	"study-gator-backend/graph/gqlcontext"
	"time"
	"context"
	"io"
	"path/filepath"
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-pkgz/auth"
	"github.com/go-pkgz/auth/avatar"
	"github.com/go-pkgz/auth/token"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/upload"

)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	SECRET, exists := os.LookupEnv("SG_SECRET")
	URL := "https://studygator-api.chasemacdonnell.net"
	CID, cid_exists := os.LookupEnv("SG_GITHUB_CID")
	CSECRET, csecret_exists := os.LookupEnv("SG_GITHUB_CSECRET")
	ORIGINS := []string{"https://studygator.chasemacdonnell.net", "https://studygator-api.chasemacdonnell.net"}
	COOKIE_DOMAIN := "chasemacdonnell.net"
	SAMESITE := http.SameSiteNoneMode
	if !exists {
		SECRET = "dev-secret"
		URL = "http://localhost:8080"
		COOKIE_DOMAIN = "localhost"
		SAMESITE = http.SameSiteDefaultMode
		ORIGINS = []string{"http://localhost:5173", "http://localhost:8080", "https://studygator.chasemacdonnell.net", "https://studygator-api.chasemacdonnell.net"}
	}

	// This CID and CSECRET are for a localhost only zero permission OAuth2 provider with Github, should be no problems with including this secret in the repo.
	// The generated tokens have zero access to your github account and are only used for auth on this platform.
	if !cid_exists {
		CID = "1f3b435de0687155aa4f"
	}

	if !csecret_exists {
		CSECRET = "dbb30cf61105d17f12324ebdb9770be4f0f49553"
	}

	println("URL: ", URL)
	println("SECRET: ", SECRET)
	println("CID: ", CID)
	println("CSECRET: ", CSECRET)

	// Auth Options
	options := auth.Opts{
		SecretReader: token.SecretFunc(func(id string) (string, error) { // secret key for JWT
			return SECRET, nil
		}),
		TokenDuration:   time.Minute * 5, // token expires in 5 minutes
		CookieDuration:  time.Hour * 24,  // cookie expires in 1 day and will enforce re-login
		Issuer:          "studygator",
		URL:             URL,
		JWTCookieDomain: COOKIE_DOMAIN,
		SameSiteCookie:  SAMESITE,
		SecureCookies:   exists,
		AvatarStore:     avatar.NewLocalFS("./avatars"),
		// GraphQL by nature only uses POST requests, eliminating most XSRF vulneabilities.,
		// Plus XSRF tokens are a quite controversial topic, and arguably dont work.
		DisableXSRF: true,
		Validator: token.ValidatorFunc(func(_ string, claims token.Claims) bool {
			return true
		}),
	}

	service := auth.NewService(options)
	service.AddProvider("github", CID, CSECRET) // add github provider

	m := service.Middleware()

	router := chi.NewRouter()
	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   ORIGINS,
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				if !exists && r.Host == "localhost" {
					return true
				}
				return r.Host == "studygator.chasemacdonnell.net" || r.Host == "studygator-api.chasemacdonnell.net"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/", playground.Handler("StudyGator API", "/query"))
	router.Handle("/query", m.Auth(gqlcontext.AuthMiddleware()(srv)))

	authRoutes, avaRoutes := service.Handlers()

	router.Mount("/auth", authRoutes)  // add auth handlers
	router.Mount("/avatar", avaRoutes) // add avatar handler

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}

	// added under here (file upload)
	r := chi.NewRouter()

	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(upload.Transport{})

	srv.Use(extension.Introspection{})

	// Handle HTTP file upload requests
	r.Post("/upload", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(10 << 20) // 10 MB limit
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		// Save the file locally
		f, err := os.OpenFile(filepath.Join("./uploads", handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		w.Write([]byte("Uploaded successfully\n"))
	})

	// Handle GraphQL requests
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	// Start HTTP server
	http.ListenAndServe(":8080", r)
}

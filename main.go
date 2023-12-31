package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/julianinsua/RSSAgregator/internal/database"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("Unspecified PORT in .env file")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("unspecified databse conection url in .env file")
	}

	conn, e := sql.Open("postgres", dbURL)
	if e != nil {
		log.Fatal("unable to connect to the database", e)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, 12*time.Hour)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*s"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Post("/user", apiCfg.handleCreateUser)
	v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handleGetUserByApiKey))
	v1Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handleAddFeed))
	v1Router.Get("/feeds", apiCfg.handleGetAllFeeds)
	v1Router.Post("/feed/follow", apiCfg.middlewareAuth(apiCfg.handleCreateFeedFollow))
	v1Router.Get("/feed/follow", apiCfg.middlewareAuth(apiCfg.handleListFeedFollows))
	v1Router.Delete("/feed/follow/{ffID}", apiCfg.middlewareAuth(apiCfg.handleUnfollowFeed))
	v1Router.Get("/post", apiCfg.middlewareAuth(apiCfg.handleGetUserPosts))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port: %v", portString)
	log.Println("It's a beautifull day in the server")
	e = srv.ListenAndServe()
	if e != nil {
		log.Fatal(e)
	}
}

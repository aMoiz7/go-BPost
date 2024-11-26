package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	
)

func main(){
       
	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == ""{
		log.Fatal("port not found")
	}
   	router := chi.NewRouter()
	

   router.Use(cors.AllowAll().Handler)

   server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
   log.Printf("server running on %v "  , portString)
routerv1 := chi.NewRouter()

	// Health check endpoint
	routerv1.Get("/healthz", handlerreadiness)

	// Mount v1 router under /v1 path
	router.Mount("/v1", routerv1)


    routerv1.Get("/err",errorHandler  )

   err:= server.ListenAndServe()

   if err!=nil{
	log.Fatal(err)
   }
   


}
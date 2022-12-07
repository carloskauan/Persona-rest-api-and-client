package routes

import(
  "net/http"
  "api/controllers"
  "log"
  "github.com/gorilla/mux"
  "github.com/gorilla/handlers"
  "api/middleware"
)

func HandleRequest() {
  rmux := mux.NewRouter()
  rmux.Use(middleware.ContentTypeSet)
  rmux.HandleFunc("/api/persona", controllers.AllPersonas).Methods("Get")
  rmux.HandleFunc("/api/persona/{id}", controllers.OnePersona).Methods("Get")
  rmux.HandleFunc("/api/persona", controllers.Register).Methods("Post")
  rmux.HandleFunc("/api/persona/{id}", controllers.Delete).Methods("Delete")
  rmux.HandleFunc("/api/persona/{id}", controllers.Edit).Methods("Put")
  log.Fatal(http.ListenAndServe(":7070", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(rmux)))
}

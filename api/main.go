package main

import(
  "fmt"
  "api/routes"
  "api/database"
)

func main() {
  database.ConnectDB()

  fmt.Println("Iniciando res-api com go")
  routes.HandleRequest()
}

package database

import(
  "log"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)
 var(
   DB *gorm.DB
   err error
 )

 func ConnectDB(){
   connector := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
   DB, err = gorm.Open(postgres.Open(connector))
   if err != nil{
     log.Panic("Error ao conectar com o banco de dados")
   }
 }

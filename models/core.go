package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "fmt"
  "belajar-ech0-framework/config"
  "os"
)

type Models struct {
  Db *gorm.DB
}

func (m Models) Connect(db_connection, db_string string) *gorm.DB {
  config.LoadConfig();
  DB_CONNECTION := os.Getenv("DB_CONNECTION")
  DB_STRING := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")" + "/" + os.Getenv("DB_DATABASE")
  
  m.Db, err := gorm.Open(DB_CONNECTION, DB_DATABASE)
  if err != nil {
    fmt.Println("Error : ", err.Error());
    return
  }
  
  return m.Db
}

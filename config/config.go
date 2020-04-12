package config

import (
  "github.com/joho/godotenv"
  "os"
  "fmt"
  "gopkg.in/alecthomas/kingpin.v2"
)

func LoadConfig() {
  flag := kingpin.New("Aplikasi CRUD", "Aplikasi CRUD dengan echo framework dan golang")
  
  // menagkap flag --env=".env"
  flagEnv := flag.Flag("env", "Lokasi file .env (environment) atau file custom .env").Default(".env").String() // required = flag wajib ada, string = isi flag harus string pakai kutip
  kingpin.MustParse(flag.Parse(os.Args[1:]))
  
  // ambil file env yang berisi konfigurasi
  err := godotenv.Load(*flagEnv)
  if err != nil {
    // jika file .env atau semacamnya tidak ditemukan, stop jalan aplikasi
    fmt.Println("Error pada aplikasi!")
    fmt.Println("File environment " + *flagEnv + " tidak ditemukan!")
    return
  }
}

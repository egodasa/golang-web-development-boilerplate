package config

import (
  godotenv "github.com/joho/godotenv"
  os "os"
  fmt "fmt"
  kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func LoadConfig() {
  var flag kingpin *kingpin.Application = kingpin.New("Aplikasi CRUD", "Aplikasi CRUD dengan echo framework dan golang")
  
  // menagkap flag --env=".env"
  var flagEnv *kingpin.FlagCause = flag.Flag("env", "Lokasi file .env (environment) atau file custom .env").Default(".env").String() // required = flag wajib ada, string = isi flag harus string pakai kutip
  kingpin.MustParse(flag.Parse(os.Args[1:]))
  
  // ambil file env yang berisi konfigurasi
  var err error = godotenv.Load(*flagEnv)
  if err != nil {
    // jika file .env atau semacamnya tidak ditemukan, stop jalan aplikasi
    fmt.Println("Error pada aplikasi!")
    fmt.Println("File environment " + *flagEnv + " tidak ditemukan!")
    return
  }
}

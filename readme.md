# Contoh pemakaian echo framework (go language)
Struktur file/folder dibuat dengan konsep MVC

# Library yang digunakan pada project ini
1. Echo Framework (https://github.com/labstack/echo), merupakan framework golang
1. Godotenv (https://github.com/joho/godotenv), untuk membaca isi file .env
1. Kingpin (https://github.com/alecthomas/kingpin), untuk membaca argument/flag pada CLI
1. GORM (https://github.com/jinzhu/gorm), untuk ORM 

# Struktur file/folder
| File/Folder | Kegunaan |
| ------ | ------ |
| `main.go` | Program akan dijalankan dari file ini. Semua controller dan router akan dipanggil disini |
| `api/` (WIP) | Folder yang berisi controller khusus RESTFul API |
| `controller/` | Folder yang berisi controller atau proses halaman |
| `controller/core.go` | File yang berisi struct yang akan digunakan untuk semua controller. Semacam variabel/method global untuk semua controller |
| `models/` (WIP) | Folder yang berisi model. Model menggunakan library ***GORM*** |
| `views/` (WIP) | Folder yang berisi file views format .HTML |
| `.env` | File yang berisi pengaturan aplikasi |

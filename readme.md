# Contoh pemakaian echo framework (go language)
Struktur file/folder dibuat dengan konsep MVC

# Library yang digunakan pada project ini
1. Gin Framework (github.com/gin-gonic/gin), merupakan framework golang
1. Godotenv (https://github.com/joho/godotenv), untuk membaca isi file .env
1. Kingpin (https://github.com/alecthomas/kingpin), untuk membaca argument/flag pada CLI
1. Beego (https://github.com/astaxie/beego/), untuk ORM (Hanya ORM nya saja. Karena Beego adalah framework full)
1. Squirrel (https://github.com/Masterminds/squirrel), untuk Query Builder 

# Struktur file/folder
| File/Folder | Kegunaan |
| ------ | ------ |
| `main.go` | Program akan dijalankan dari file ini. Semua controller dan router akan dipanggil disini |
| `api/` | Folder yang berisi controller khusus RESTFul API |
| `controller/` | Folder yang berisi controller atau proses halaman |
| `controller/core.go` | File yang berisi struct yang akan digunakan untuk semua controller. Semacam variabel/method global untuk semua controller |
| `models/` | Folder yang berisi model. |
| `views/` (WIP) | Folder yang berisi file views format .HTML |
| `websockets/` (WIP) | Folder yang berisi method untuk websockets. Agar nanti bisa digunakan di controller dll |
| `.env` | File yang berisi pengaturan aplikasi |

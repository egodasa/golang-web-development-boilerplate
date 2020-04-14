package models

var MobilColumn = []Column{{
    Name: "id_mobil",
    Type: "int",
    Fillable: false,
    IsPk: true,
    AutoIncrement: true,
  },
  {
    Name: "kode_mobil",
    Type: "varchar",
    Fillable: true,
  },
  {
    Name: "merk",
    Type: "varchar",
    Fillable: true,
  },  
  {
    Name: "tipe",
    Type: "varchar",
    Fillable: true,
  },  
  {
    Name: "harga",
    Type: "int",
    Fillable: true,
  },  
  {
    Name: "warna",
    Type: "varchar",
    Fillable: true,
  },  
  {
    Name: "penggerak",
    Type: "varchar",
    Fillable: true,
  },  
  {
    Name: "banyak_roda",
    Type: "int",
    Fillable: true,
  },  
  {
    Name: "banyak_bangku",
    Type: "int",
    Fillable: true,
  },  
  {
    Name: "jenis_mesin",
    Type: "varchar",
    Fillable: true,
  },  
  {
    Name: "mesin",
    Type: "varchar",
    Fillable: true,
  },
}

// inisaliasi model mobil
// nanti variabel ini akan digunakan di controller
var ModelMobil = &Mobil{Models{
 tableName: "tb_mobil",
 ColumnList: MobilColumn,
}};

// struct models digabung dengan struct mobil
// agar kita bisa menambahkan custom methos selain method dasar CRUD ke struct mobil
// kurang lebih seperti pewarisan
// dimana struct mobil mendapatkan warisan berupa struct models
// dan struct mobil bisa dimodifikasi methodnya
type Mobil struct {
  Models
}

// contoh custom method dari mobil
// dimana method ini tidak ada distruct models
func (m *Mobil) CariMobil(kataKunci string) string {
  return "Pencarian"
}


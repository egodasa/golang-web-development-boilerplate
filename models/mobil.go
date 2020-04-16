package models

var MobilColumn = []Column{{
    Name: "id_mobil",
    Type: "int",
    Fillable: false,
    IsPk: true,
    AutoIncrement: true,
  },
  {
    Name: "id_perusahaan",
    Type: "int",
    Fillable: true,
  },
  {
    Name: "nm_mobil",
    Type: "varchar",
    Fillable: true,
  },  
  {
    Name: "jenis_penggerak",
    Type: "varchar",
    Fillable: true,
  },  
  {
    Name: "banyak_roda",
    Type: "int",
    Fillable: true,
  },  
  {
    Name: "id_jenis",
    Type: "int",
    Fillable: true,
  },  
  {
    Name: "harga",
    Type: "int",
    Fillable: true,
  },  
}

// inisaliasi model mobil
// nanti variabel ini akan digunakan di controller
var ModelMobil = &Models{
 tableName: "tb_mobil",
 ColumnList: MobilColumn,
};

// struct models digabung dengan struct mobil
// agar kita bisa menambahkan custom methos selain method dasar CRUD ke struct mobil
// kurang lebih seperti pewarisan
// dimana struct mobil mendapatkan warisan berupa struct models
// dan struct mobil bisa dimodifikasi methodnya
type Mobil Models

// contoh custom method dari mobil
// dimana method ini tidak ada distruct models
func (m *Mobil) CariMobil(kataKunci string) string {
  return "Pencarian"
}


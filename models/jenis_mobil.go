package models

var JenisMobilColumn = []Column{{
    Name: "id_jenis",
    Type: "int",
    Fillable: false,
    IsPk: true,
    AutoIncrement: true,
  },
  {
    Name: "nm_jenis",
    Type: "varhcar",
    Fillable: true,
  },
}

// inisaliasi model perusahaan
// nanti variabel ini akan digunakan di controller
var ModelJenisMobil = &Models{
 tableName: "tb_jenis_mobil",
 ColumnList: JenisMobilColumn,
};

// struct models digabung dengan struct perusahaan
// agar kita bisa menambahkan custom methos selain method dasar CRUD ke struct perusahaan
// kurang lebih seperti pewarisan
// dimana struct perusahaan mendapatkan warisan berupa struct models
// dan struct perusahaan bisa dimodifikasi methodnya
type JenisMobil Models



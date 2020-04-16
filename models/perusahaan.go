package models

var PerusahaanColumn = []Column{{
    Name: "id_perusahaan",
    Type: "int",
    Fillable: false,
    IsPk: true,
    AutoIncrement: true,
  },
  {
    Name: "nm_perusahaan",
    Type: "varhcar",
    Fillable: true,
  },
}

// inisaliasi model perusahaan
// nanti variabel ini akan digunakan di controller
var ModelPerusahaan = &Models{
 tableName: "tb_perusahaan",
 ColumnList: PerusahaanColumn,
};

// struct models digabung dengan struct perusahaan
// agar kita bisa menambahkan custom methos selain method dasar CRUD ke struct perusahaan
// kurang lebih seperti pewarisan
// dimana struct perusahaan mendapatkan warisan berupa struct models
// dan struct perusahaan bisa dimodifikasi methodnya
type Perusahaan Models

// contoh custom method dari perusahaan
// dimana method ini tidak ada distruct models
func (m *Perusahaan) CariPerusahaan(kataKunci string) string {
  return "Pencarian"
}


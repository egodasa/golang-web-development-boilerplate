package models

var PerusahaanColumn = map[string]Column{
  "id_perusahaan": Column{
    Name: "id_perusahaan",
    Type: "int",
    Fillable: false,
    IsPk: true,
    AutoIncrement: true,
  },
  "nm_perusahaan": Column{
    Name: "nm_perusahaan",
    Type: "varhcar",
    Fillable: true,
  },
}

// inisaliasi model perusahaan
// nanti variabel ini akan digunakan di controller
type Perusahaan struct {
  *Models
}

var ModelPerusahaan Perusahaan = Perusahaan{NewModels("tb_perusahaan", PerusahaanColumn)}

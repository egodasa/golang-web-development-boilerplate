package models

var JenisMobilColumn = map[string]Column{
  "id_jenis": Column{
    Name: "id_jenis",
    Type: "int",
    Fillable: false,
    IsPk: true,
    AutoIncrement: true,
  },
  "nm_jenis": Column{
    Name: "nm_jenis",
    Type: "varhcar",
    Fillable: true,
  },
}

// inisaliasi model perusahaan
// nanti variabel ini akan digunakan di controller
var ModelJenisMobil *Models = NewModels("tb_jenis_mobil", JenisMobilColumn, nil)

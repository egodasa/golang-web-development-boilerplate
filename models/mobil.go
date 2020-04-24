package models

var MobilColumn = map[string]Column{
  "id_mobil": Column{
    Name: "id_mobil",
    Type: "int",
    Fillable: false,
    IsPk: true,
    AutoIncrement: true,
  },
  "id_perusahaan": Column{
    Name: "id_perusahaan",
    Type: "int",
    Fillable: true,
  },
  "nm_mobil": Column{
    Name: "nm_mobil",
    Type: "varchar",
    Fillable: true,
  },
  "jenis_penggerak": Column{
    Name: "jenis_penggerak",
    Type: "varchar",
    Fillable: true,
  },
  "banyak_roda": Column{
    Name: "banyak_roda",
    Type: "int",
    Fillable: true,
  },
  "id_jenis": Column{
    Name: "id_jenis",
    Type: "int",
    Fillable: true,
  },
  "harga": Column{
    Name: "harga",
    Type: "int",
    Fillable: true,
  },
} 

// inisaliasi model mobil
// nanti variabel ini akan digunakan di controller
var ModelMobil *Models = NewModels("tb_mobil", MobilColumn, []IModels{ModelPerusahaan, ModelJenisMobil})


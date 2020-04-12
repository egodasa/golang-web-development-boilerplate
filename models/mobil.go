package models

type ModelMobil struct {
  ID int `json:"id_mobil"`
  Merk string `json:"merk"`
  Tipe string `json:"tipe"`
  Harga int `json:"harga"`
  Warna []string `json:"warna"`
  Penggerak string `json:"penggerak"`
  BanyakRoda int `json:"banyak_roda"`
  BanyakBangku int `json:"banyak_bangku"`
  JenisMesin string `json:"jenis_mesin"`
  Mesin string `json:"mesin"`
}

var dataMobil = []ModelMobil {{
    ID: 1,
    Merk: "Toyota",
    Tipe: "Kijang Innova Venturer",
    Harga: 657000000,
    Warna: []string{"Merah", "Hitam", "Putih", "Silver"},
    Penggerak: "FR",
    BanyakRoda: 4,
    BanyakBangku: 6,
    JenisMesin: "Diesel",
    Mesin: "D4D",
  },
  {
    ID: 2,
    Merk: "Suzuki",
    Tipe: "Ertiga GX",
    Harga: 210000000,
    Warna: []string{"Merah", "Hitam", "Putih", "Silver"},
    Penggerak: "FF",
    BanyakRoda: 4,
    BanyakBangku: 7,
    JenisMesin: "Bensin",
    Mesin: "I4",
  },
  {
    ID: 3,
    Merk: "Suzuki",
    Tipe: "Karimun Wagon R",
    Harga: 108000000,
    Warna: []string{"Merah", "Hitam", "Putih", "Silver", "Biru"},
    Penggerak: "FF",
    BanyakRoda: 4,
    BanyakBangku: 4,
    JenisMesin: "Bensin",
    Mesin: "I4",
  },
}

func (m ModelMobil) Get() []ModelMobil {
  data := dataMobil
  return data
}
func (m ModelMobil) Find(id int) interface{} {
  // lakukan pencarian mobil berdasarkan ID
  for index, value := range dataMobil {
    if id == value.ID {
      return dataMobil[index]
    }
  }
  return struct{}{}
}


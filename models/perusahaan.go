package models

import "reflect"

var PerusahaanColumn = []Column{
	{
		Name:          "id_perusahaan",
		Type:          reflect.Int,
		Fillable:      false,
		IsPk:          true,
		AutoIncrement: true,
	},
	{
		Name:     "nm_perusahaan",
		Type:     reflect.String,
		Fillable: true,
	},
}

// inisaliasi model perusahaan
// nanti variabel ini akan digunakan di controller
type Perusahaan struct {
	*Models
}

var ModelPerusahaan Perusahaan = Perusahaan{NewModels("tb_perusahaan", PerusahaanColumn)}

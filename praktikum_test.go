package _be2024

import (
	"testing"
	"fmt"
	"github.com/jul003/BE_Tb/model"
	"github.com/jul003/BE_Tb/module"
)

func TestInsertGadget(t *testing.T) {
	name := "Oppo"
	merk := "Reno 4"
	harga := 3500000.00
	spesifikasi := model.Spesifikasi{
		Prosesor:     "Snapdragon 750",
		RAM:          8,
		Storage:      128,
	}
	deskripsi := "RAM OPPO Reno 4 adalah 8GB. Spesifikasi RAM 8GB tersebut tergolong besar dan sudah cukup untuk mendukung pemakaian multitasking sehari-hari tanpa hambatan."
	insertedID, err := module.InsertGadget(module.MongoConn, "gadget2024", name, merk, harga, spesifikasi, deskripsi)
	if err != nil {
		t.Errorf("Error inserting data: %v", err)
	}
	fmt.Printf("Data berhasil disimpan dengan id %s", insertedID.Hex())
}

func TestGetAll(t *testing.T) {
	data := module.GetDataGadget(module.MongoConn, "gadget2024")
	fmt.Println(data)
}

package main

import (
	"fmt"

	moduleProject "github.com/ZainalAbiddin/moduleProject"
)

func main() {
	nama := "admin"
	var kegiatan = [4]string{
		"Beladiri",
		"Memanah",
		"English Club",
		"Komputer Club",
	}
	moduleProject.Perkenalan()
	moduleProject.Pengguna(nama)
	visi, misi := moduleProject.VisiMisi()
	fmt.Println(visi)
	fmt.Println(misi)
	moduleProject.Ekstrakulikuler(kegiatan)
	fmt.Println("")
	fmt.Println("tes repository versi 1.0.4 ")
}

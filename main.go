package main

import (
	"fmt"

	moduleProject "github.com/ZainalAbiddin/moduleProject/v2"
)

// struct
type Nilaiipa struct {
	fisika, biologi, kimia float64
}

type Nilaimurid struct {
	Namamurid   string
	kelas       int
	Rataannilai []Nilaiipa
}

func main() {
	nama := "admin"
	var kegiatan = [4]string{
		"Beladiri",
		"Memanah",
		"English Club",
		"Komputer Club",
	}

	moduleProject.Perkenalan()
	moduleProject.Verifikasi(nama, moduleProject.Ujinama)
	visi, misi := moduleProject.VisiMisi()
	fmt.Println("===================")
	fmt.Println("Visi MA MAQDIS : ", visi)
	fmt.Println("Misi MA MAQDIS : ", misi)
	fmt.Println("===================")
	moduleProject.Ekstrakulikuler(kegiatan)
	fmt.Println("===================")

	//method struct tanpa interface
	bernilai := Nilaimurid{
		Namamurid: "Kelas Sulaiman",
		kelas:     3,
		Rataannilai: []Nilaiipa{
			Nilaiipa{
				fisika:  75.50,
				biologi: 80.50,
				kimia:   60.05,
			},
			Nilaiipa{
				fisika:  60.35,
				biologi: 76.19,
				kimia:   50.55,
			},
			Nilaiipa{
				fisika:  90.50,
				biologi: 89.99,
				kimia:   82.88,
			},
		},
	}

	// anonymous function
	func(user string, Uji func(string) bool) {
		if Uji(user) {

			bernilai.Nilai()
		} else {
			fmt.Printf("gagal")
		}
	}(nama, moduleProject.Ujinama)

	// anonymous struct
	contactPerson := struct {
		ustadzHadi   string
		ustadzMiftah string
	}{
		"0857-8257-5108",
		"0857-7693-6338",
	}
	fmt.Println("Contact Person :")
	fmt.Printf("%+v\n", contactPerson)
	fmt.Println("")

	//interface
	media := []moduleProject.Mediasocial{moduleProject.Instagram{}, moduleProject.Youtube{}}
	for _, medsos := range media {

		fmt.Println("Media Sosial :", medsos.Media())
	}
	fmt.Println("===================")

}

func (n Nilaimurid) Nilai() {
	fmt.Println(n.Namamurid)
	fmt.Println("kelas :", n.kelas)
	for _, nilai := range n.Rataannilai {
		fmt.Println("===================")
		fmt.Println(nilai.fisika)
		fmt.Println(nilai.biologi)
		fmt.Println(nilai.kimia)
	}
	fmt.Println("===================")
}

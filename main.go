package main

import "fmt"

// Creational Pattern: Factory Method untuk membuat objek Dosen
func NewDosen(course string) *Dosen {
	return &Dosen{
		observers: make([]Observer, 0),
		course:    course,
	}
}

// Structural Pattern: Pola Observer
type Observer interface {
	Update(string)
}

// Interface Subjek
type Subject interface {
	Attach(Observer)
	Notify(string)
}

// Subjek Konkret
type Dosen struct {
	observers []Observer
	course    string
}

func (d *Dosen) Attach(observer Observer) {
	d.observers = append(d.observers, observer)
}

func (d *Dosen) Notify(message string) {
	for _, observer := range d.observers {
		observer.Update(message)
	}
}

// Behavioral Pattern: Pengamat Konkret
type Mahasiswa struct {
	name string
}

func (m *Mahasiswa) Update(message string) {
	fmt.Printf("Mahasiswa %s menerima: %s\n", m.name, message)
}

func main() {
	// Membuat objek Dosen menggunakan Factory Method
	dosen := NewDosen("Ilmu Komputer")

	// Mahasiswa sebagai observer
	mahasiswa1 := &Mahasiswa{name: "Faiz"}
	mahasiswa2 := &Mahasiswa{name: "Nathan"}

	// Mahasiswa melakukan subscribe pada Dosen
	dosen.Attach(mahasiswa1)
	dosen.Attach(mahasiswa2)

	// Dosen mengirimkan pembaruan kepada semua Mahasiswa
	dosen.Notify("Tugas baru telah diunggah.")

	// Output:
	// Mahasiswa Alice menerima: Tugas baru telah diunggah.
	// Mahasiswa Bob menerima: Tugas baru telah diunggah.
}

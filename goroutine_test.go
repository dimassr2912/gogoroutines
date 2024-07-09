package go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // Kalau tanpa goroutine, akan menunggu statement ini dikerjakan hingga selesai
	fmt.Println("Ups")

	// Kadang goroutine tidak akan dijalankan karena akan di kill (terlalu lama) sehingga butuh menunggu agar aplikasi tidak dihentikan dulu
	time.Sleep(1 * time.Second) // Menunggu 1 detik
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 100000; i++ { // Tidak berjalan berurutan , sesuai dengan cara kerja goroutine yang jika lama prosesnya, akan ambil goroutine yang lain dahulu, lalu kembali lagi ke goroutine yang lama
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}

/**
 * Tidak bisa jika function return value karena tidak bisa ditangkap
 */

package go_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string) // Membuat channel dengan tipe data string

	// Menutup channel
	defer close(channel)
	go func() {
		time.Sleep(2 * time.Second)
		// Mengirim data ke channel
		channel <- "Rama"
		fmt.Println("Selesai mengirim")
	}()

	// Menerima data
	data := <-channel // Jika baru di deklarasi
	fmt.Println(data)
	// fmt.Println(<-channel)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	// Mengirim data ke channel
	channel <- "Rama"
	fmt.Println("Selesai mengirim")
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	data := <-channel // Jika baru di deklarasi
	fmt.Println(data)
	// fmt.Println(<-channel)
	time.Sleep(5 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	// Mengirim data ke channel
	channel <- "Rama"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)

}

Goroutine: Sebuah thread ringan yang dikelola Go Runtime
    Ukurannya 2KB yang biasanya thread adalah 1MB / 1000KB
    Meskipun thread, tetapi Berjalan secara concurrent

Cara kerja goroutine: 
    Dijalankan Go Scheduler (Fitur yang berada dalam Go Runtime), dimana jumlah threadnya sebanyak GOMAXPROCS (Sebanyak jumlah core CPU)
    Ada tiga istilah: 
        G: Goroutine
        M: Thread (Machine)
        P: Processor

Untuk menjalankannya: Gunakan go saat sebelum function yang berjalan secara asynchronous
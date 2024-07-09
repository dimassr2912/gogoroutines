Parallel: Membagi menjadi kecil, dan dijalankan bersamaan pada waktu yang sama 
    Contoh: 
        10 Soal dikerjakan 10 orang, jadi masing-masing mendapatkan pekerjaan
        Menjalankan aplikasi sekaligus (Chrome, Vscode, spotify)
    Process vs thread: 
        Process: 
            Sebuah eksekusi program (Menjalankan chrome, spotify (Aplikasi itu disebut process))
            Konsumsi memori besar 
            Terisolasi dengan process lain (Tidak akan berkomunikasi antar chrom dan spotify)
            lama dijalankan dan dihentikan
        Thread: 
            Segmen dari process (Membuka chrome lalu membuka banyak tab)
            Di dalam process ada banyak thread 
            Konsumsi memori kecil 
            Berhubungan jika dalam process yang lain (Bisa mengirim dari satu thread ke thread lain)
            Cepat dijalankan dan dihentikan

Concurrency: Menjalankan pekerjaan bergantian
Paralel vs Concurrency:
    Paralel: Butuh banyak thread
    Concurrency: Sedikit thread

Default golang: Concurrency

Cpu-bound: Algoritma yang hanya butuh Cpu untuk menjalankan, contoh Machine learning, lebih bagus menggunakan paralel 

I/O-bound: Kebalikan cpu-bound, bergantung pada input output (membaca database, dari file, dll): Lebih bagus menggunakan concurrency
    Contoh: Saat memanggil database 1 detik, daripada menunggu, bisa mengerjakan hal lain
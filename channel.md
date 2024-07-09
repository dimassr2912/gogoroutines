Channel: Komunikasi synchronous yang bisa dilakukan goroutine
    Jika ingin menggunakan function return value gunakan channel
    Di channel terdapat pengirim dan penerima (duanya goroutine) tetapi berbeda goroutine
    Karena sychronous: Saat melakukan pengiriman data ke channel, goroutine akan terbloc sampai ada yang menerima

Karakteristik channel: 
    Hanya bisa menampung satu data
    Hanya bisa menerima satu jenis data 
    Channel bisa diambil dari satu goroutine 
    Channel harus di close, agar tidak terjadi memory lea

Channel direpresentasikan dengan tipe data chan

Tidak perlu mengirim menggunakan pointer

Kadang channel hanya ingin mengirim saja / menerima saja
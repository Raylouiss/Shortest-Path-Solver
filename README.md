# Tucil3_13521054_13521143
Tugas Besar 3 IF2211 Strategi Algoritma Implementasi Algoritma UCS dan A* untuk Menentukan Lintasan Terpendek

## Daftar Isi
* [Deskripsi Singkat Program](#deskripsi-singkat-program)
* [Requirements](#requirements)
* [Cara Menjalankan Program](#cara-menjalankan-program)
* [Dibuat Oleh](#dibuat-oleh)

## Deskripsi Singkat Program
Program ini merupakan program yang menggunakan algoritma A* dan UCS untuk mendapatkan lintasan terdekat dari suatu titik ke titik lainnya pada suatu graf maupun peta. Untuk visualisasi, telah dibuat GUI sederhana dengan menggunakan bantuan Tkinter yang dihubungkan menggunakan metode interprocess communication (IPC) ke algoritma yang terdapat pada file berbahasa Go.

## Struktur Program
```bash
├─── test
│   ├─── Origin.txt
│   ├─── tesAlunAlun.txt
│   ├─── tesBuahBatu.txt
│   └─── tesITB.txt
├─── doc
│   └─── Tucil3_13521054_13521143.pdf
├─── src
│   ├─── Algorithm
│   │     ├─── Astar.go
│   │     ├─── UCS.go
│   ├─── Class
│   │     ├─── graph.go
│   ├─── gui.py
│   └─── main.go
├─── go.mod
├─── go.sum
├─── README.md
└─── run.bat                                    
```

## Requirements
* Python 3.11
* go1.20.3
## Cara Compile dan Menjalankan Program
1. Clone folder with `git clone https://github.com/Raylouiss/Tucil3_13521054_13521143.git`
2. Go to the folder `cd Tucil3_13521054_13521143.git`
3. Run the program `./run.bat`

## Cara Menggunakan Program

### Menjalankan Go
Untuk menjalankan go, diperlukan untuk menekan tombol Upload File terlebih dahulu. Lalu, akan muncul windows security alert, lalu tekan allow access. Maka go akan dijalankan dan terhubung dengan GUI.

### Memberi Input
Sistem akan menerima 4 buah input, yaitu file txt, starting point, goal point, dan jenis algoritma yang ingin digunakan. Untuk memasukkan input file txt, diperlukan untuk menekan tombol upload file dan pilih file txt yang sesuai. Setelah file txt dipilih, peta akan menampilkan graph dari file txt tersebut beserta dengan lokasi - lokasinya. Lalu, untuk starting point dan goal point diperlukan untuk diisi sesuai dengan salah satu dari nama lokasi yang terdapat pada file atau yang telah ditampilkan pada peta. Terakhir, untuk jenis algoritma yang ingin digunakan, diperlukan untuk menekan salah satu dari radio button yang ada sesuai dengan algoritma yang diinginkan.

### Mencari Lintasan Terdekat
Untuk mencari lintasan terdekat, diperlukan untuk menekan tombol search, lalu sistem akan menampilkan path pada peta dari lokasi start ke lokasi goal dan menampilkan lokasi mana saja yang dilewati, serta berapa total jarak lintasan tersebut.

### Mengganti Peta
Sistem akan menyediakan 3 tipe tampilan peta, yaitu openStreetMap, google normal, dan google satelite. Untuk mengganti tampilan peta, diperlukan untuk menekan tombol berwarna biru di bagian kiri bawah dan tampilan dapat dipilih sesuai dengan yang diinginkan.

### Error
Bila terjadi error atau pada tahap menjalankan go tidak muncul windows security alert, tutup semua terminal yang ada terlebih dahulu, lalu compile dan jalankan program dari awal kembali.

## Dibuat Oleh
* Nama : Wilson Tansil
* NIM : 13521054
* Prodi/Jurusan : STEI/Teknik Informatika
* Profile Github : Tansil011019
* Kelas : K01
##
* Nama: Raynard Tanadi
* NIM: 13521143
* Prodi/Jurusan: STEI/Teknik Informatika
* Profile Github : Raylouiss
* Kelas : K01

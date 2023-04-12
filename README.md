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
![windows_security_alert](https://user-images.githubusercontent.com/92111319/231442538-4628f2de-f1fb-4e65-abdd-883b2c6f1577.jpg)

### Format File txt
Format file txt diharuskan untuk seperti contoh berikut yang terdiri dari jumlah node di baris pertama, lalu nama latitude dan longitude node, serta adjacency Matrix.
![format_file_txt](https://user-images.githubusercontent.com/92111319/231442561-606d0934-d91f-4bec-bb0b-d3e71cdadfa5.jpg)

### Memberi Input
Sistem akan menerima 4 buah input, yaitu file txt, starting point, goal point, dan jenis algoritma yang ingin digunakan. Untuk memasukkan input file txt, diperlukan untuk menekan tombol upload file dan pilih file txt yang sesuai. Setelah file txt dipilih, peta akan menampilkan graph dari file txt tersebut beserta dengan lokasi - lokasinya. Lalu, untuk starting point dan goal point diperlukan untuk diisi sesuai dengan salah satu dari nama lokasi yang terdapat pada file atau yang telah ditampilkan pada peta. Terakhir, untuk jenis algoritma yang ingin digunakan, diperlukan untuk menekan salah satu dari radio button yang ada sesuai dengan algoritma yang diinginkan.
![input](https://user-images.githubusercontent.com/92111319/231442559-8910b5af-1b91-4338-8ae6-7e7012b493cd.jpg)

### Mencari Lintasan Terdekat
Untuk mencari lintasan terdekat, diperlukan untuk menekan tombol search, lalu sistem akan menampilkan path pada peta dari lokasi start ke lokasi goal dan menampilkan lokasi mana saja yang dilewati, serta berapa total jarak lintasan tersebut.
![output](https://user-images.githubusercontent.com/92111319/231442554-5283f950-3cda-4db0-9120-357d22e21b83.jpg)

### Fitur Peta
Sistem akan menyediakan 3 tipe tampilan peta, yaitu openStreetMap, google normal, dan google satelite. Untuk mengganti tampilan peta, diperlukan untuk menekan tombol berwarna biru di bagian kiri bawah dan tampilan dapat dipilih sesuai dengan yang diinginkan. Selain itu, peta juga dapat digeser dan dapat diatur untuk melakukan zoom in maupun zoom out.
![normal](https://user-images.githubusercontent.com/92111319/231442550-18f4fd43-d15e-4505-aff9-bb9a9efe6af9.jpg)
![satelite](https://user-images.githubusercontent.com/92111319/231442545-bceab82b-d78c-4574-932c-d230b81a33ec.jpg)

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

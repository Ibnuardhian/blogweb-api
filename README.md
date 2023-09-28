# BlogWeb API

API ini adalah sebuah sistem back-end yang memungkinkan manajemen artikel pada sebuah blog melalui operasi CRUD, menggunakan Go, Nodejs dan MySql.

## Daftar Isi

- [Deskripsi](#deskripsi)
- [Persyaratan](#persyaratan)
- [Setup dan Instalasi](#setup-dan-instalasi)
- [Endpoint API](#endpoint-api)
  - [Pendaftaran dan Login](#pendaftaran-dan-login)
  - [Manajemen Artikel](#manajemen-artikel)
  - [Upload Gambar](#upload-gambar)
- [Lisensi dan Kontribusi](#lisensi-dan-kontribusi)
- [Kontak dan Dukungan](#kontak-dan-dukungan)
- [Credits](#credits)

## Deskripsi

BlogWeb API menyediakan operasi sebagai berikut:
- Pendaftaran dan login pengguna
- Membuat, mengambil, memperbarui, dan menghapus artikel
- Mengunggah gambar

## Persyaratan
- Go 1.16 (atau terbaru)
- MySQL (atau database lainnya yang kompatibel)
- Git

## Setup dan Instalasi

### 1. Kloning Repositori
```sh
git clone https://github.com/Ibnuardhian/blogweb-api
cd blogweb-api
```

### 2. Konfigurasi Database
Ubah env menggunakan 
```
mv .env.example .env
```
Buat database dan konfigurasi koneksi pada `.env`.

### 3. Install Dependencies
```sh
go mod tidy
```

### 4. Install Nodemon
```sh
npm install nodemon -g
```
### 5. Menjalankan Server
```sh
nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go
```

## Endpoint API

### Pendaftaran dan Login
- **Register**: 
  ```sh
  POST /api/register
  ```

- **Login**: 
  ```sh
  POST /api/login
  ```

### Manajemen Artikel
- **Membuat Artikel Baru**: 
  ```sh
  POST /api/post
  ```

- **Mengambil Semua Artikel**: 
  ```sh
  GET /api/allpost
  ```

- **Mengambil Artikel Berdasarkan ID**: 
  ```sh
  GET /api/allpost/:id
  ```

- **Mengambil Artikel Unik**: 
  ```sh
  GET /api/uniquepost
  ```

- **Update Artikel Berdasarkan ID**: 
  ```sh
  PUT /api/updatepost/:id
  ```

- **Menghapus Artikel Berdasarkan ID**: 
  ```sh
  DELETE /api/deletepost/:id
  ```

### Upload Gambar
- **Mengunggah Gambar**: 
  ```sh
  POST /api/upload-image/
  ```

- **Mengakses Gambar yang Diunggah**: 
  ```sh
  GET /api/uploads/{filename}
  ```

## Lisensi dan Kontribusi

Projek ini berlisensi di bawah [Lisensi MIT](LICENSE).

Kontribusi sangat diterima. Silakan buat Pull Request atau buka Issue jika Anda menemui bug atau ingin menambahkan fitur baru.

## Kontak dan Dukungan

Untuk bantuan lebih lanjut, silakan hubungi [Nama Anda atau Tim Anda](mailto:email@example.com).

## Credits

Terima kasih kepada semua kontributor yang telah membantu pengembangan proyek ini.

---

Jangan lupa untuk menggantikan placeholder seperti `email@example.com` dengan informasi yang benar dan relevan dengan proyek Anda.

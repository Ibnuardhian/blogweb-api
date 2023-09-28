# BlogWeb-API
* Sebuah API CRUD Mysql yang dapat mengatur sebuah Website Blog
* Dibuat dengan Golang dan Mysql

## Isi
- [Instalasi](#instalasi)
- [Dokumentasi](#dokumentasi)
# Instalasi
Clone Repositori ini
```
git clone https://github.com/Ibnuardhian/blogweb-api
```
Instal dependensi yang diperlukan menggunakan
```
go mod tidy
```
Start Mysql di Xampp
Buat Database dengan nama gowebblog di Mysql
Ubah file env
```
mv .env.example .env
```
Lalu Jalankan kode dibawah dengan 
```
nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go
```
## Dokumentasi
* POST {Host}/api/register
* API untuk Verifikasi dan mendaftarkan user
* Response
```
{
  "message": "Account successfully created",
  "user": {
    "id": 2,
    "first_name": "New",
    "last_name": "Data",
    "email": "dummy@gmail.com",
    "phone": "1234567890"
  }
}
```
* POST {Host}/api/login
* API untuk user dapat login
* Response
```
{
  "message": "You have successfully login!",
  "user": {
    "id": 2,
    "first_name": "New",
    "last_name": "Data",
    "email": "dummy@gmail.com",
    "phone": "1234567890"
  }
}
```
* POST {Host}/api/post
* API untuk user membuat sebuah postingan
* Response
```
{
  "message": "Congratulations!, Your post is live"
}
```
* GET {Host}/api/allpost?page=[your page number]
* API untuk melihat semua post
* Response
```
[
  {
    "data": {
      "id": 1,
      "Title": "Here the test",
      "Desc": "Just Testing it",
      "Image": "https://images.unsplash.com/photo-1498462440456-0dba182e775b?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8c3BsYXNofGVufDB8fDB8fHww&auto=format&fit=crop&w=500&q=60",
      "UserID": "1",
      "user": {
        "id": 1,
        "first_name": "New",
        "last_name": "Data",
        "email": "dummy@gmail.com",
        "phone": "1234567890"
      }
    }
  },
  {
    "meta": {
      "last_page": 1,
      "page": 1,
      "total": 1
    }
  }
]

* GET {Host}/api/allpost/:id
* API untuk melihat post dengan id tertentu
* Response
{
  "data": {
    "id": 1,
    "Title": "Here the test",
    "Desc": "Just Testing it",
    "Image": "https://images.unsplash.com/photo-1498462440456-0dba182e775b?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8c3BsYXNofGVufDB8fDB8fHww&auto=format&fit=crop&w=500&q=60",
    "UserID": "1",
    "user": {
      "id": 1,
      "first_name": "New",
      "last_name": "Data",
      "email": "dummy@gmail.com",
      "phone": "1234567890"
    }
  }
}
```
* GET {Host}/api/updatepost/:id
* API untuk mengubah postingan
* Response
```
{
  "message": "success update post"
}

* GET {Host}/api/allpost/:id
* API untuk melihat postingan dengan id tertentu
* Response
{
  "data": {
    "id": 1,
    "Title": "Updated Post!",
    "Desc": "Updated",
    "Image": "https://plus.unsplash.com/premium_photo-1680871320511-8be2085ff281?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MXx8c3BsYXNofGVufDB8fDB8fHww&auto=format&fit=crop&w=500&q=60",
    "UserID": "1",
    "user": {
      "id": 1,
      "first_name": "New",
      "last_name": "Data",
      "email": "dummy@gmail.com",
      "phone": "1234567890"
    }
  }
}
```
* GET {Host}/api/uniquepost
* API untuk melihat postingan dari user tertentu
* Response
```
[
  {
    "id": 1,
    "Title": "Updated Post!",
    "Desc": "Updated",
    "Image": "https://plus.unsplash.com/premium_photo-1680871320511-8be2085ff281?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxzZWFyY2h8MXx8c3BsYXNofGVufDB8fDB8fHww&auto=format&fit=crop&w=500&q=60",
    "UserID": "1",
    "user": {
      "id": 1,
      "first_name": "New",
      "last_name": "Data",
      "email": "dummy@gmail.com",
      "phone": "1234567890"
    }
  },
  {
    "id": 2,
    "Title": "Here the test 2",
    "Desc": "Just Testing it 2",
    "Image": "https://images.unsplash.com/photo-1509233725247-49e657c54213?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1949&q=80",
    "UserID": "1",
    "user": {
      "id": 1,
      "first_name": "New",
      "last_name": "Data",
      "email": "dummy@gmail.com",
      "phone": "1234567890"
    }
  }
]
```
* GET {Host}/api/deletepost/:id
* API untuk menghapus postingan
* Response
```
{
  "status": "Post deleted "
}
```
* POST {Host}/api/upload-image/
* API untuk mengupload sebuah gambar
* Response
```
{
  "url": "http://{Host}:{Port}/api/uploads/[Generated string]-[filename].png"
}
```
## Have A Great Day!

# Article API

REST API untuk mengelola article, dibuat dengan Golang, Gin, dan MySQL.

## Cara menjalankan

1. Buat database MySQL bernama `post_article`
2. Sesuaikan koneksi di `database/database.go` (user dan password MySQL)
3. Jalankan migrasi: `go run migrate/main.go`
4. Jalankan server: `go run .`

Server berjalan di `http://localhost:8080`.

## Endpoint

| Method | URL | Keterangan |
|---|---|---|
| POST | /article/ | Membuat article baru |
| GET | /article/:limit/:offset | Daftar article dengan paging |
| GET | /article/:id | Detail article |
| PUT | /article/:id | Mengubah article |
| DELETE | /article/:id | Menghapus article |

## Postman

Collection ada di folder `postman/`. Import filenya ke Postman.
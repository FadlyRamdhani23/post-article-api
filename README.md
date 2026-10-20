# Article API

Backend untuk test Sharing Vision. Dibuat pakai Golang, Gin, dan MySQL.

## Cara menjalankan

1. Buat database dulu di MySQL:

```sql
CREATE DATABASE article;
```

2. Cek koneksi di `database/database.go`, sesuaikan user dan password MySQL.

3. Install dependency dan jalankan migrasi:

```bash
go mod tidy
go run migrate/main.go
```

4. Jalankan servernya:

```bash
go run .
```

Server jalan di `http://localhost:8080`.

## Endpoint

- `POST /article/` : buat article baru
- `GET /article/:limit/:offset` : list article (pakai paging)
- `GET /article/:id` : detail article
- `PUT /article/:id` : ubah article
- `DELETE /article/:id` : hapus article

Contoh body untuk POST dan PUT:

```json
{
  "title": "Judul article minimal 20 karakter",
  "content": "Isi article minimal 200 karakter",
  "category": "Teknologi",
  "status": "publish"
}
```

Status yang boleh: `publish`, `draft`, `thrash`.

## Postman

File collection ada di folder `Postman`, tinggal di-import.
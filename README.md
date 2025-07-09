# Backend API (Go + Gin)

Sebuah framework backend REST API berbasis Golang dan Gin, dengan struktur modular dan siap dikembangkan sesuai kebutuhan, dalam framework ini sudah termasuk manajemen user, role, dan autentikasi JWT.

---

## Fitur Utama

- **Autentikasi JWT** (Login, Register)
- **Manajemen User** (CRUD + Pagination ala Laravel)
- **Manajemen Role** (CRUD + Pagination)
- **Middleware Auth**
- **Struktur folder modular (MVC + Repository)**
- **Konfigurasi via .env**
- **Response JSON konsisten dan standar**
- **Pagination response mirip Laravel**

---

## Struktur Folder

```
/controllers   # Handler endpoint
/helpers       # Fungsi bantu (hash, pagination, dsb)
/middlewares   # Middleware (auth, dsb)
/models        # Model database (GORM)
/repositories  # Query ke database
/routes        # Routing API
/structs       # Struct untuk request/response
/config        # Konfigurasi aplikasi
main.go        # Bootstrap aplikasi
.env           # Konfigurasi environment
```

---

## Instalasi & Menjalankan

1. **Clone repo ini**
2. **Copy `.env.example` ke `.env`** dan sesuaikan konfigurasi (DB, JWT_SECRET, dsb)
3. **Install dependency**
   ```sh
   go mod tidy
   ```
4. **Jalankan migrasi database** (jika ada)
5. **Jalankan aplikasi**
   ```sh
   go run main.go
   ```
   atau
   ```sh
   go build && ./POS
   ```

---

## Contoh Endpoint

- `POST   /api/login`
- `POST   /api/register`
- `GET    /api/user?page=1&limit=10`
- `GET    /api/user/:id`
- `POST   /api/role`
- `GET    /api/role?page=1&limit=10`
- dst.

---

## Standar Response

```json
{
  "success": true,
  "message": "Berhasil mengambil data user",
  "data": {
    "current_page": 1,
    "per_page": 10,
    "total": 25,
    "last_page": 3,
    "from": 1,
    "to": 10,
    "data": [
      { "id": 1, "name": "John Doe", ... }
    ]
  }
}
```

---

## Kontribusi

- Pull request dan issue sangat terbuka!
- Ikuti standar penamaan dan struktur folder yang sudah ada.

---

## Lisensi

MIT License

---

## Database & Dependency

### Database

- **PostgreSQL**  
  Framework ini menggunakan PostgreSQL sebagai database utama.  
  Koneksi database diatur melalui file `.env`:
  ```
  DB_HOST=localhost
  DB_PORT=5432
  DB_USER=postgres
  DB_PASS=password
  DB_NAME=pos_db
  ```
- ORM yang digunakan: **GORM** (https://gorm.io/)  
  Anda dapat menyesuaikan koneksi di file konfigurasi sesuai kebutuhan.

### Dependency Utama

- [Gin](https://github.com/gin-gonic/gin) — HTTP web framework
- [GORM](https://github.com/go-gorm/gorm) — ORM untuk Golang
- [GoDotEnv](https://github.com/joho/godotenv) — Loader file `.env`
- [golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt) — Library JWT untuk autentikasi
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) — Untuk hash password
- [validator](https://github.com/go-playground/validator) — Validasi struct request

### Cara Install Dependency

Semua dependency akan otomatis terinstall saat menjalankan:
```sh
go mod tidy
```

---

**Catatan:**  
- Pastikan PostgreSQL sudah berjalan dan kredensial di `.env` sudah benar sebelum menjalankan aplikasi.
- Untuk development/testing, Anda bisa menggunakan database lain yang didukung GORM dengan sedikit penyesuaian

---

**Dikembangkan dengan ❤️ oleh Candra
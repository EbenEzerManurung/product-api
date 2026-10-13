<div align="center">

<img width="200" height="200" alt="Product API Logo" src="[https://cdn-icons-png.flaticon.com/512/3233/3233509.png"](https://img.shields.io/badge/Go-1.24-00ADD8?style=for-the-badge&logo=go&logoColor=white) />

<br/>

# Product Management REST API

**Backend REST API modern untuk manajemen produk dengan autentikasi JWT, rate limiting, dan dukungan Docker.**

<br/>

![Go](https://img.shields.io/badge/Go-1.24-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![MySQL](https://img.shields.io/badge/MySQL-8.0-4479A1?style=for-the-badge&logo=mysql&logoColor=white)
![Gin](https://img.shields.io/badge/Gin-1.10-008ECF?style=for-the-badge&logo=gin&logoColor=white)
![GORM](https://img.shields.io/badge/GORM-1.25-4CAF50?style=for-the-badge&logo=go&logoColor=white)

![JWT](https://img.shields.io/badge/Auth-JWT-0F766E?style=flat-square&logo=jsonwebtokens&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=flat-square&logo=docker&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-F59E0B?style=flat-square)
![Status](https://img.shields.io/badge/Status-Active-10B981?style=flat-square)

<br/>

**[Overview](#-overview)** &nbsp;•&nbsp; **[Features](#-features)** &nbsp;•&nbsp; **[Tech Stack](#-tech-stack)** &nbsp;•&nbsp; **[Architecture](#-architecture)** &nbsp;•&nbsp; **[Getting Started](#-getting-started)** &nbsp;•&nbsp; **[API Endpoints](#-api-endpoints)** &nbsp;•&nbsp; **[Postman Testing](#-postman-testing)** &nbsp;•&nbsp; **[Roadmap](#-roadmap)**

</div>

<br/>

---

## 📖 Overview

**Product Management REST API** adalah backend service yang dibangun dengan **Go** dan **Gin Framework** untuk mengelola data produk secara efisien. API ini menyediakan operasi CRUD lengkap, autentikasi berbasis JWT, rate limiting untuk mencegah penyalahgunaan, dan dukungan CORS untuk integrasi dengan aplikasi frontend dari berbagai domain.

Project ini dirancang mengikuti best practice industri: **clean layered architecture** (`config`, `controller`, `service`, `repository`, `model`), dependency injection, dan penggunaan **GORM** sebagai ORM untuk MySQL.

<br/>

<table>
  <tr>
    <td width="33%" valign="top">
      <h4>🔐 Authentication</h4>
      JWT-based authentication dengan access token &amp; refresh token. Password di-hash menggunakan bcrypt.
    </td>
    <td width="33%" valign="top">
      <h4>⚡ Rate Limiting</h4>
      Proteksi endpoint sensitif dengan rate limit per IP untuk mencegah abuse &amp; brute-force.
    </td>
    <td width="33%" valign="top">
      <h4>🐳 Docker Ready</h4>
      Multi-stage Dockerfile + docker-compose untuk deployment yang konsisten di semua environment.
    </td>
  </tr>
  <tr>
    <td width="33%" valign="top">
      <h4>🌐 CORS Support</h4>
      Cross-Origin Resource Sharing aktif — siap diintegrasikan dengan frontend dari domain manapun.
    </td>
    <td width="33%" valign="top">
      <h4>📦 Fresh Migrate &amp; Seeder</h4>
      Satu perintah untuk drop, migrate ulang, dan seed data dummy: <code>go run main.go fresh:migrate --seeder</code>.
    </td>
    <td width="33%" valign="top">
      <h4>✅ Input Validation</h4>
      Validasi server-side untuk semua input — title, price, category, dan images wajib diisi.
    </td>
  </tr>
</table>

<br/>

---

## ✨ Features

<details open>
<summary><b>🛍️ Product Management (CRUD)</b></summary>
<br/>

- **Create** produk baru dengan validasi lengkap (title, price > 0, category, min. 1 image)
- **Read** semua produk dengan dukungan:
  - 🔍 Search berdasarkan judul (`?search=keyword`)
  - 🏷️ Filter berdasarkan kategori (`?category=Clothes`)
  - 📄 Pagination (`?limit=10&page=2`)
- **Read** detail produk by ID dengan handling **404 Not Found** yang proper
- **Update** produk berdasarkan ID dengan tracking `updated_by` & `updated_at`
- **Delete** produk berdasarkan ID (soft delete dengan GORM)

</details>

<details open>
<summary><b>🔐 Authentication</b></summary>
<br/>

- **Register** user baru dengan konfirmasi password
- **Login** dengan username & password → menghasilkan:
  - `authentication_token` (access token, expires 15 menit)
  - `refresh_token` (expires 7 hari)
- **Password hashing** menggunakan bcrypt (cost factor default)
- **JWT claims** menyimpan `sub`, `username`, `type`, `exp`, `iat`

</details>

<details open>
<summary><b>🛡️ Security Features</b></summary>
<br/>

- **Rate Limiting** per IP address:
  - Auth endpoints (register/login): **3 request / 60 detik**
  - Protected endpoints (create/update/delete product): **1 request / 5 detik**
- **CORS** aktif untuk semua origin (development-friendly)
- **JWT validation** di setiap endpoint yang dilindungi
- **Parameterized queries** via GORM untuk mencegah SQL injection
- **Input validation** dengan Gin binding tags

</details>

<details>
<summary><b>🗄️ Database Layer</b></summary>
<br/>

- **GORM** sebagai ORM dengan auto-migration
- **Custom type** `JSONStringSlice` untuk menyimpan array URL gambar sebagai JSON di MySQL
- **Soft delete** dengan `gorm.DeletedAt`
- **Timestamps** otomatis (`created_at`, `updated_at`)
- **Audit fields**: `created_by`, `created_by_id`, `updated_by`, `updated_by_id`

</details>

<details>
<summary><b>🐳 Deployment &amp; Tooling</b></summary>
<br/>

- **Multi-stage Dockerfile** — image final hanya ~15 MB (Alpine-based)
- **docker-compose.yml** — menjalankan API + MySQL dengan satu perintah
- **Custom CLI command** di `main.go`:
  - `go run main.go serve` — start server
  - `go run main.go fresh:migrate` — drop + re-migrate
  - `go run main.go fresh:migrate --seeder` — drop + re-migrate + seed
- **Environment variables** via `.env` dengan godotenv

</details>

<br/>

---

## 🛠️ Tech Stack

| Layer | Technology |
|---|---|
| **Language** | Go 1.24 |
| **Web Framework** | Gin 1.10 |
| **ORM** | GORM 1.25 + MySQL Driver |
| **Database** | MySQL 8.0 |
| **Authentication** | JWT (`golang-jwt/jwt/v5`) + bcrypt |
| **Rate Limiting** | `golang.org/x/time/rate` |
| **CORS** | `gin-contrib/cors` |
| **Config** | `joho/godotenv` |
| **Documentation** | Swagger (`swaggo/gin-swagger`) |
| **Container** | Docker &amp; docker-compose |

<br/>

---

## 🏗️ Architecture

Request mengalir melalui layer yang terpisah dengan jelas:

```text
┌──────────────────────────────────────────┐
│          Client (Postman / Browser)      │
└──────────────────────────────────────────┘
                     │  REST API (JSON)
                     ▼
┌──────────────────────────────────────────┐
│            Middleware Layer              │
│   CORS · Logger · Auth · Rate Limiter    │
└──────────────────────────────────────────┘
                     │
                     ▼
┌──────────────────────────────────────────┐
│           Route Layer (route/)           │
│       SetupRouter · Route Grouping       │
└──────────────────────────────────────────┘
                     │
                     ▼
┌──────────────────────────────────────────┐
│       Controller Layer (controller/)     │
│   Bind JSON · Validate · Format Response │
└──────────────────────────────────────────┘
                     │
                     ▼
┌──────────────────────────────────────────┐
│         Service Layer (service/)         │
│      Business Logic · JWT · Hashing      │
└──────────────────────────────────────────┘
                     │
                     ▼
┌──────────────────────────────────────────┐
│      Repository Layer (repository/)      │
│         GORM · SQL · Transactions        │
└──────────────────────────────────────────┘
                     │
                     ▼
┌──────────────────────────────────────────┐
│              MySQL Database              │
│         users · products (tables)        │
└──────────────────────────────────────────┘

<details> <summary><b>📁 Project Structure</b></summary> <br/>
text
product-api/
├── config/
│   └── config.go                 # Load environment variables
│
├── controller/
│   ├── auth_controller.go        # Auth HTTP handlers
│   └── product_controller.go     # Product HTTP handlers
│
├── database/
│   └── database.go               # DB connection + FreshMigrate
│
├── middleware/
│   ├── auth.go                   # JWT validation middleware
│   ├── cors.go                   # CORS configuration
│   └── rate_limit.go             # Per-IP rate limiter
│
├── model/
│   └── model.go                  # GORM models (User, Product)
│
├── repository/
│   ├── product_repository.go     # Product data access
│   └── user_repository.go        # User data access
│
├── route/
│   └── routes.go                 # Router setup
│
├── seeder/
│   └── seeder.go                 # Seed dummy data
│
├── service/
│   ├── auth_service.go           # Auth business logic
│   └── product_service.go        # Product business logic
│
├── docs/                         # Swagger docs (generated)
├── .env                          # Environment variables
├── .dockerignore
├── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── main.go                       # CLI entry point
└── README.md
</details>
🚀 Getting Started
Prerequisites
Go 1.24+ — Download

MySQL 8.0+ — atau XAMPP / Laragon

Docker Desktop (opsional) — Download

Postman (opsional) — untuk testing

Installation (Lokal)
bash
# 1. Clone repository
git clone https://github.com/USERNAME/product-api.git
cd product-api

# 2. Buat database
mysql -u root -p -e "CREATE DATABASE product_api;"

# 3. Setup environment variables
cp .env.example .env
# Edit .env sesuai kredensial MySQL Anda

# 4. Download dependencies
go mod tidy

# 5. Migrate + seed data awal
go run main.go fresh:migrate --seeder

# 6. Jalankan server
go run main.go serve
Server berjalan di: http://localhost:8080

Installation (Docker)
bash
# 1. Clone repository
git clone https://github.com/USERNAME/product-api.git
cd product-api

# 2. Build image
docker compose build

# 3. Jalankan MySQL
docker compose up -d mysql

# 4. Tunggu MySQL healthy (~30 detik)
docker compose ps

# 5. Migrate + seeder
docker compose run --rm api ./server fresh:migrate --seeder

# 6. Jalankan API
docker compose up -d api
Server berjalan di: http://localhost:8080

Demo Account (dari Seeder)
Username	Password	Role
jhon_doe	supersecret	User
⚠️ Warning: Akun demo hanya untuk development. Ganti sebelum deploy ke production.


⚙️ Environment Variables
Buat file .env di root project:

env
# Database
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=product_api

# JWT
JWT_SECRET=your-super-secret-jwt-key-min-32-chars
JWT_ACCESS_EXPIRY=15m
JWT_REFRESH_EXPIRY=168h

# Server
SERVER_PORT=8080
CACHE_TTL=60s

📡 API Endpoints
🔓 Public Endpoints
Method	Endpoint	Description
POST	/api/auth/register	Registrasi user baru
POST	/api/auth/login	Login & dapatkan JWT token
GET	/api/products	List semua produk (+ search, filter, pagination)
GET	/api/products/:id	Detail produk by ID
🔒 Protected Endpoints (butuh Bearer Token)
Method	Endpoint	Description
POST	/api/products	Buat produk baru
PUT	/api/products/:id	Update produk
DELETE	/api/products/:id	Hapus produk
📋 Query Parameters untuk GET /api/products
Parameter	Contoh	Description
search	?search=T-Shirt	Cari produk berdasarkan judul (LIKE)
category	?category=Clothes	Filter berdasarkan kategori (exact match)
limit	?limit=10	Jumlah item per halaman (default: 10)
page	?page=2	Nomor halaman (default: 1)
🧪 Contoh Request
1. Register User
bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "jhon_doe",
    "password": "supersecret",
    "password_confirmation": "supersecret"
  }'
Response (201 Created):

json
{
  "message": "User berhasil didaftarkan",
  "user": {
    "id": 1,
    "username": "jhon_doe"
  }
}
2. Login
bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "jhon_doe", "password": "supersecret"}'
Response (200 OK):

json
{
  "authentication_token": "eyJhbGciOiJIUzI1NiIs...",
  "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
}
3. Create Product (butuh token)
bash
curl -X POST http://localhost:8080/api/products \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <TOKEN>" \
  -d '{
    "title": "Awesome T-Shirt",
    "price": 99.99,
    "description": "High-quality cotton t-shirt",
    "category": "Clothes",
    "images": ["https://placeimg.com/640/480/any"]
  }'
Response (201 Created):

json
{
  "id": 1,
  "title": "Awesome T-Shirt",
  "price": 99.99,
  "description": "High-quality cotton t-shirt",
  "category": "Clothes",
  "images": ["https://placeimg.com/640/480/any"],
  "created_at": "2026-10-13T12:26:04+07:00",
  "created_by": "jhon_doe",
  "created_by_id": "1",
  "updated_at": "2026-10-13T12:26:04+07:00",
  "updated_by": "jhon_doe",
  "updated_by_id": "1"
}
4. Get Products dengan Filter
bash
# Search by title
curl "http://localhost:8080/api/products?search=T-Shirt"

# Filter by category
curl "http://localhost:8080/api/products?category=Electronics"

# Pagination
curl "http://localhost:8080/api/products?limit=2&page=1"

# Kombinasi
curl "http://localhost:8080/api/products?search=Wallet&category=Accessories&limit=5"

🧪 Postman Testing
Koleksi Postman untuk testing semua endpoint bisa diimpor dari file postman_collection.json (tersedia di root project).

Alur Testing yang Disarankan
text
1. POST /api/auth/login           → dapatkan token
2. Copy authentication_token
3. POST /api/products             → buat produk (dengan Bearer token)
4. GET  /api/products             → lihat semua produk
5. GET  /api/products/:id         → detail produk
6. PUT  /api/products/:id         → update produk (dengan Bearer token)
7. DELETE /api/products/:id       → hapus produk (dengan Bearer token)
Setup Environment di Postman
Variable	Value
base_url	http://localhost:8080
token	<paste setelah login>
Testing Rate Limit
Auth endpoints: Kirim 4 request login dalam <60 detik → request ke-4 akan dapat 429 Too Many Requests

Product write endpoints: Kirim 2 request POST dalam <5 detik → request ke-2 akan dapat 429


📸 Screenshots
<details open> <summary><b>🖥️ API Response Examples</b></summary> <br/>
Login Success — Response with JWT Tokens

<img width="100%" alt="Login Response" src="docs/images/login-response.png" />
GET Products — List with Pagination

<img width="100%" alt="Product List" src="docs/images/product-list.png" />
POST Product — Create with Authorization

<img width="100%" alt="Create Product" src="docs/images/create-product.png" />
Rate Limit — 429 Too Many Requests

<img width="100%" alt="Rate Limit" src="docs/images/rate-limit.png" /></details><details> <summary><b>🐳 Docker Deployment</b></summary> <br/>
docker compose up -d — Containers Running

<img width="100%" alt="Docker Running" src="docs/images/docker-running.png" />
Fresh Migrate & Seeder Output

<img width="100%" alt="Migrate Seeder" src="docs/images/migrate-seeder.png" /></details>
🗺️ Roadmap
Shipped ✅	Planned ⬜
✅ Product CRUD (create, read, update, delete)	⬜ Caching dengan Redis
✅ JWT authentication (access + refresh)	⬜ File upload untuk product images
✅ bcrypt password hashing	⬜ Product categories sebagai tabel terpisah
✅ Rate limiting per IP	⬜ Real-time notifications (WebSocket)
✅ CORS support	⬜ Unit testing (testify + mockery)
✅ Search, filter, pagination	⬜ CI/CD dengan GitHub Actions
✅ Soft delete dengan GORM	⬜ API versioning (v1, v2)
✅ Docker & docker-compose	⬜ Swagger documentation UI
✅ Fresh migrate & seeder command	⬜ GraphQL endpoint (alternative)
✅ Input validation	⬜ Elasticsearch untuk search
✅ Multi-stage Docker build	⬜ Kubernetes manifests
✅ Environment-based config	⬜ Email verification on register

🎨 Design Decisions
Why Go? — Go menawarkan performa tinggi, concurrency model yang efisien (goroutines), static typing, dan kompilasi cepat. Cocok untuk backend REST API yang harus melayani banyak concurrent request.

Why Gin? — Gin adalah web framework Go paling populer dengan performa tinggi, middleware ekosistem yang kaya, dan sintaks routing yang bersih. Binding & validation bawaan menghemat waktu development.

Why GORM? — GORM mempercepat development dengan auto-migration, relasi model, hooks, dan soft delete. Menghilangkan boilerplate SQL untuk operasi CRUD standar.

Why JWT? — Stateless authentication — tidak perlu session storage di server, cocok untuk horizontal scaling. Access token + refresh token pattern memberikan security & UX yang seimbang.

Why Rate Limiting? — Melindungi endpoint sensitif (login, register) dari brute-force attack, dan mencegah abuse pada endpoint write (create/update/delete).

Why Layered Architecture? — Pemisahan concern yang jelas antara controller (HTTP), service (business logic), repository (data access), dan model (entity). Memudahkan testing, maintenance, dan scaling tim.


🤝 Contributing
Contributions, issues, dan feature requests sangat diterima!

Fork repository ini

Buat feature branch: git checkout -b feature/amazing-feature

Commit perubahan: git commit -m "Add amazing feature"

Push ke branch: git push origin feature/amazing-feature

Buat Pull Request


📄 License
Project ini dilisensikan di bawah MIT License — lihat file LICENSE untuk detail.


<div align="center"><img width="150" height="150" alt="Logo" src="https://cdn-icons-png.flaticon.com/512/3233/3233509.png" />
Dibuat dengan ❤️ oleh [Nama Anda]
Backend Developer • Go Enthusiast

https://img.shields.io/badge/GitHub-USERNAME-181717?style=for-the-badge&logo=github&logoColor=white
https://img.shields.io/badge/LinkedIn-Connect-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white
https://img.shields.io/badge/Email-Contact-EA4335?style=for-the-badge&logo=gmail&logoColor=white


Jika project ini membantu Anda, mohon berikan ⭐

<sub>Built with Go 1.24, Gin, GORM, MySQL 8, and Docker</sub>

</div> ```

# 🦴 Go Fiber + GORM Skeleton

A production-ready starter project for building RESTful APIs using **Go Fiber**, **GORM**, **PostgreSQL**, and **Makefile**.  
Includes clean architecture layout, database migration, Swagger docs, and more.

---

## 📦 Tech Stack

- [Fiber](https://github.com/gofiber/fiber) — Fast HTTP Web Framework for Go
- [GORM](https://gorm.io/) — ORM for Golang
- [PostgreSQL](https://www.postgresql.org/)
- [golang-migrate](https://github.com/golang-migrate/migrate) — SQL migrations
- [Swag](https://github.com/swaggo/swag) — Swagger API docs
- [mockery](https://github.com/vektra/mockery) — Interface mocking for tests
- [Air](https://github.com/cosmtrek/air) — Live-reloading for development

---

## ⚙️ Setup

### 1. Clone project

```bash
git clone https://github.com/yourname/go-fiber-skeleton.git
cd go-fiber-skeleton
```


2. Setup .env

PGSQL_USERNAME=postgres
PGSQL_PASSWORD=secret
PGSQL_HOST=localhost
PGSQL_PORT=5432
PGSQL_DATABASE_NAME=mydb
PGSQL_URI=postgres://postgres:secret@localhost:5432/mydb?sslmode=disable

3. Install dependencies

go mod tidy



⸻

🚀 Usage

Run dev with live reload
```bash
make dev
```

Run manually
```bash
make run
```

Build binary
```bash
make build
```

⸻

📂 Migration Commands

make migrate create=create_users_table      # Create new migration
make migrate_up                             # Apply all migrations
make migrate_down                           # Rollback latest migration
make auto_migrate                           # Use GORM to auto-migrate all models

⸻

📄 API Documentation

Generate Swagger docs:

make apidoc

View Swagger UI at:
http://localhost:<PORT>/swagger/index.html

⸻

🧪 Test

make test          # Run unit tests
make coverage      # Show coverage report
make mock d=UserService  # Generate mock for interface

⸻

✨ Credit

Built with ❤️ by [your name / team] — feel free to fork and contribute!

---

.PHONY: build run test test-verbose test-coverage setup-db clean help

help:
	@echo "Sistem Manajemen Pemesanan Ruang - Makefile Commands"
	@echo ""
	@echo "Available commands:"
	@echo "  make build          - Build aplikasi"
	@echo "  make run            - Jalankan aplikasi"
	@echo "  make test           - Jalankan unit tests"
	@echo "  make test-verbose   - Jalankan tests dengan verbose output"
	@echo "  make test-coverage  - Jalankan tests dengan coverage report"
	@echo "  make setup-db       - Setup database MySQL"
	@echo "  make clean          - Hapus binary yang sudah di-build"
	@echo "  make help           - Tampilkan help ini"

build:
	@echo "Building aplikasi..."
	@go build -o pesan-ruang.exe main.go
	@echo "Build completed: pesan-ruang.exe"

run:
	@echo "Starting server..."
	@go run main.go

test:
	@echo "Running tests..."
	@go test ./services -v

test-verbose:
	@echo "Running tests with verbose output..."
	@go test ./services -v -run .

test-coverage:
	@echo "Running tests with coverage..."
	@go test ./services -v -cover

setup-db:
	@echo "Setting up database..."
	@mysql -u root < database/schema.sql

clean:
	@echo "Cleaning up..."
	@if exist pesan-ruang.exe del pesan-ruang.exe
	@if exist laporan_*.json del laporan_*.json
	@if exist laporan_*.csv del laporan_*.csv
	@echo "Cleanup completed"

deps:
	@echo "Downloading dependencies..."
	@go mod download
	@go mod tidy
	@echo "Dependencies downloaded"

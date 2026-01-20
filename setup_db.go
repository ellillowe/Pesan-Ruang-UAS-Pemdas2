package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"pesan-ruang/config"

	_ "github.com/go-sql-driver/mysql"
)

func SetupDatabase() {
	// 1. Load konfigurasi
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Gagal load config: %v", err)
	}

	// 2. Konek ke MySQL tanpa nama database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Gagal konek ke MySQL: %v", err)
	}
	defer db.Close()

	// 3. Buat Database
	fmt.Printf("Resetting database '%s'...\n", cfg.Database.Name)
	// Drop database lama jika ada agar bersih (seperti script .bat)
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", cfg.Database.Name))
	if err != nil {
		log.Fatalf("Gagal drop database: %v", err)
	}

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s", cfg.Database.Name))
	if err != nil {
		log.Fatalf("Gagal membuat database: %v", err)
	}

	// 4. Baca file schema
	schemaContent, err := os.ReadFile("database/schema.sql")
	if err != nil {
		log.Fatalf("Gagal membaca file schema.sql: %v", err)
	}

	// 5. Konek ke database yang baru dibuat untuk import tabel
	dsnWithDB := fmt.Sprintf("%s%s?multiStatements=true", dsn, cfg.Database.Name)
	dbSchema, err := sql.Open("mysql", dsnWithDB)
	if err != nil {
		log.Fatalf("Gagal konek ke database baru: %v", err)
	}
	defer dbSchema.Close()

	fmt.Println("Mengimport tabel...")
	if _, err := dbSchema.Exec(string(schemaContent)); err != nil {
		log.Fatalf("Gagal import schema: %v", err)
	}

	fmt.Println("✅ Setup database berhasil! Silakan jalankan 'go run main.go'")
}

# Quick Start Guide - Panduan Cepat

## ⚡ Mulai dalam 5 Menit

### Prerequisites
- Go 1.21+
- MySQL Server (running)
- Terminal/Command Prompt

### Step 1: Setup Database (2 menit)

**Windows:**
```cmd
# Edit setup-db.bat dengan credentials MySQL Anda
setup-db.bat
```

**Linux/Mac:**
```bash
# Edit setup-db.sh dengan credentials MySQL Anda
chmod +x setup-db.sh
./setup-db.sh
```

**Manual** (Jika script tidak berfungsi):
```bash
# 1. Buka MySQL client
mysql -u root -p

# 2. Jalankan SQL dari database/schema.sql
source database/schema.sql
exit
```

### Step 2: Update Config (30 detik)

Edit `config.json`:
```json
{
  "database": {
    "host": "localhost",
    "port": "3306",
    "user": "root",              ← Ubah sesuai MySQL user Anda
    "password": "",              ← Ubah sesuai MySQL password
    "name": "pesan_ruang_db"     ← Ubah jika nama database berbeda
  },
  "server": {
    "port": "8080"
  }
}
```

### Step 3: Run Application (30 detik)

```bash
# Download dependencies (first time only)
go mod download

# Run the application
go run main.go
```

Expected output:
```
Starting server on port 8080
```

✅ Aplikasi running di: **http://localhost:8080**

### 🎨 Web Interface (Dashboard)

Buka browser dan akses: **http://localhost:8080**

Fitur yang tersedia:
- 📊 **Dashboard**: Statistics total ruang, pemesanan, dan status
- 🚪 **Ruang**: Kelola daftar ruang (tambah, edit, hapus)
- 📅 **Pemesanan**: Kelola pemesanan ruang (create, approve, reject)
- 📈 **Laporan**: Generate laporan bulanan JSON & CSV

UI menggunakan Bootstrap 5 dengan design responsive untuk desktop dan mobile.

## 🧪 Quick Testing

### Test 1: Check Health
```bash
curl http://localhost:8080/
```

### Test 2: Get All Rooms
```bash
curl http://localhost:8080/rooms
```

### Test 3: Create Room
```bash
curl -X POST http://localhost:8080/rooms \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Ruang Baru",
    "type": "kelas",
    "capacity": 50
  }'
```

### Test 4: Create Booking
```bash
curl -X POST http://localhost:8080/bookings \
  -H "Content-Type: application/json" \
  -d '{
    "room_id": 1,
    "user_id": 2,
    "date": "2025-02-20",
    "start_time": "08:00",
    "end_time": "10:00",
    "purpose": "Test Booking",
    "status": "pending"
  }'
```

### Test 5: Approve Booking
```bash
curl -X POST "http://localhost:8080/bookings/1?action=approve"
```

### Test 6: Generate Report
```bash
curl "http://localhost:8080/report?year=2025&month=2"
```

Check generated files:
```bash
ls laporan_*
cat laporan_202502.json
cat laporan_202502.csv
```

## ✅ Run Unit Tests

```bash
# All tests
go test ./services -v

# Specific test
go test ./services -v -run TestHasTimeConflict

# With coverage
go test ./services -v -cover
```

Expected output:
```
=== RUN   TestHasTimeConflict
=== RUN   TestHasTimeConflict/No_conflict_-_different_time_slots
=== RUN   TestHasTimeConflict/Conflict_-_overlapping_times
...
ok      pesan-ruang/services    0.005s  coverage: 45.0%
```

## 📁 Project Files Overview

```
pesan-ruang/
├── main.go                    ← Jalankan ini
├── config.json                ← Edit config DB
├── README.md                  ← Full documentation
├── ARCHITECTURE.md            ← Technical details
├── RUBRIC_COMPLIANCE.md       ← Rubric mapping
├── API_TESTING.md             ← API examples
│
├── config/
│   └── config.go              ← DB connection
├── models/
│   └── models.go              ← Data structures
├── repository/
│   ├── room_repository.go     ← Room CRUD
│   ├── booking_repository.go  ← Booking CRUD
│   └── user_repository.go     ← User queries
├── services/
│   ├── booking_service.go     ← Business logic
│   ├── room_service.go        ← Business logic
│   ├── booking_service_test.go ← Unit tests
│   └── room_service_test.go   ← Unit tests
├── handlers/
│   ├── room_handler.go        ← HTTP handlers
│   └── booking_handler.go     ← HTTP handlers
└── database/
    └── schema.sql             ← DB schema
```

## 🔍 Sample Data

Database automatically populated dengan:

**Users** (user_id):
- 1: Admin User
- 2: Dr. Budi (Dosen)
- 3: Dr. Ani (Dosen)

**Rooms** (room_id):
- 1: Ruang A101 (Kelas, 40 orang)
- 2: Ruang A102 (Kelas, 35 orang)
- 3: Lab Komputer 1 (Lab, 30 orang)
- 4: Lab Fisika 1 (Lab, 25 orang)

## 💡 Tips

1. **Test Conflict Detection**: Create booking 08:00-09:00, try 08:30-09:30 (should fail)

2. **Check Database**: 
   ```bash
   mysql -u root pesan_ruang_db
   SELECT * FROM bookings;
   SELECT * FROM rooms;
   ```

3. **View Logs**: Check terminal output untuk error messages

4. **Reset Database**: Re-run setup-db script untuk start fresh

5. **Change Port**: Edit config.json server.port untuk port berbeda

## ❌ Troubleshooting

| Problem | Solution |
|---------|----------|
| "Failed to initialize database" | Check MySQL credentials di config.json |
| "database pesan_ruang_db does not exist" | Run setup-db script |
| "Port 8080 already in use" | Change port di config.json |
| "Module not found" | Run `go mod download` |
| "Connection refused" | Check MySQL server is running |

## 📚 Full Documentation

- **Architecture**: See `ARCHITECTURE.md`
- **All Endpoints**: See `API_TESTING.md`
- **Setup Details**: See `README.md`
- **Rubric Mapping**: See `RUBRIC_COMPLIANCE.md`

## 🎯 Next Steps

1. ✅ Database ready
2. ✅ Config updated
3. ✅ App running
4. ✅ API tested
5. ✅ Tests passing
6. ✅ Ready for submission!

---

**Questions?** Check the documentation files or review the code comments.

Happy coding! 🚀

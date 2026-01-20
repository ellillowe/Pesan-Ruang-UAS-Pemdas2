# Sistem Manajemen Pemesanan Ruang - Room Booking Management System

Aplikasi web berbasis Go untuk mengelola pemesanan ruang kelas dan lab di kampus.

**Status**: ✅ Production Ready  
**Latest Version**: 1.2.1 (January 19, 2026 - Phase 1: File Persistence & Analytics)  
**Docs**: [CHANGELOG](CHANGELOG.md) | [QUICKSTART](QUICKSTART.md) | [API](API_TESTING.md) | [Phase 1](PHASE_1_IMPLEMENTATION.md)

## 🎯 Fitur Utama

- ✅ **Authentication** - Login/Register dengan bcrypt hashing
- ✅ **Room Management** - CRUD ruang (Create, Read, Update, Delete)
- ✅ **Booking System** - Membuat, approve, reject pemesanan
- ✅ **Conflict Detection** - Prevent double-booking otomatis
- ✅ **Dashboard** - Admin dashboard + User dashboard
- ✅ **Reports** - Monthly reports dengan file persistence (JSON & CSV)
- ✅ **Role-Based** - Admin vs Siswa access control
- ✨ **NEW Phase 1**:
  - ✅ Server-side file persistence untuk laporan
  - ✅ Download laporan endpoint dengan security checks
  - ✅ Pending approvals notification endpoint
  - ✅ Analytics stats untuk dashboard

## 🚀 Quick Start

```bash
# Setup database (first time)
go run setup_db.go

# Start server
go run main.go

# Open browser
# http://localhost:8080
```

Demo accounts:
- Admin: `admin@campus.edu` / `admin123`
- Siswa: `siswa@campus.edu` / `siswa123`

## 📁 Project Structure

```
pesan-ruang/
├── config/               # Database config
├── models/               # Data structures
├── repository/           # Data access layer
├── services/             # Business logic
├── handlers/             # HTTP handlers
├── database/schema.sql   # Database schema
├── static/
│   ├── login.html        # Login page
│   ├── admin-dashboard.html
│   ├── user-dashboard.html
│   └── js/ & css/
├── main.go
└── config.json
```

## 🔌 API Endpoints (12 Total)

| Method | Endpoint | Purpose |
|--------|----------|---------|
| POST | `/auth/login` | Login user |
| POST | `/auth/register` | Register user |
| GET | `/rooms` | List all rooms |
| POST | `/rooms` | Create room |
| PUT | `/rooms/{id}` | Update room |
| DELETE | `/rooms/{id}` | Delete room |
| GET | `/bookings` | List bookings |
| POST | `/bookings` | Create booking |
| POST | `/bookings/{id}?action=approve` | Approve booking |
| POST | `/bookings/{id}?action=reject` | Reject booking |
| GET | `/report` | Generate report |
| GET | `/api/health` | Health check |

## 🛠️ Tech Stack

- **Backend**: Go 1.21+ dengan Clean Architecture
- **Frontend**: Bootstrap 5 + Vanilla JavaScript
- **Database**: MySQL dengan 3 tables
- **Testing**: 18 unit tests (all passing)

## API Endpoints

### 🆕 Phase 1 Features - File Persistence & Analytics

#### GET /report?year=YYYY&month=MM
Generate dan simpan monthly report sebagai JSON & CSV

```bash
curl "http://localhost:8080/report?year=2026&month=1"
```

Response:
```json
{
  "message": "Report generated for 2026-01",
  "summary": [
    {
      "room_name": "Ruang A103",
      "total_bookings": 2,
      "total_waktu_pemesanan": "5 jam 50 menit"
    }
  ],
  "files": {
    "json": "laporan_202601.json",
    "csv": "laporan_202601.csv"
  }
}
```

#### GET /pending-approvals
Ambil data pending bookings untuk notification

```bash
curl http://localhost:8080/pending-approvals
```

#### GET /analytics/stats
Ambil data analytics untuk dashboard charts

```bash
curl http://localhost:8080/analytics/stats
```

#### GET /reports/list
List semua saved report files

```bash
curl http://localhost:8080/reports/list
```

#### GET /download/laporan/{filename}
Download laporan file (JSON atau CSV)

```bash
# Download JSON
curl http://localhost:8080/download/laporan/laporan_202601.json > report.json

# Download CSV
curl http://localhost:8080/download/laporan/laporan_202601.csv > report.csv
```

---

### Rooms Management

#### GET /rooms
Mengambil daftar semua ruang
```bash
curl http://localhost:8080/rooms
```

#### POST /rooms
Menambah ruang baru
```bash
curl -X POST http://localhost:8080/rooms \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Ruang A101",
    "type": "kelas",
    "capacity": 40
  }'
```## 📚 Documentation

See detailed docs in:
- **[QUICKSTART.md](QUICKSTART.md)** - Setup in 5 minutes
- **[ARCHITECTURE.md](ARCHITECTURE.md)** - System design details
- **[API_TESTING.md](API_TESTING.md)** - API examples with cURL
- **[UI_DOCUMENTATION.md](UI_DOCUMENTATION.md)** - Frontend guide
- **[CHANGELOG.md](CHANGELOG.md)** - Version history & changes

## 📋 Database Schema

3 tables with relationships:
- **users** - Store user accounts (admin/siswa)
- **rooms** - Store classroom/lab data
- **bookings** - Store booking records with status

See `database/schema.sql` for complete schema.

## ✨ Key Features in Detail

### 1. Authentication
- Login/Register system with bcrypt password hashing
- Role-based access (admin vs siswa)
- Session management via localStorage

### 2. Room Management
- Create new rooms (admin only)
- **Edit room details** (name, type, capacity, status)
- Delete rooms
- View active/inactive status

### 3. Booking System
- Create bookings with date/time/purpose
- Automatic conflict detection
- Booking workflow: pending → approved/rejected
- Cancel pending bookings (siswa)
- Approve/reject bookings (admin)
    end_time TIME NOT NULL,
    purpose VARCHAR(255) NOT NULL,
    status ENUM('pending', 'approved', 'rejected') DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (room_id) REFERENCES rooms(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);
```

### 4. Conflict Detection
- Automatically prevent double-booking
- Check room, date, and time overlaps
- Validates: `new_start < existing_end AND new_end > existing_start`

### 5. Reports
- Generate monthly reports (JSON & CSV)
- Summary of room usage
- Total hours booked per room

## 📊 Testing

Run all unit tests:
```bash
go test ./services -v
```

**Status**: All 18 tests passing ✅
- Conflict detection tests
- Hours calculation tests
- Room validation tests
- Edge case tests

## 🏗️ Architecture

```
Handlers (HTTP Layer)
    ↓
Services (Business Logic)
    ↓
Repository (Data Access)
    ↓
Database (MySQL)
```

- **Clean Architecture** - Separation of concerns
- **Repository Pattern** - Abstract database operations
- **Service Layer** - Centralized business logic
- **CORS Enabled** - Cross-origin requests supported

## 👥 User Roles

### Admin
- Create/Edit/Delete rooms
- Approve or reject bookings
- View all bookings
- Generate reports

### Siswa (User)
- Browse available rooms
- Create bookings
- View own bookings
- Cancel own pending bookings

## 📝 License & Credits

Academic project for Web Backend Programming course with Go.

---

**Questions?** See [CHANGELOG.md](CHANGELOG.md) for latest updates.

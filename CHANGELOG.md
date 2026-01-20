# CHANGELOG - Dokumentasi Perubahan & Pengembangan

## Version 1.2.0 - January 19, 2026 (Date/Time Format Fixes & Report Enhancement)

### 🐛 Critical Bug Fixes

#### 1. **Database Date Format Parsing Issue**
- **Problem**: Laporan menampilkan "Tidak ada data untuk bulan ini" meskipun ada pemesanan di dashboard
- **Root Cause**: Database DATE column di-scan sebagai timestamp ISO 8601 (`2026-01-20T00:00:00Z`), tetapi code parse dengan format `2006-01-02`
- **Solution**: Scan DATE sebagai `sql.NullTime`, format dengan `.Format("2006-01-02")`
- **Files Fixed**: `repository/booking_repository.go` (4 methods)
- **Result**: ✅ Report sekarang menampilkan data dengan benar

#### 2. **Time Format Parsing Issue (HH:MM:SS)**
- **Problem**: Laporan error "parsing time 07:00:00: extra text :00"
- **Root Cause**: Database TIME column berkformat `HH:MM:SS`, tetapi parsing hanya `15:04`
- **Solution**: Update `CalculateTotalHours` untuk support both formats
- **Files Fixed**: `services/booking_service.go`
- **Result**: ✅ Hours calculation support HH:MM dan HH:MM:SS

### ✨ UI Enhancements

#### 3. **Report Time Display Format**
- **Improvement**: Ubah format tampilan jam dari desimal ke format yang lebih readable
- **Before**: `5.833333333333334 jam`
- **After**: `5 jam 50 menit`
- **Implementation**: Buat fungsi `FormatHoursToWaktu()` di service layer
- **Files Modified**: 
  - `services/booking_service.go` - Tambah `FormatHoursToWaktu()` function
  - `services/booking_service.go` - Update `ReportSummary` struct (field baru: `TotalWaktuPemesanan`)
  - `static/js/reports-admin.js` - Update display logic
- **Result**: ✅ Report sekarang menampilkan waktu dalam format "X jam Y menit"

### ✅ Verification

```
GET /report?year=2026&month=1 Response:
{
    "summary": [
        {"room_name": "Ruang A103", "total_bookings": 2, "total_waktu_pemesanan": "5 jam 50 menit"},
        {"room_name": "Ruang A101", "total_bookings": 1, "total_waktu_pemesanan": "2 menit"}
    ]
}
```

---

## Version 1.1.0 - January 19, 2026

### ✨ Fitur Baru

#### 1. **Frontend Rebuild & Restructuring**
- ✅ Pisahkan halaman login dari dashboard utama (`login.html`, `register.html`)
- ✅ Buat admin dashboard terpisah (`admin-dashboard.html`)
- ✅ Buat user/siswa dashboard terpisah (`user-dashboard.html`)
- ✅ Setiap halaman adalah file standalone, bukan combined
- ✅ Routing utama diubah dari `/` → `index.html` menjadi `/` → `login.html`

#### 2. **JavaScript API Integration**
- ✅ Buat `admin-dashboard.js` dengan fungsi:
  - `checkAuth()` - Validasi role admin
  - `loadDashboard()` - Load statistik
  - `loadRooms()` - List ruang dengan button edit
  - `editRoom()` - Modal untuk edit ruang
  - `saveEditRoom()` - Update ruang via PUT endpoint
  - `saveRoom()` - Tambah ruang baru
  - `deleteRoom()` - Hapus ruang
  - `loadBookings()` - List pemesanan
  - `approveBooking()` - Setujui pemesanan
  - `rejectBooking()` - Tolak pemesanan

- ✅ Buat `user-dashboard.js` dengan fungsi:
  - `checkAuth()` - Validasi role siswa (block admin)
  - `loadDashboard()` - Load statistik personal
  - `loadRoomsGrid()` - Tampilkan ruang sebagai card grid
  - `openBookingModal()` - Modal form untuk booking
  - `saveBooking()` - Submit pemesanan baru
  - `loadMyBookings()` - List pemesanan saya
  - `cancelBooking()` - Batalkan pemesanan

#### 3. **Database Connection Fix**
- ✅ Fix DSN connection string: tambah `?parseTime=true`
- ✅ Ini fix error: "unsupported Scan, storing driver.Value type []uint8 into type *time.Time"
- ✅ Go MySQL driver sekarang properly parse TIMESTAMP columns
- **File Modified**: `config/config.go`

#### 4. **Feature: Edit Ruang (Edit Room)**
- ✅ Tambah button "Edit" di setiap row tabel ruang
- ✅ Buat modal form untuk edit ruang
- ✅ Form bisa edit: nama, tipe, kapasitas, status (aktif/tidak aktif)
- ✅ Save changes via PUT `/rooms/{id}` endpoint
- ✅ Tabel auto-refresh setelah edit
- **Files Modified**: 
  - `static/admin-dashboard.html` - Tambah Edit button & modal
  - `static/js/admin-dashboard.js` - Tambah `editRoom()` & `saveEditRoom()`

#### 5. **Authentication Flow**
- ✅ Demo accounts dengan bcrypt-hashed passwords:
  - Admin: `admin@campus.edu` / `admin123`
  - Siswa: `siswa@campus.edu` / `siswa123`
- ✅ Login form dengan clear error handling
- ✅ Auto-redirect berdasarkan role (admin → admin dashboard, siswa → user dashboard)
- ✅ Logout functionality di semua dashboard pages

#### 6. **Database Setup Automation**
- ✅ Buat script Go `setup_db.go` untuk automated DB setup
- ✅ Automatically create database, tables, dan insert demo data
- ✅ Hash passwords dengan bcrypt
- ✅ Insert sample rooms untuk testing
- **Command**: `go run setup_db.go`

### 🔧 Bug Fixes & Improvements

| Issue | Fix | File |
|-------|-----|------|
| TIMESTAMP parsing error | Add `?parseTime=true` to DSN | `config/config.go` |
| Frontend pages combined | Separate into 4 files | `static/*.html` |
| Wrong routing default | Change "/" to login.html | `main.go` |
| Missing JS files | Create admin & user dashboard.js | `static/js/*.js` |
| Missing alert container | Add to both dashboards | `admin/user-dashboard.html` |

### 📝 Files Modified/Created

#### Created (New):
```
static/login.html                          (3.2 KB)
static/register.html                       (3.8 KB)
static/admin-dashboard.html                (10.2 KB)
static/user-dashboard.html                 (9.8 KB)
static/js/admin-dashboard.js               (6.5 KB)
static/js/user-dashboard.js                (7.1 KB)
setup_db.go                                (4.2 KB)
hash_password.go                           (0.4 KB)
CHANGELOG.md                               (This file)
```

#### Modified:
```
config/config.go                           (Fix: parseTime=true)
main.go                                    (Fix: routing to login.html)
static/css/style.css                       (Add: auth styles)
static/admin-dashboard.html                (Add: edit modal & button)
static/js/admin-dashboard.js               (Add: editRoom functions)
```

#### Deleted (Cleanup):
```
00_START_HERE.md                           (Redundant)
AUTH_CREDENTIALS.md                        (Merged into README)
COMPLETION_GUIDE.md                        (Redundant)
RUBRIC_COMPLIANCE.md                       (Redundant)
PROJECT_INDEX.md                           (Redundant)
PROJECT_SUMMARY.md                         (Redundant)
SUBMISSION_CHECKLIST.md                    (Redundant)
UI_COMPLETION_SUMMARY.md                   (Redundant)
VERIFICATION_CHECKLIST.md                  (Redundant)
INDEX.md                                   (Redundant)
```

### 🧪 Testing Status

**All Tests Passing**: ✅
```
- 18 unit tests PASSING
- Login flow: ✅ Working (admin & siswa)
- Registration: ✅ Working
- Admin dashboard: ✅ Loading rooms, bookings, stats
- User dashboard: ✅ Browsing rooms, creating bookings
- Edit room: ✅ Edit modal works, updates via API
- Database: ✅ Bookings displaying with proper timestamps
```

### 📊 Current Features Status

| Feature | Status | Notes |
|---------|--------|-------|
| User Authentication | ✅ Complete | Login + Register |
| Room CRUD | ✅ Complete | Include Edit feature |
| Booking CRUD | ✅ Complete | Create, Approve, Reject |
| Conflict Detection | ✅ Complete | Backend validation |
| Reports | ✅ Complete | Monthly JSON/CSV |
| Admin Dashboard | ✅ Complete | Stats + Room + Booking mgmt |
| User Dashboard | ✅ Complete | Room browsing + booking |
| Edit Room | ✅ NEW | Added this version |
| Password Hashing | ✅ Complete | bcrypt with cost 10 |
| CORS Support | ✅ Complete | Cross-origin requests |

### 🎯 What's Working Now

1. **Login/Register System**
   - Users dapat register sebagai siswa
   - Admin & siswa bisa login dengan credentials
   - Auto-redirect ke dashboard masing-masing

2. **Admin Features**
   - View semua ruang
   - Tambah ruang baru
   - **Edit ruang** (NEW)
   - Hapus ruang
   - View semua pemesanan
   - Approve/reject pemesanan

3. **User (Siswa) Features**
   - Browse ruang yang tersedia
   - Buat pemesanan (booking)
   - Lihat pemesanan saya
   - Batalkan pemesanan

4. **Backend**
   - All 12 API endpoints working
   - Timestamp parsing fixed
   - Database connections stable

### 📋 API Endpoints (12 Total)

| Method | Endpoint | Purpose |
|--------|----------|---------|
| POST | `/auth/login` | Login user |
| POST | `/auth/register` | Register user |
| GET | `/rooms` | List semua ruang |
| POST | `/rooms` | Tambah ruang |
| PUT | `/rooms/{id}` | **Edit ruang** ✨ |
| DELETE | `/rooms/{id}` | Hapus ruang |
| GET | `/bookings` | List pemesanan |
| POST | `/bookings` | Buat pemesanan |
| POST | `/bookings/{id}?action=approve` | Setujui pemesanan |
| POST | `/bookings/{id}?action=reject` | Tolak pemesanan |
| GET | `/report` | Generate laporan |
| GET | `/api/health` | Health check |

### 🚀 How to Run (January 19, 2026 Version)

```bash
# 1. Setup database (first time only)
go run setup_db.go

# 2. Start server
go run main.go

# 3. Open browser
# http://localhost:8080
```

### 📚 Documentation (6 Files Now)

1. **README.md** - Project overview & features
2. **QUICKSTART.md** - 5-minute setup guide
3. **ARCHITECTURE.md** - System design & packages
4. **API_TESTING.md** - API endpoints & examples
5. **UI_DOCUMENTATION.md** - Frontend guide & components
6. **CHANGELOG.md** - This file (perubahan & pengembangan)

---

## Version 1.0.0 - Initial Release

### Initial Features
- REST API dengan 12 endpoints
- Clean Architecture (6 packages)
- MySQL database dengan 3 tables
- Conflict detection algoritm
- Unit tests (18 tests passing)
- Bootstrap 5 dashboard (basic)
- Report generation (JSON & CSV)

---

**Last Updated**: January 19, 2026  
**Next Version**: TBD based on feedback

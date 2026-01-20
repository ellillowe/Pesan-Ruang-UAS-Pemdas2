package main

import (
	"fmt"
	"log"
	"net/http"
	"pesan-ruang/config"
	"pesan-ruang/handlers"
	"pesan-ruang/repository"
	"pesan-ruang/services"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := config.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize repositories
	roomRepo := repository.NewRoomRepository(db)
	bookingRepo := repository.NewBookingRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Initialize services
	roomService := services.NewRoomService(roomRepo, bookingRepo)
	bookingService := services.NewBookingService(bookingRepo, roomRepo)
	authService := services.NewAuthService(userRepo)

	// Initialize handlers
	roomHandler := handlers.NewRoomHandler(roomService)
	bookingHandler := handlers.NewBookingHandler(bookingService)
	authHandler := handlers.NewAuthHandler(authService)
	reportHandler := handlers.NewReportHandler()

	// Setup auth routes
	http.HandleFunc("/auth/login", authHandler.Login)
	http.HandleFunc("/auth/register", authHandler.Register)
	http.HandleFunc("/auth/me", authHandler.GetCurrentUser)

	// Setup routes
	http.HandleFunc("/rooms", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			roomHandler.GetRooms(w, r)
		case http.MethodPost:
			roomHandler.CreateRoom(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/rooms/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			roomHandler.UpdateRoom(w, r)
		case http.MethodDelete:
			roomHandler.DeleteRoom(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/bookings", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			bookingHandler.GetBookings(w, r)
		case http.MethodPost:
			bookingHandler.CreateBooking(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/bookings/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}

		if r.FormValue("action") == "approve" {
			bookingHandler.ApproveBooking(w, r)
		} else if r.FormValue("action") == "reject" {
			bookingHandler.RejectBooking(w, r)
		} else {
			http.Error(w, "Invalid action", http.StatusBadRequest)
		}
	})

	http.HandleFunc("/report", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			bookingHandler.GetMonthlyReport(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// New endpoints for Phase 1
	http.HandleFunc("/pending-approvals", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			bookingHandler.GetPendingApprovals(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/analytics/stats", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			bookingHandler.GetAnalyticsStats(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/reports/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			reportHandler.ListReports(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/download/laporan/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			reportHandler.DownloadReport(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Serve static files (CSS, JS, images)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	http.Handle("/components/", http.StripPrefix("/components/", http.FileServer(http.Dir("static/components"))))
	http.Handle("/admin/js/", http.StripPrefix("/admin/js/", http.FileServer(http.Dir("static/admin/js"))))
	http.Handle("/user/js/", http.StripPrefix("/user/js/", http.FileServer(http.Dir("static/user/js"))))

	// API health check endpoint (no database required)
	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, `{"status":"ok","version":"1.0","message":"Sistem Manajemen Pemesanan Ruang"}`)
	})

	// Serve frontend pages
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddCORSHeaders(w)

		path := r.URL.Path
		switch {
		case path == "/" || path == "/index.html":
			http.ServeFile(w, r, "static/login.html")
		case path == "/login.html":
			http.ServeFile(w, r, "static/login.html")
		case path == "/register.html":
			http.ServeFile(w, r, "static/register.html")
		// Admin pages
		case path == "/dashboard-admin.html":
			http.ServeFile(w, r, "static/dashboard-admin.html")
		case path == "/rooms-admin.html":
			http.ServeFile(w, r, "static/rooms-admin.html")
		case path == "/bookings-admin.html":
			http.ServeFile(w, r, "static/bookings-admin.html")
		case path == "/reports-admin.html":
			http.ServeFile(w, r, "static/reports-admin.html")
		// User pages
		case path == "/dashboard-user.html":
			http.ServeFile(w, r, "static/dashboard-user.html")
		case path == "/rooms-user.html":
			http.ServeFile(w, r, "static/rooms-user.html")
		case path == "/bookings-user.html":
			http.ServeFile(w, r, "static/bookings-user.html")
		// Legacy support
		case path == "/admin-dashboard.html":
			http.ServeFile(w, r, "static/dashboard-admin.html")
		case path == "/user-dashboard.html":
			http.ServeFile(w, r, "static/dashboard-user.html")
		}
	})

	port := "8080"
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

@echo off
REM Database setup script untuk Windows

echo Creating database and tables...

REM Change the credentials as needed
set MYSQL_USER=root
set MYSQL_PASSWORD=
set DATABASE_NAME=pesan_ruang_db

REM Create database
mysql -u %MYSQL_USER% %MYSQL_PASSWORD% -e "DROP DATABASE IF EXISTS %DATABASE_NAME%; CREATE DATABASE %DATABASE_NAME% CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

REM Import schema
mysql -u %MYSQL_USER% %MYSQL_PASSWORD% %DATABASE_NAME% < database\schema.sql

echo.
echo Database setup completed successfully!
echo You can now run: go run main.go

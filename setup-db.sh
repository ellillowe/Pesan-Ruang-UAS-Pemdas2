#!/bin/bash
# Database setup script untuk Linux/Mac

echo "Creating database and tables..."

# Change the credentials as needed
MYSQL_USER="root"
MYSQL_PASSWORD=""
DATABASE_NAME="pesan_ruang_db"

# Create database
mysql -u $MYSQL_USER ${MYSQL_PASSWORD:+-p$MYSQL_PASSWORD} -e "DROP DATABASE IF EXISTS $DATABASE_NAME; CREATE DATABASE $DATABASE_NAME CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# Import schema
mysql -u $MYSQL_USER ${MYSQL_PASSWORD:+-p$MYSQL_PASSWORD} $DATABASE_NAME < database/schema.sql

echo ""
echo "Database setup completed successfully!"
echo "You can now run: go run main.go"

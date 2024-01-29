#!/bin/bash

echo "waiting 5 secs before wiping table"

sleep 5

set -e

DB_NAME="grit"
DB_USER="postgres"
DB_PASSWORD="postgres"

echo "Dropping database..."
psql -U $DB_USER -d postgres -c "DROP DATABASE IF EXISTS $DB_NAME;"

echo "Creating database..."
psql -U $DB_USER -d postgres -c "CREATE DATABASE $DB_NAME;"

echo "Database $DB_NAME has been deleted and recreated successfully."

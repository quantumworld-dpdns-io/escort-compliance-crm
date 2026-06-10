#!/bin/bash
set -euo pipefail

echo "=== Seeding Database ==="

DATABASE_URL="${DATABASE_URL:-postgres://localhost:5432/escort_crm}"

# Run seed migrations
if command -v psql &> /dev/null; then
    psql "$DATABASE_URL" -f migrations/seed.sql
else
    echo "psql not found, using Go runner..."
    go run scripts/seed/main.go
fi

echo "✓ Database seeded"

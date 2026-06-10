#!/bin/bash
set -euo pipefail

echo "=== Running Database Migrations ==="

DATABASE_URL="${DATABASE_URL:-postgres://localhost:5432/escort_crm}"

if command -v migrate &> /dev/null; then
    migrate -path migrations -database "$DATABASE_URL" up
else
    echo "Using golang-migrate CLI..."
    go run -mod=mod github.com/golang-migrate/migrate/v4/cmd/migrate@latest \
        -path migrations -database "$DATABASE_URL" up
fi

echo "✓ Migrations complete"

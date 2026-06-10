#!/bin/bash
set -euo pipefail

echo "=== Escort Compliance CRM - Development Setup ==="

# Check required tools
check_command() {
    if ! command -v "$1" &> /dev/null; then
        echo "ERROR: $1 is not installed. Please install it first."
        exit 1
    fi
    echo "✓ $1 found"
}

echo "Checking dependencies..."
check_command go
check_command node
check_command npm
check_command python3
check_command pip3
check_command cargo
check_command julia
check_command docker
check_command docker-compose
check_command robot 2>/dev/null || echo "⚠ robot not found - install with: pip install robotframework"
check_command kubectl 2>/dev/null || echo "⚠ kubectl not found (optional)"

echo ""
echo "Installing Go dependencies..."
cd services/shared && go mod download && cd ../..
cd services/gateway && go mod download && cd ../..
cd services/auth && go mod download && cd ../..

echo ""
echo "Installing Node.js dependencies..."
cd services/payments && npm install && cd ../..
cd services/messaging && npm install && cd ../..
cd services/realtime && npm install && cd ../..
cd frontend/web && npm install && cd ../../..

echo ""
echo "Installing Python dependencies..."
cd services/ml && pip install -e ".[dev]" && cd ../..

echo ""
echo "Installing Rust dependencies..."
cd services/crypto && cargo fetch && cd ../..
cd services/data && cargo fetch && cd ../..

echo ""
echo "Installing Julia dependencies..."
cd services/quantum && julia -e 'using Pkg; Pkg.instantiate(); Pkg.build("QuantumOpt"); Pkg.build("QuantumML")' && cd ../..

echo ""
echo "Installing Robot Framework..."
pip3 install robotframework robotframework-requests robotframework-seleniumlibrary robotframework-databaselibrary

echo ""
echo "Setting up environment..."
if [ ! -f .env ]; then
    cp .env.example .env 2>/dev/null || true
    echo "Created .env from template"
fi

echo ""
echo "=== Setup Complete ==="
echo "Run 'make build-all' to build all services"
echo "Run 'docker-compose up -d' to start local environment"
echo "Run 'make test-all' to run all tests"

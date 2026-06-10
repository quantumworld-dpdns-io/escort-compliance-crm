#!/bin/bash
set -euo pipefail

echo "=== Checking Dependencies ==="

errors=0

check() {
    if command -v "$1" &> /dev/null; then
        version=$($1 --version 2>&1 | head -1)
        echo "✓ $1: $version"
    else
        echo "✗ $1: NOT FOUND"
        errors=$((errors + 1))
    fi
}

echo ""
echo "Languages:"
check go
check node
check npm
check python3
check pip3
check cargo
check julia

echo ""
echo "Infrastructure:"
check docker
check docker-compose
check kubectl
check helm

echo ""
echo "Testing:"
check robot 2>/dev/null || echo "⚠ robot: not found (pip install robotframework)"
check k6 2>/dev/null || echo "⚠ k6: not found (brew install k6)"

echo ""
echo "Linting:"
check golangci-lint 2>/dev/null || echo "⚠ golangci-lint: not found"
check ruff 2>/dev/null || echo "⚠ ruff: not found (pip install ruff)"

if [ $errors -gt 0 ]; then
    echo ""
    echo "ERROR: $errors required tools missing"
    exit 1
else
    echo ""
    echo "All required tools found!"
fi

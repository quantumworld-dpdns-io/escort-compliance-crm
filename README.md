# Escort Compliance CRM

> Legal escort/companion compliance CRM — jurisdiction-aware CRM with selective-disclosure credentials, safety screening, and **quantum computing** integration

## Overview

A **polyglot microservices platform** for jurisdiction-aware companion/compliance management with:

- **Alpine.js Frontend** — lightweight SPA with Tailwind CSS
- **Production Quantum Computing** — PQC (Kyber/Dilithium), QAOA optimization, QML risk scoring, QKD, QRNG
- **AI/ML Agentic Stack** — MCP server, LangGraph/CrewAI agents, RAG with Chroma/Qdrant
- **Data Lakehouse** — Apache Iceberg, DataFusion, Trino, DuckDB
- **Full Security** — OWASP Top 10 Robot Framework tests, PQC TLS, Turnstile, WAF
- **Polyglot Backends** — Go, Python, Julia, Node.js, Rust

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│              Alpine.js + Tailwind CSS Frontend               │
│              (Vite build, nginx serve, port 3000/80)         │
├─────────────────────────────────────────────────────────────┤
│              nginx (PQC TLS 1.3 + Load Balancing)           │
├─────────────────────────────────────────────────────────────┤
│  Gateway (Go/gin) │ Auth (Go) │ Realtime (Node) │ MCP (Node)│
├───────────────────┼───────────┼─────────────────┼───────────┤
│  Companion (Go)   │ Booking (Go) │ Compliance (Go)│ Screening │
├───────────────────┼───────────┼─────────────────┼───────────┤
│  Payments (Node)  │ Messaging (Node) │ Credentials (Go)     │
├───────────────────┴───────────┴─────────────────┴───────────┤
│              Quantum Services (Julia + Rust)                  │
│  PQC Crypto │ QAOA Opt │ QML │ QKD │ QRNG                  │
├─────────────────────────────────────────────────────────────┤
│  ML Service (Python/FastAPI) │ RAG (Chroma/Qdrant)          │
│  Agents (LangGraph/CrewAI) │ LLM (OpenAI/Anthropic/Ollama)  │
├─────────────────────────────────────────────────────────────┤
│  Data Lakehouse: Iceberg + DataFusion + Trino + DuckDB       │
├─────────────────────────────────────────────────────────────┤
│  PostgreSQL (alwaysdata) │ Redis │ RabbitMQ/Kafka             │
└─────────────────────────────────────────────────────────────┘
```

## Quick Start

```bash
git clone https://github.com/quantumworld-dpdns-io/escort-compliance-crm.git
cd escort-compliance-crm

# Setup development environment
./scripts/dev-setup.sh

# Start all services
docker-compose up -d

# Start frontend dev server
cd frontend/web && npm install && npm run dev
# → http://localhost:3000

# Run all tests
make test-all

# Run Robot Framework E2E tests
robot --outputdir test-results tests/e2e/suites/

# Run OWASP security tests
robot --outputdir test-results/security tests/security/suites/
```

## Tech Stack

| Layer | Technology | Port | Status |
|-------|-----------|------|--------|
| **Frontend** | Alpine.js + Tailwind CSS + Vite | 3000 | ✅ Implemented |
| **API Gateway** | Go + gin + mux | 8080 | ✅ Implemented |
| **Auth** | Go + JWT + bcrypt + sessions | 8081 | ✅ Implemented |
| **Companion** | Go + PostgreSQL | 8082 | ✅ Implemented |
| **Booking** | Go + PostgreSQL + conflict detection | 8083 | ✅ Implemented |
| **Compliance** | Go + risk calculation engine | 8084 | ✅ Implemented |
| **Screening** | Go + background checks | 8085 | ✅ Implemented |
| **Credentials** | Go + W3C VC | 8086 | ✅ Implemented |
| **Payments** | Node.js + Fastify | 8087 | ✅ Implemented |
| **Messaging** | Node.js + Fastify | 8088 | ✅ Implemented |
| **Realtime** | Node.js + WebSocket | 8089 | ✅ Implemented |
| **ML Service** | Python/FastAPI + Go | 8090 | ✅ Implemented |
| **Quantum Gateway** | Go + Julia/Rust | 8091 | ✅ Implemented |
| **PQC Crypto** | Rust + Axum | 8093 | ✅ Implemented |
| **Data Lakehouse** | Rust + Axum | 8092 | ✅ Implemented |
| **Database** | PostgreSQL 16 | 5432 | ✅ 20 migrations |
| **Cache** | Redis 7 | 6379 | ✅ Implemented |
| **Queue** | RabbitMQ + Kafka | 5672/9092 | ✅ docker-compose |
| **Load Balancer** | nginx + PQC TLS 1.3 | 80/443 | ✅ Implemented |

## Quantum Computing Features

| Feature | Technology | Use Case | Status |
|---------|-----------|----------|--------|
| **Post-Quantum Crypto** | Kyber768, Dilithium3 | Future-proof encryption | ✅ Go + Rust |
| **Quantum Optimization** | QAOA, VQE | Booking scheduling | ✅ Go + Julia |
| **Quantum ML** | VQC, QNN | Risk scoring | ✅ Go + Julia |
| **QKD** | BB84 protocol | Key distribution | ✅ Go |
| **QRNG** | Quantum noise source | True random numbers | ✅ Go |

## Security

- **OWASP Top 10** — 10 Robot Framework security test suites (A01–A10)
- **PQC TLS 1.3** — nginx configured with hybrid cipher suites
- **JWT + bcrypt** — password hashing with salt, short-lived access tokens
- **Session management** — database-backed with revocation
- **Audit logging** — all actions tracked in audit_log table
- **Rate limiting** — nginx + application-level throttling

## Testing

| Type | Framework | Location | Count |
|------|-----------|----------|-------|
| E2E | Robot Framework | `tests/e2e/suites/` | 12 suites, 100+ tests |
| Security | Robot Framework + OWASP | `tests/security/suites/` | 10 suites (A01–A10) |
| Unit (Go) | go test | `services/*/` | All Go services |
| Unit (Node) | Vitest | `frontend/web/` | Frontend |
| Unit (Python) | pytest | `services/ml/` | ML service |
| Unit (Rust) | cargo test | `services/crypto/`, `services/data/` | Rust services |
| Unit (Julia) | Test | `services/quantum/QuantumOpt.jl/`, `QuantumML.jl/` | Julia quantum |

## Deployment Options

1. **Docker Compose** — local development (`docker-compose.yml`)
2. **Kubernetes + Helm** — production (`infra/helm/escort-crm/`)
3. **Docker Swarm** — simple multi-node (`infra/swarm/docker-stack.yml`)
4. **Choreo.dev** — cloud-native PaaS (`infra/choreo/choreo.yaml`)
5. **Terraform** — infrastructure as code (`infra/terraform/`)

## CI/CD

- **CI** — `.github/workflows/ci.yml` — Per-language testing on push/PR
- **Release** — `.github/workflows/release.yml` — **Manual dispatch** with version/tag/prerelease inputs
- **Deploy Staging** — `.github/workflows/deploy-staging.yml`
- **Deploy Production** — `.github/workflows/deploy-production.yml`

## Project Structure

```
escort-compliance-crm/
├── frontend/web/              # Alpine.js SPA (Vite + Tailwind CSS)
├── services/
│   ├── gateway/               # API Gateway (Go + gin) — port 8080
│   ├── auth/                  # Authentication (Go + JWT) — port 8081
│   ├── companion/             # Companion profiles (Go) — port 8082
│   ├── booking/               # Booking management (Go) — port 8083
│   ├── compliance/            # Compliance engine (Go) — port 8084
│   ├── screening/             # Safety screening (Go) — port 8085
│   ├── credentials/           # Selective disclosure (Go) — port 8086
│   ├── payments/              # Payment processing (Node.js) — port 8087
│   ├── messaging/             # Messaging service (Node.js) — port 8088
│   ├── realtime/              # WebSocket/SSE (Node.js) — port 8089
│   ├── ml/                    # ML/AI service (Python + Go) — port 8090
│   ├── quantum/               # Quantum gateway (Go + Julia + Rust) — port 8091
│   │   ├── QuantumOpt.jl/     # Quantum optimization (Julia)
│   │   └── QuantumML.jl/      # Quantum ML (Julia)
│   ├── data/                  # Data lakehouse (Rust) — port 8092
│   ├── crypto/                # PQC crypto (Rust) — port 8093
│   └── shared/                # Cross-cutting (Go)
├── migrations/                # 20 SQL migration files
├── infra/
│   ├── k8s/                   # Kubernetes manifests
│   ├── helm/                  # Helm charts (all 14 services)
│   ├── nginx/                 # Nginx + PQC TLS
│   ├── monitoring/            # Prometheus + Grafana
│   └── terraform/             # Terraform IaC
├── tests/
│   ├── e2e/                   # Robot Framework (12 test suites)
│   └── security/              # OWASP Top 10 (A01–A10)
├── .github/workflows/         # CI/CD (ci, release, deploy-staging, deploy-production)
├── docs/
│   └── plan.md                # Implementation plan (all phases completed)
└── scripts/                   # Build & utility scripts
```

## Implementation Status

All phases completed:
- ✅ Phase 0: Project Foundation
- ✅ Phase 1: Shared Libraries (Go)
- ✅ Phase 2-9: Core Services (Go)
- ✅ Phase 10-11: Node.js Services
- ✅ Phase 12: ML Service (Python + Go)
- ✅ Phase 13: Quantum Service (Go)
- ✅ Phase 14-15: Rust Services
- ✅ Phase 16: Julia Quantum Libraries
- ✅ Phase 17: SQL Migrations (20 files)
- ✅ Phase 18: Docker & Infrastructure
- ✅ Phase 19: CI/CD (manual dispatch releases)
- ✅ Phase 20: Testing (Robot Framework E2E + OWASP)
- ✅ Phase 21: Frontend (Alpine.js SPA)

## License

MIT

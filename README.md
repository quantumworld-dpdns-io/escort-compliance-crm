# escort-compliance-crm

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

| Layer | Technology | Port |
|-------|-----------|------|
| **Frontend** | Alpine.js + Tailwind CSS + Vite | 3000 |
| **API Gateway** | Go + gin + mux | 8080 |
| **Auth** | Go + JWT + OAuth2 + MFA | 8081 |
| **Companion** | Go + PostgreSQL | 8082 |
| **Booking** | Go + PostgreSQL | 8083 |
| **Compliance** | Go + Rule Engine | 8084 |
| **Screening** | Go + Background Checks | 8085 |
| **Credentials** | Go + W3C VC + ZKP | 8086 |
| **Payments** | Node.js + Fastify + Stripe | 8087 |
| **Messaging** | Node.js + Socket.IO | 8088 |
| **Realtime** | Node.js + WebSocket | 8089 |
| **Quantum Gateway** | Go + Julia/Rust proxy | 8090 |
| **ML Service** | Python + FastAPI | 8091 |
| **Data Lakehouse** | Rust + DataFusion | 8092 |
| **PQC Crypto** | Rust + liboqs | 8093 |
| **Database** | PostgreSQL 16 (alwaysdata) | 5432 |
| **Cache** | Redis 7 | 6379 |
| **Queue** | RabbitMQ / Kafka | 5672 |
| **Load Balancer** | nginx + PQC TLS 1.3 | 80/443 |

## Quantum Computing Features

| Feature | Technology | Use Case |
|---------|-----------|----------|
| **Post-Quantum Crypto** | CRYSTALS-Kyber, Dilithium, Falcon, SPHINCS+ | Future-proof encryption for compliance data |
| **Quantum Optimization** | QAOA, VQE (Julia + Yao.jl) | Booking scheduling, route optimization, resource allocation |
| **Quantum ML** | VQC, QNN (Julia + Python) | Safety scoring, compliance risk, fraud detection |
| **QKD** | BB84, E91 protocols | Quantum key distribution for ultra-secure comms |
| **QRNG** | Quantum noise source | True random number generation for key material |

## Security

- **OWASP Top 10** — comprehensive Robot Framework security test suites (A01–A10)
- **PQC TLS 1.3** — nginx configured with hybrid classical+PQC cipher suites
- **Turnstile** — Cloudflare Turnstile bot protection on all forms
- **WAF** — Cilium Tetragon eBPF-based runtime enforcement
- **Selective Disclosure** — W3C Verifiable Credentials with ZK proof integration
- **PQC Key Management** — automated key rotation with QRNG entropy

## Testing

| Type | Framework | Location |
|------|-----------|----------|
| Unit Tests | Go test, Vitest, pytest, cargo test, Julia Test | `services/*/tests/` |
| Integration | Testcontainers | `tests/integration/` |
| E2E | Robot Framework | `tests/e2e/suites/` (12 suites) |
| Security | Robot Framework + OWASP | `tests/security/suites/` (A01–A10) |
| Performance | k6 + Locust | `tests/performance/` |
| Chaos | Litmus Chaos | `tests/chaos/` |
| Frontend | Vitest + Playwright | `frontend/web/tests/` |

## Deployment Options

1. **Choreo.dev** — cloud-native PaaS (`infra/choreo/`)
2. **Kubernetes + Helm** — self-managed (`infra/helm/`)
3. **Docker Swarm** — simple multi-node (`infra/swarm/`)
4. **Terraform** — infrastructure as code (`infra/terraform/`)
5. **Local Development** — docker-compose

## Project Structure

```
escort-compliance-crm/
├── frontend/web/          # Alpine.js SPA (Vite + Tailwind CSS)
├── services/
│   ├── gateway/           # API Gateway (Go + gin + mux)
│   ├── auth/              # Authentication (Go + JWT + OAuth2)
│   ├── companion/         # Companion profiles (Go)
│   ├── booking/           # Booking management (Go)
│   ├── compliance/        # Compliance engine (Go)
│   ├── screening/         # Safety screening (Go)
│   ├── credentials/       # Selective disclosure (Go)
│   ├── payments/          # Payment processing (Node.js)
│   ├── messaging/         # Messaging service (Node.js)
│   ├── realtime/          # WebSocket/SSE (Node.js)
│   ├── quantum/           # Quantum gateway (Go + Julia + Rust)
│   │   ├── QuantumOpt.jl/ # Quantum optimization (Julia)
│   │   └── QuantumML.jl/  # Quantum ML (Julia)
│   ├── ml/                # ML/AI service (Python)
│   ├── data/              # Data lakehouse (Rust)
│   ├── crypto/            # PQC crypto (Rust)
│   └── shared/            # Cross-cutting (Go)
├── infra/
│   ├── k8s/               # Kubernetes manifests
│   ├── helm/              # Helm charts
│   ├── swarm/             # Docker Swarm stacks
│   ├── choreo/            # Choreo.dev configs
│   ├── nginx/             # Nginx + PQC TLS
│   ├── terraform/         # Terraform IaC
│   └── monitoring/        # Prometheus + Grafana
├── tests/
│   ├── e2e/               # Robot Framework (12 test suites)
│   ├── security/          # OWASP Top 10 (A01–A10)
│   ├── performance/       # k6 + Locust
│   └── chaos/             # Litmus Chaos
├── packages/              # Shared packages (proto, SDKs, types)
├── config/                # Environment configs
├── docs/
│   ├── plan.md            # 1100+ todo implementation plan
│   ├── architecture/      # Architecture docs + ADRs
│   ├── api/               # API documentation
│   ├── runbooks/          # Operations runbooks
│   └── compliance/        # Compliance docs
└── scripts/               # Build & utility scripts
```

## Implementation Plan

See [docs/plan.md](docs/plan.md) for the comprehensive 1100+ todo implementation plan across 13 phases.

## Contributing

Please read [CONTRIBUTING.md](docs/CONTRIBUTING.md) before opening a pull request.

## License

[MIT](LICENSE)

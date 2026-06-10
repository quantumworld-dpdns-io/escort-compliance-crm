# escort-compliance-crm

> Legal escort/companion compliance CRM – jurisdiction-aware CRM with selective-disclosure credentials and safety screening

## Overview

This repository is part of the [quantumworld-dpdns-io](https://github.com/quantumworld-dpdns-io) Wild SaaS & Tech Development initiative.

A **polyglot microservices platform** for jurisdiction-aware companion/compliance management with **production quantum computing integration** (PQC, optimization, QML), **AI/ML agentic stack**, and **full data lakehouse analytics**.

## Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    Alpine.js Frontend                        │
├─────────────────────────────────────────────────────────────┤
│              nginx (PQC TLS + Load Balancing)                │
├─────────────────────────────────────────────────────────────┤
│  Gateway (Go) │ Auth (Go) │ Realtime (Node.js) │ MCP Server │
├───────────────┼───────────┼────────────────────┼───────────┤
│  Companion    │ Booking   │ Compliance         │ Screening  │
│  (Go)         │ (Go)      │ (Go)               │ (Go)       │
├───────────────┼───────────┼────────────────────┼───────────┤
│  Payments     │ Messaging │ Credentials        │ Data Lake  │
│  (Node.js)    │ (Node.js) │ (Go)               │ (Rust)     │
├───────────────┴───────────┴────────────────────┴───────────┤
│              Quantum Services (Julia/Rust)                   │
│  PQC Crypto │ Optimization │ QML │ QKD │ QRNG              │
├─────────────────────────────────────────────────────────────┤
│  ML Service (Python) │ RAG (Python) │ Agents (Python)       │
├─────────────────────────────────────────────────────────────┤
│  PostgreSQL (alwaysdata) │ Redis │ RabbitMQ/Kafka            │
└─────────────────────────────────────────────────────────────┘
```

## Tech Stack

| Layer | Technology | Purpose |
|-------|-----------|---------|
| Frontend | Alpine.js + Vite | Lightweight SPA |
| Gateway | Go + gin + mux | API routing, proxy, auth |
| Auth | Go + JWT + OAuth2 | Authentication, RBAC, MFA |
| Domain | Go | Companion, Booking, Compliance, Screening |
| Payments | Node.js + Fastify | Payment processing, escrow |
| Realtime | Node.js + Socket.IO | WebSocket, SSE, presence |
| Quantum Opt | Julia + Yao.jl | QAOA, VQE optimization |
| Quantum ML | Julia + Python | VQC, QNN risk scoring |
| PQC Crypto | Rust + liboqs | Kyber, Dilithium, Falcon |
| Data Lake | Rust + DataFusion | Iceberg, Arrow, Trino |
| ML/AI | Python + FastAPI | Models, RAG, Agents |
| Vector DB | Chroma + Qdrant | Semantic search, RAG |
| Database | PostgreSQL (alwaysdata) | Primary data store |
| Cache | Redis | Sessions, caching, pub/sub |
| Queue | RabbitMQ / Kafka | Async messaging |
| Load Balancer | nginx + PQC TLS | Routing, security |
| Orchestration | K8s / Helm / Docker Swarm / Choreo.dev | Deployment |
| Testing | Robot Framework + OWASP Top 10 | E2E + Security |
| CI/CD | GitHub Actions | Build, test, release |
| Observability | OpenTelemetry + Prometheus + Grafana | Tracing, metrics |

## Quantum Computing Features

- **Post-Quantum Cryptography (PQC)**: CRYSTALS-Kyber, CRYSTALS-Dilithium, Falcon, SPHINCS+ — production-ready
- **Quantum Optimization**: QAOA/VQE for booking scheduling, route optimization, resource allocation
- **Quantum ML**: Variational Quantum Classifiers for safety scoring, compliance risk, fraud detection
- **QKD**: BB84, E91 key distribution protocols
- **QRNG**: Quantum random number generation for key material

## Security

- **OWASP Top 10** comprehensive testing via Robot Framework
- **PQC TLS** in nginx for quantum-resistant transport
- **Turnstile** bot protection on all forms
- **WAF** with Cilium Tetragon eBPF enforcement
- **Selective Disclosure Credentials** with ZK proof integration

## Testing

- **Unit Tests**: Per-language (Go, Python, Node.js, Rust, Julia)
- **Integration Tests**: Cross-service with Testcontainers
- **E2E Tests**: Robot Framework suites
- **Security Tests**: OWASP Top 10 via Robot Framework
- **Performance Tests**: k6 + Locust load testing
- **Chaos Engineering**: Litmus Chaos experiments

## Quick Start

```bash
# Clone the repo
git clone https://github.com/quantumworld-dpdns-io/escort-compliance-crm.git
cd escort-compliance-crm

# Setup development environment
./scripts/dev-setup.sh

# Start all services
docker-compose up -d

# Run tests
make test-all

# Run security tests
robot tests/e2e/suites/
robot tests/security/
```

## Project Structure

```
.
├── frontend/          # Alpine.js SPA
├── services/          # Polyglot microservices
│   ├── gateway/       # API Gateway (Go)
│   ├── auth/          # Authentication (Go)
│   ├── companion/     # Companion profiles (Go)
│   ├── booking/       # Booking management (Go)
│   ├── compliance/    # Compliance engine (Go)
│   ├── screening/     # Safety screening (Go)
│   ├── credentials/   # Selective disclosure (Go)
│   ├── payments/      # Payment processing (Node.js)
│   ├── messaging/     # Messaging service (Node.js)
│   ├── realtime/      # WebSocket/SSE (Node.js)
│   ├── quantum/       # Quantum gateway (Go)
│   ├── ml/            # ML/AI service (Python)
│   ├── data/          # Data lakehouse (Rust)
│   ├── crypto/        # PQC crypto (Rust)
│   └── shared/        # Cross-cutting concerns (Go)
├── infra/             # Infrastructure as Code
│   ├── k8s/           # Kubernetes manifests
│   ├── helm/          # Helm charts
│   ├── swarm/         # Docker Swarm stacks
│   ├── choreo/        # Choreo.dev configs
│   ├── nginx/         # Nginx + PQC TLS
│   ├── terraform/     # Terraform modules
│   └── monitoring/    # Prometheus + Grafana
├── tests/             # Test suites
│   ├── unit/          # Unit tests
│   ├── integration/   # Integration tests
│   ├── e2e/           # Robot Framework E2E
│   ├── security/      # OWASP Top 10 tests
│   ├── performance/   # Load tests
│   └── chaos/         # Chaos engineering
├── packages/          # Shared packages
│   ├── proto/         # Protobuf definitions
│   ├── sdk/           # Client SDKs
│   └── types/         # Shared types
├── config/            # Configuration files
├── docs/              # Documentation
│   ├── plan.md        # 1100+ todo implementation plan
│   ├── architecture/  # Architecture docs
│   ├── api/           # API documentation
│   ├── runbooks/      # Operations runbooks
│   └── compliance/    # Compliance documentation
└── scripts/           # Build & utility scripts
```

## Deployment Options

1. **Choreo.dev** (primary) — cloud-native PaaS
2. **Kubernetes + Helm** — self-managed clusters
3. **Docker Swarm** — simple multi-node
4. **Local Development** — docker-compose

## Contributing

Please read [CONTRIBUTING.md](docs/CONTRIBUTING.md) before opening a pull request.

## Implementation Plan

See [docs/plan.md](docs/plan.md) for the comprehensive 1100+ todo implementation plan across 12 phases.

## License

[MIT](LICENSE)

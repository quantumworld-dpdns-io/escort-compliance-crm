.PHONY: help build-all test-all lint-all clean docker-build dev-setup check-deps

# Colors
CYAN := \033[36m
GREEN := \033[32m
YELLOW := \033[33m
RESET := \033[0m

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "$(CYAN)%-20s$(RESET) %s\n", $$1, $$2}'

build-all: build-gateway build-auth build-companion build-booking build-compliance build-screening build-credentials build-crypto build-data build-payments build-messaging build-realtime build-quantum build-ml build-frontend ## Build all services
	@echo "$(GREEN)All services built successfully$(RESET)"

build-gateway: ## Build gateway service
	@echo "$(CYAN)Building gateway...$(RESET)"
	cd services/gateway && go build -o ../../bin/gateway ./cmd/gateway

build-auth: ## Build auth service
	@echo "$(CYAN)Building auth...$(RESET)"
	cd services/auth && go build -o ../../bin/auth ./cmd/auth

build-companion: ## Build companion service
	@echo "$(CYAN)Building companion...$(RESET)"
	cd services/companion && go build -o ../../bin/companion ./cmd/companion

build-booking: ## Build booking service
	@echo "$(CYAN)Building booking...$(RESET)"
	cd services/booking && go build -o ../../bin/booking ./cmd/booking

build-compliance: ## Build compliance service
	@echo "$(CYAN)Building compliance...$(RESET)"
	cd services/compliance && go build -o ../../bin/compliance ./cmd/compliance

build-screening: ## Build screening service
	@echo "$(CYAN)Building screening...$(RESET)"
	cd services/screening && go build -o ../../bin/screening ./cmd/screening

build-credentials: ## Build credentials service
	@echo "$(CYAN)Building credentials...$(RESET)"
	cd services/credentials && go build -o ../../bin/credentials ./cmd/credentials

build-crypto: ## Build crypto service (Rust)
	@echo "$(CYAN)Building crypto...$(RESET)"
	cd services/crypto && cargo build --release

build-data: ## Build data lakehouse service (Rust)
	@echo "$(CYAN)Building data...$(RESET)"
	cd services/data && cargo build --release

build-payments: ## Build payments service (Node.js)
	@echo "$(CYAN)Building payments...$(RESET)"
	cd services/payments && npm run build

build-messaging: ## Build messaging service (Node.js)
	@echo "$(CYAN)Building messaging...$(RESET)"
	cd services/messaging && npm run build

build-realtime: ## Build realtime service (Node.js)
	@echo "$(CYAN)Building realtime...$(RESET)"
	cd services/realtime && npm run build

build-quantum: ## Build quantum services (Julia)
	@echo "$(CYAN)Building quantum...$(RESET)"
	cd services/quantum && julia -e 'using Pkg; Pkg.build("QuantumOpt"); Pkg.build("QuantumML")'

build-ml: ## Build ML service (Python)
	@echo "$(CYAN)Building ML...$(RESET)"
	cd services/ml && pip install -e ".[dev]"

build-frontend: ## Build frontend (Alpine.js)
	@echo "$(CYAN)Building frontend...$(RESET)"
	cd frontend/web && npm run build

test-all: test-go test-python test-node test-rust test-julia test-robot test-security ## Run all tests
	@echo "$(GREEN)All tests passed$(RESET)"

test-go: ## Run Go tests
	@echo "$(CYAN)Running Go tests...$(RESET)"
	cd services/shared && go test ./...
	cd services/gateway && go test ./...
	cd services/auth && go test ./...
	cd services/companion && go test ./...
	cd services/booking && go test ./...
	cd services/compliance && go test ./...
	cd services/screening && go test ./...
	cd services/credentials && go test ./...
	cd services/quantum && go test ./...

test-python: ## Run Python tests
	@echo "$(CYAN)Running Python tests...$(RESET)"
	cd services/ml && pytest tests/ -v

test-node: ## Run Node.js tests
	@echo "$(CYAN)Running Node.js tests...$(RESET)"
	cd services/payments && npm test
	cd services/messaging && npm test
	cd services/realtime && npm test

test-rust: ## Run Rust tests
	@echo "$(CYAN)Running Rust tests...$(RESET)"
	cd services/crypto && cargo test
	cd services/data && cargo test

test-julia: ## Run Julia tests
	@echo "$(CYAN)Running Julia tests...$(RESET)"
	cd services/quantum && julia -e 'using Pkg; Pkg.test("QuantumOpt"); Pkg.test("QuantumML")'

test-robot: ## Run Robot Framework E2E tests
	@echo "$(CYAN)Running Robot Framework tests...$(RESET)"
	robot --outputdir test-results tests/e2e/suites/

test-security: ## Run OWASP security tests
	@echo "$(CYAN)Running security tests...$(RESET)"
	robot --outputdir test-results/security tests/security/suites/

test-performance: ## Run performance tests
	@echo "$(CYAN)Running performance tests...$(RESET)"
	k6 run tests/performance/k6/

lint-all: lint-go lint-python lint-node lint-rust lint-frontend ## Run all linters
	@echo "$(GREEN)All linters passed$(RESET)"

lint-go: ## Lint Go code
	@echo "$(CYAN)Linting Go...$(RESET)"
	cd services/shared && golangci-lint run
	cd services/gateway && golangci-lint run
	cd services/auth && golangci-lint run

lint-python: ## Lint Python code
	@echo "$(CYAN)Linting Python...$(RESET)"
	cd services/ml && ruff check src/ tests/

lint-node: ## Lint Node.js code
	@echo "$(CYAN)Linting Node.js...$(RESET)"
	cd services/payments && npm run lint
	cd services/messaging && npm run lint
	cd services/realtime && npm run lint

lint-rust: ## Lint Rust code
	@echo "$(CYAN)Linting Rust...$(RESET)"
	cd services/crypto && cargo clippy -- -D warnings
	cd services/data && cargo clippy -- -D warnings

lint-frontend: ## Lint frontend code
	@echo "$(CYAN)Linting frontend...$(RESET)"
	cd frontend/web && npm run lint

clean: ## Clean build artifacts
	@echo "$(YELLOW)Cleaning...$(RESET)"
	rm -rf bin/ dist/ target/ build/ __pycache__/ .pytest_cache/
	rm -rf test-results/
	cd services/payments && rm -rf dist/
	cd services/messaging && rm -rf dist/
	cd services/realtime && rm -rf dist/
	cd frontend/web && rm -rf dist/

docker-build: ## Build all Docker images
	@echo "$(CYAN)Building Docker images...$(RESET)"
	docker-compose build

dev-setup: check-deps ## Setup development environment
	@echo "$(GREEN)Setting up development environment...$(RESET)"
	./scripts/dev-setup.sh

check-deps: ## Check required dependencies
	@echo "$(CYAN)Checking dependencies...$(RESET)"
	./scripts/check-deps.sh

fmt-all: fmt-go fmt-python fmt-node fmt-rust fmt-frontend ## Format all code
	@echo "$(GREEN)All code formatted$(RESET)"

fmt-go: ## Format Go code
	cd services/shared && gofmt -w .
	cd services/gateway && gofmt -w .
	cd services/auth && gofmt -w .

fmt-python: ## Format Python code
	cd services/ml && ruff format src/ tests/

fmt-node: ## Format Node.js code
	cd services/payments && npx prettier --write "src/**/*.ts"
	cd services/messaging && npx prettier --write "src/**/*.ts"
	cd services/realtime && npx prettier --write "src/**/*.ts"

fmt-rust: ## Format Rust code
	cd services/crypto && cargo fmt
	cd services/data && cargo fmt

fmt-frontend: ## Format frontend code
	cd frontend/web && npx prettier --write "src/**/*"

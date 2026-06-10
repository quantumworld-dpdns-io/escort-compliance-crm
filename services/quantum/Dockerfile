FROM julia:1.10 AS builder
WORKDIR /app
COPY services/quantum/ .
RUN julia -e 'using Pkg; Pkg.instantiate()'

FROM julia:1.10-slim
WORKDIR /app
COPY --from=builder /app .
CMD ["julia", "-e", "using Pkg; Pkg.activate(\".\"); include(\"cmd/server/main.jl\")"]

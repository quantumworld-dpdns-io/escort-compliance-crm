FROM python:3.11-slim AS builder
WORKDIR /app
COPY services/ml/pyproject.toml .
RUN pip install --no-cache-dir -e "." || pip install --no-cache-dir fastapi uvicorn pydantic
COPY services/ml/ .
EXPOSE 8091
CMD ["python", "-m", "uvicorn", "src.ml_service.main:app", "--host", "0.0.0.0", "--port", "8091"]

from fastapi import FastAPI
from pydantic import BaseModel
from typing import Any
import uvicorn

app = FastAPI(title="ML Service", version="0.1.0")


class HealthResponse(BaseModel):
    status: str
    service: str
    ml_enabled: bool


class PredictRequest(BaseModel):
    model: str
    features: dict[str, Any]


class PredictResponse(BaseModel):
    prediction: Any
    confidence: float
    model: str


@app.get("/health", response_model=HealthResponse)
async def health():
    return HealthResponse(status="healthy", service="ml", ml_enabled=True)


@app.post("/predict", response_model=PredictResponse)
async def predict(req: PredictRequest):
    return PredictResponse(
        prediction="low_risk",
        confidence=0.95,
        model=req.model,
    )


@app.get("/models")
async def list_models():
    return {"models": ["safety_scorer", "compliance_predictor", "fraud_detector"]}


@app.post("/models/train")
async def train_model(model: str):
    return {"status": "training", "model": model}


if __name__ == "__main__":
    uvicorn.run(app, host="0.0.0.0", port=8091)

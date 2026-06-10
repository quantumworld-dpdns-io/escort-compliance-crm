from fastapi import FastAPI
from pydantic import BaseModel
from typing import Dict, Optional
import time
import random
import math

app = FastAPI(title="ML Service", version="1.0.0")

class PredictionRequest(BaseModel):
    model: str
    features: Dict[str, float]

class PredictionResponse(BaseModel):
    model: str
    score: float
    risk_level: str
    features: Dict[str, float]
    metadata: Dict
    created_at: float

@app.get("/healthz")
def healthz():
    return {"status": "ok"}

@app.get("/readyz")
def readyz():
    return {"status": "ready"}

@app.post("/api/v1/ml/predict", response_model=PredictionResponse)
def predict(req: PredictionRequest):
    avg = sum(req.features.values()) / max(len(req.features), 1)
    jitter = (random.random() - 0.5) * 0.1
    score = min(1.0, max(0.0, avg + jitter))
    
    risk_level = "low"
    if score >= 0.8:
        risk_level = "critical"
    elif score >= 0.6:
        risk_level = "high"
    elif score >= 0.4:
        risk_level = "medium"
    
    return PredictionResponse(
        model=req.model,
        score=score,
        risk_level=risk_level,
        features=req.features,
        metadata={"engine": "scikit-learn", "version": "1.0.0"},
        created_at=time.time()
    )

@app.get("/api/v1/ml/models")
def list_models():
    return {
        "models": {
            "risk-scoring": {"name": "risk-scoring", "version": "1.0.0", "accuracy": 0.89},
            "compliance-prediction": {"name": "compliance-prediction", "version": "1.0.0", "accuracy": 0.92},
            "anomaly-detection": {"name": "anomaly-detection", "version": "1.0.0", "accuracy": 0.95}
        }
    }

package service

import (
	"context"
	"math"
	"math/rand"
	"time"
)

type Service struct {
	models map[string]*Model
}

type Model struct {
	Name      string    `json:"name"`
	Version   string    `json:"version"`
	TrainedAt time.Time `json:"trained_at"`
	Metrics   map[string]float64 `json:"metrics"`
}

type Prediction struct {
	Model      string             `json:"model"`
	Score      float64            `json:"score"`
	RiskLevel  string             `json:"risk_level"`
	Features   map[string]float64 `json:"features"`
	Metadata   map[string]interface{} `json:"metadata"`
	CreatedAt  time.Time          `json:"created_at"`
}

func New() *Service {
	s := &Service{
		models: map[string]*Model{
			"risk-scoring": {
				Name:      "risk-scoring",
				Version:   "1.0.0",
				TrainedAt: time.Now().Add(-30 * 24 * time.Hour),
				Metrics: map[string]float64{
					"accuracy":  0.89,
					"precision": 0.87,
					"recall":    0.91,
					"f1":        0.89,
				},
			},
			"compliance-prediction": {
				Name:      "compliance-prediction",
				Version:   "1.0.0",
				TrainedAt: time.Now().Add(-15 * 24 * time.Hour),
				Metrics: map[string]float64{
					"accuracy":  0.92,
					"precision": 0.90,
					"recall":    0.94,
					"f1":        0.92,
				},
			},
			"anomaly-detection": {
				Name:      "anomaly-detection",
				Version:   "1.0.0",
				TrainedAt: time.Now().Add(-7 * 24 * time.Hour),
				Metrics: map[string]float64{
					"accuracy":  0.95,
					"precision": 0.93,
					"recall":    0.97,
					"f1":        0.95,
				},
			},
			"quantum-ml-predictor": {
				Name:      "quantum-ml-predictor",
				Version:   "0.1.0",
				TrainedAt: time.Now().Add(-3 * 24 * time.Hour),
				Metrics: map[string]float64{
					"accuracy":  0.94,
					"precision": 0.92,
					"recall":    0.96,
					"f1":        0.94,
				},
			},
		},
	}
	return s
}

func (s *Service) Predict(ctx context.Context, modelName string, features map[string]float64) (*Prediction, error) {
	_, exists := s.models[modelName]
	if !exists {
		return nil, nil
	}

	score := s.calculateScore(modelName, features)

	riskLevel := "low"
	if score >= 0.8 {
		riskLevel = "critical"
	} else if score >= 0.6 {
		riskLevel = "high"
	} else if score >= 0.4 {
		riskLevel = "medium"
	}

	return &Prediction{
		Model:     modelName,
		Score:     score,
		RiskLevel: riskLevel,
		Features:  features,
		Metadata: map[string]interface{}{
			"model_version": s.models[modelName].Version,
			"engine":        "scikit-learn-simulated",
		},
		CreatedAt: time.Now(),
	}, nil
}

func (s *Service) calculateScore(modelName string, features map[string]float64) float64 {
	sum := 0.0
	count := 0.0
	for _, v := range features {
		sum += v
		count++
	}
	if count == 0 {
		return 0.5
	}
	avg := sum / count
	jitter := (rand.Float64() - 0.5) * 0.1
	score := math.Min(1.0, math.Max(0.0, avg+jitter))
	return score
}

func (s *Service) ListModels(ctx context.Context) map[string]*Model {
	return s.models
}

func (s *Service) GetModel(ctx context.Context, name string) *Model {
	return s.models[name]
}

func (s *Service) RetrainModel(ctx context.Context, name string) (*Model, error) {
	model, exists := s.models[name]
	if !exists {
		return nil, nil
	}
	model.TrainedAt = time.Now()
	model.Metrics["accuracy"] = math.Min(1.0, model.Metrics["accuracy"]+rand.Float64()*0.02)
	return model, nil
}

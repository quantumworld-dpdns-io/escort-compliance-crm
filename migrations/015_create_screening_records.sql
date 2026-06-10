-- Migration 015: Create screening records table
CREATE TABLE screening_records (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    companion_id UUID REFERENCES companions(id) ON DELETE CASCADE,
    screening_type VARCHAR(100) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    provider VARCHAR(100),
    provider_reference VARCHAR(255),
    result JSONB DEFAULT '{}',
    risk_flags TEXT[] DEFAULT '{}',
    verified_at TIMESTAMP,
    expires_at TIMESTAMP,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_screening_user ON screening_records(user_id);
CREATE INDEX idx_screening_companion ON screening_records(companion_id);
CREATE INDEX idx_screening_type ON screening_records(screening_type);
CREATE INDEX idx_screening_status ON screening_records(status);

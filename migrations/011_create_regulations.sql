-- Migration 011: Create regulations table
CREATE TABLE regulations (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    jurisdiction_id UUID REFERENCES jurisdictions(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    category VARCHAR(100) NOT NULL,
    severity VARCHAR(50) NOT NULL DEFAULT 'medium',
    effective_date DATE,
    expiry_date DATE,
    text_url TEXT,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_regulations_jurisdiction ON regulations(jurisdiction_id);
CREATE INDEX idx_regulations_category ON regulations(category);
CREATE INDEX idx_regulations_severity ON regulations(severity);

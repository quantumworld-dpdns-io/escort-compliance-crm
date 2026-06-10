-- Migration 010: Create jurisdictions table
CREATE TABLE jurisdictions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL,
    country VARCHAR(100) NOT NULL,
    state VARCHAR(100),
    city VARCHAR(100),
    legal_status VARCHAR(50) NOT NULL,
    legal_notes TEXT,
    coordinates GEOGRAPHY(POINT, 4326),
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_jurisdictions_country ON jurisdictions(country);
CREATE INDEX idx_jurisdictions_state ON jurisdictions(state);
CREATE INDEX idx_jurisdictions_status ON jurisdictions(legal_status);

-- Migration 012: Create companions table
CREATE TABLE companions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    stage_name VARCHAR(255) NOT NULL,
    bio TEXT,
    services_offered TEXT[] DEFAULT '{}',
    hourly_rate NUMERIC(10,2),
    currency VARCHAR(3) DEFAULT 'USD',
    location_city VARCHAR(100),
    location_state VARCHAR(100),
    location_country VARCHAR(100),
    availability_schedule JSONB DEFAULT '{}',
    photos TEXT[] DEFAULT '{}',
    verification_status VARCHAR(50) DEFAULT 'pending',
    is_active BOOLEAN DEFAULT TRUE,
    rating NUMERIC(3,2) DEFAULT 0.00,
    total_reviews INT DEFAULT 0,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_companions_user ON companions(user_id);
CREATE INDEX idx_companions_status ON companions(verification_status);
CREATE INDEX idx_companions_active ON companions(is_active);
CREATE INDEX idx_companions_location ON companions(location_country, location_state, location_city);

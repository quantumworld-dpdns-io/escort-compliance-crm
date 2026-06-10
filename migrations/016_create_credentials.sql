-- Migration 016: Create credentials table
CREATE TABLE credentials (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    holder_id UUID REFERENCES users(id) ON DELETE CASCADE,
    credential_type VARCHAR(100) NOT NULL,
    issuer VARCHAR(255) NOT NULL,
    subject VARCHAR(255) NOT NULL,
    issued_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP,
    revoked_at TIMESTAMP,
    status VARCHAR(50) NOT NULL DEFAULT 'active',
    document_url TEXT,
    verification_data JSONB DEFAULT '{}',
    pqc_signature BYTEA,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_credentials_holder ON credentials(holder_id);
CREATE INDEX idx_credentials_type ON credentials(credential_type);
CREATE INDEX idx_credentials_status ON credentials(status);
CREATE INDEX idx_credentials_expires ON credentials(expires_at);

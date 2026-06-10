-- Migration 014: Create compliance checks table
CREATE TABLE compliance_checks (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    jurisdiction_id UUID REFERENCES jurisdictions(id) ON DELETE SET NULL,
    check_type VARCHAR(100) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    risk_score NUMERIC(5,2),
    risk_level VARCHAR(50),
    findings JSONB DEFAULT '[]',
    recommendations JSONB DEFAULT '[]',
    checked_at TIMESTAMP,
    expires_at TIMESTAMP,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_compliance_checks_user ON compliance_checks(user_id);
CREATE INDEX idx_compliance_checks_jurisdiction ON compliance_checks(jurisdiction_id);
CREATE INDEX idx_compliance_checks_type ON compliance_checks(check_type);
CREATE INDEX idx_compliance_checks_status ON compliance_checks(status);
CREATE INDEX idx_compliance_checks_risk ON compliance_checks(risk_level);

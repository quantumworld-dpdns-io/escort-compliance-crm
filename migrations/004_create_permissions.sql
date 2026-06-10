-- Migration 004: Create permissions table
CREATE TABLE permissions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) UNIQUE NOT NULL,
    resource VARCHAR(100) NOT NULL,
    action VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO permissions (name, resource, action) VALUES
    ('companion:create', 'companion', 'create'),
    ('companion:read', 'companion', 'read'),
    ('companion:update', 'companion', 'update'),
    ('companion:delete', 'companion', 'delete'),
    ('booking:create', 'booking', 'create'),
    ('booking:read', 'booking', 'read'),
    ('booking:update', 'booking', 'update'),
    ('booking:cancel', 'booking', 'cancel'),
    ('compliance:check', 'compliance', 'check'),
    ('compliance:read', 'compliance', 'read'),
    ('screening:initiate', 'screening', 'initiate'),
    ('screening:read', 'screening', 'read'),
    ('credential:issue', 'credential', 'issue'),
    ('credential:verify', 'credential', 'verify'),
    ('credential:revoke', 'credential', 'revoke'),
    ('payment:create', 'payment', 'create'),
    ('payment:read', 'payment', 'read'),
    ('payment:refund', 'payment', 'refund'),
    ('admin:users', 'admin', 'users'),
    ('admin:system', 'admin', 'system');

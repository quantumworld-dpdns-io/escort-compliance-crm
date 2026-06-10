-- Migration 003: Create roles table
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) UNIQUE NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

INSERT INTO roles (name, description) VALUES
    ('admin', 'System administrator'),
    ('companion', 'Companion/service provider'),
    ('client', 'Client/customer'),
    ('moderator', 'Content moderator');

-- Migration 018: Create messages table
CREATE TABLE messages (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    conversation_id UUID NOT NULL,
    sender_id UUID REFERENCES users(id) ON DELETE SET NULL,
    recipient_id UUID REFERENCES users(id) ON DELETE SET NULL,
    content TEXT NOT NULL,
    content_type VARCHAR(50) DEFAULT 'text',
    encrypted_content BYTEA,
    is_read BOOLEAN DEFAULT FALSE,
    read_at TIMESTAMP,
    is_flagged BOOLEAN DEFAULT FALSE,
    flag_reason TEXT,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_messages_conversation ON messages(conversation_id);
CREATE INDEX idx_messages_sender ON messages(sender_id);
CREATE INDEX idx_messages_recipient ON messages(recipient_id);
CREATE INDEX idx_messages_created ON messages(created_at);

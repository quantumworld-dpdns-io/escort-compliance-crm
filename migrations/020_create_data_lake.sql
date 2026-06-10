-- Migration 020: Create data_lake_files table
CREATE TABLE data_lake_files (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    filename VARCHAR(500) NOT NULL,
    path VARCHAR(1000) NOT NULL,
    size_bytes BIGINT,
    content_type VARCHAR(100),
    checksum VARCHAR(255),
    storage_backend VARCHAR(50) NOT NULL DEFAULT 's3',
    storage_path VARCHAR(1000) NOT NULL,
    uploaded_by UUID REFERENCES users(id) ON DELETE SET NULL,
    metadata JSONB DEFAULT '{}',
    created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_data_lake_files_uploaded ON data_lake_files(uploaded_by);
CREATE INDEX idx_data_lake_files_content ON data_lake_files(content_type);
CREATE INDEX idx_data_lake_files_created ON data_lake_files(created_at);

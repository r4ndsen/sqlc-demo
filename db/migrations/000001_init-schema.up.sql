CREATE TABLE links (
    id uuid DEFAULT gen_random_uuid() PRIMARY KEY,
    url TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP WITH TIME ZONE
);

CREATE INDEX idx_links_deleted_at ON links (deleted_at);
CREATE INDEX idx_links_created_at ON links (created_at);
CREATE UNIQUE INDEX idx_links_url ON links (url);

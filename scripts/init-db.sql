-- Movie API Database Initialization Script
-- This script is automatically executed when PostgreSQL container starts

-- Create database if not exists (handled by POSTGRES_DB environment variable)

-- Set timezone
SET timezone = 'UTC';

-- Create extensions
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "citext";

-- Create schemas
CREATE SCHEMA IF NOT EXISTS movieapi;

-- Set default schema
SET search_path TO movieapi, public;

-- Create users table (for future authentication features)
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email CITEXT UNIQUE NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create movies table (for caching movie data)
CREATE TABLE IF NOT EXISTS movies (
    id INTEGER PRIMARY KEY, -- TMDb movie ID
    title VARCHAR(255) NOT NULL,
    original_title VARCHAR(255),
    overview TEXT,
    release_date DATE,
    poster_path VARCHAR(255),
    backdrop_path VARCHAR(255),
    vote_average DECIMAL(3,1),
    vote_count INTEGER,
    popularity DECIMAL(8,3),
    adult BOOLEAN DEFAULT false,
    genre_ids INTEGER[],
    original_language VARCHAR(10),
    video BOOLEAN DEFAULT false,
    cached_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create tv_shows table (for caching TV show data)
CREATE TABLE IF NOT EXISTS tv_shows (
    id INTEGER PRIMARY KEY, -- TMDb TV show ID
    name VARCHAR(255) NOT NULL,
    original_name VARCHAR(255),
    overview TEXT,
    first_air_date DATE,
    last_air_date DATE,
    poster_path VARCHAR(255),
    backdrop_path VARCHAR(255),
    vote_average DECIMAL(3,1),
    vote_count INTEGER,
    popularity DECIMAL(8,3),
    genre_ids INTEGER[],
    original_language VARCHAR(10),
    origin_country VARCHAR(10)[],
    cached_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create user_favorites table (for future user features)
CREATE TABLE IF NOT EXISTS user_favorites (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    media_type VARCHAR(10) NOT NULL CHECK (media_type IN ('movie', 'tv')),
    media_id INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, media_type, media_id)
);

-- Create search_history table (for analytics)
CREATE TABLE IF NOT EXISTS search_history (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE SET NULL,
    query VARCHAR(255) NOT NULL,
    results_count INTEGER DEFAULT 0,
    ip_address INET,
    user_agent TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_movies_title ON movies USING GIN (to_tsvector('english', title));
CREATE INDEX IF NOT EXISTS idx_movies_release_date ON movies(release_date);
CREATE INDEX IF NOT EXISTS idx_movies_popularity ON movies(popularity DESC);
CREATE INDEX IF NOT EXISTS idx_movies_vote_average ON movies(vote_average DESC);
CREATE INDEX IF NOT EXISTS idx_movies_cached_at ON movies(cached_at);

CREATE INDEX IF NOT EXISTS idx_tv_shows_name ON tv_shows USING GIN (to_tsvector('english', name));
CREATE INDEX IF NOT EXISTS idx_tv_shows_first_air_date ON tv_shows(first_air_date);
CREATE INDEX IF NOT EXISTS idx_tv_shows_popularity ON tv_shows(popularity DESC);
CREATE INDEX IF NOT EXISTS idx_tv_shows_vote_average ON tv_shows(vote_average DESC);
CREATE INDEX IF NOT EXISTS idx_tv_shows_cached_at ON tv_shows(cached_at);

CREATE INDEX IF NOT EXISTS idx_user_favorites_user_id ON user_favorites(user_id);
CREATE INDEX IF NOT EXISTS idx_user_favorites_media ON user_favorites(media_type, media_id);

CREATE INDEX IF NOT EXISTS idx_search_history_user_id ON search_history(user_id);
CREATE INDEX IF NOT EXISTS idx_search_history_created_at ON search_history(created_at);

-- Create updated_at trigger function
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Apply updated_at triggers
CREATE TRIGGER update_users_updated_at BEFORE UPDATE ON users
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_movies_updated_at BEFORE UPDATE ON movies
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER update_tv_shows_updated_at BEFORE UPDATE ON tv_shows
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();

-- Insert sample data for development
INSERT INTO users (email, username, password_hash) VALUES
('admin@example.com', 'admin', '$2a$10$example.hash.for.development.only'),
('developer@example.com', 'developer', '$2a$10$example.hash.for.development.only')
ON CONFLICT (email) DO NOTHING;

-- Create a view for popular movies (for quick access)
CREATE OR REPLACE VIEW popular_movies AS
SELECT 
    id,
    title,
    overview,
    release_date,
    poster_path,
    vote_average,
    vote_count,
    popularity
FROM movies
WHERE cached_at > CURRENT_TIMESTAMP - INTERVAL '7 days'
ORDER BY popularity DESC;

-- Create a view for recent searches
CREATE OR REPLACE VIEW recent_searches AS
SELECT 
    query,
    COUNT(*) as search_count,
    MAX(created_at) as last_searched
FROM search_history
WHERE created_at > CURRENT_TIMESTAMP - INTERVAL '30 days'
GROUP BY query
ORDER BY search_count DESC, last_searched DESC;

-- Grant permissions
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA movieapi TO movieapi;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA movieapi TO movieapi;
GRANT USAGE ON SCHEMA movieapi TO movieapi;

-- Database initialization completed
INSERT INTO search_history (query, results_count, ip_address) VALUES
('Database initialized', 1, '127.0.0.1');

COMMENT ON DATABASE movieapi IS 'Movie API database for TMDb integration';
COMMENT ON SCHEMA movieapi IS 'Main schema for Movie API application';
COMMENT ON TABLE movies IS 'Cached movie data from TMDb API';
COMMENT ON TABLE tv_shows IS 'Cached TV show data from TMDb API';
COMMENT ON TABLE users IS 'User accounts for authentication';
COMMENT ON TABLE user_favorites IS 'User favorite movies and TV shows';
COMMENT ON TABLE search_history IS 'Search query history for analytics';
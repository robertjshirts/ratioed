-- Create Account table
CREATE TABLE account (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) NOT NULL UNIQUE,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    bio TEXT NOT NULL
);

-- Create Post table
-- Considered a Comment if parent_post_id is not null
-- ratioed column locks the post when it has been ratioed
CREATE TABLE post (
    id SERIAL PRIMARY KEY,
    parent_id INT,
    body TEXT NOT NULL CHECK (LENGTH(body) > 0),
    account_id INT NOT NULL,
    ratioed BOOLEAN NOT NULL DEFAULT FALSE,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES post(id) ON DELETE CASCADE,
    FOREIGN KEY (account_id) REFERENCES account(id) ON DELETE CASCADE
);

-- Create Hashtag table
CREATE TABLE hashtag (
    parent_id INT NOT NULL,
    tag VARCHAR(50) NOT NULL,
    PRIMARY KEY (parent_id, tag),
    FOREIGN KEY (parent_id) REFERENCES post(id) ON DELETE CASCADE
);

-- Create Reaction table
CREATE TABLE reaction (
    id SERIAL PRIMARY KEY,
    reaction_type VARCHAR(10) NOT NULL CHECK (reaction_type IN ('like', 'dislike')),
    parent_id INT,
    account_id INT NOT NULL,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES post(id) ON DELETE CASCADE,
    FOREIGN KEY (account_id) REFERENCES account(id) ON DELETE CASCADE
);

-- Additional indices for optimization
CREATE INDEX idx_post_account_id ON post(account_id);
CREATE INDEX idx_post_parent_id ON post(parent_id);
CREATE INDEX idx_post_timestamp ON post(timestamp);
CREATE INDEX idx_reaction_parent_id ON reaction(parent_id);
CREATE INDEX idx_reation_type ON reaction(reaction_type);
CREATE INDEX idx_hashtag_tag ON hashtag(tag);

-- Create Account table
CREATE TABLE account (
    id SERIAL PRIMARY KEY,
    email VARCHAR(100) NOT NULL UNIQUE,
    username VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    bio TEXT NOT NULL DEFAULT '',
    pfp VARCHAR(255)
);

-- Create Post table
-- Considered a Comment if parent_post_id is not null
-- ratioed column locks the post when it has been ratioed
CREATE TABLE post (
    id SERIAL PRIMARY KEY,
    parent_id INT,
    body TEXT NOT NULL,
    account_id INT NOT NULL,
    ratioed BOOLEAN NOT NULL DEFAULT FALSE,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (parent_id) REFERENCES post(id) ON DELETE CASCADE,
    FOREIGN KEY (account_id) REFERENCES account(id) ON DELETE CASCADE
);

-- Create Attachment table
-- id only exists to differentiate here, no foreign keys link to it
CREATE TABLE attachment (
    id SERIAL PRIMARY KEY,
    parent_id INT NOT NULL,
    src VARCHAR(255) NOT NULL,
    type VARCHAR(10) NOT NULL CHECK (type  IN ('image', 'video')),
    FOREIGN KEY (parent_id) REFERENCES post(id) ON DELETE CASCADE
);

-- Create Hashtag table
-- Composite key because hashtags should be unique on any post
CREATE TABLE hashtag (
    parent_id INT NOT NULL,
    tag VARCHAR(50) NOT NULL,
    PRIMARY KEY (parent_id, tag),
    FOREIGN KEY (parent_id) REFERENCES post(id) ON DELETE CASCADE
);

-- Create Reaction table
-- Composite key because there should only be one reaction per post
CREATE TABLE reaction (
    reaction_type VARCHAR(10) NOT NULL CHECK (reaction_type IN ('like', 'dislike')),
    parent_id INT NOT NULL,
    account_id INT NOT NULL,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (parent_id, account_id),
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
CREATE INDEX idx_attachment_parent_id ON attachment(parent_id)

-- Delete all records from each table before inserting new data
DELETE FROM reaction;
DELETE FROM hashtag;
DELETE FROM post;
DELETE FROM account;

-- Insert test data into account table
INSERT INTO account (id, email, username, password, bio) VALUES
(10, 'alice@example.com', 'alice', 'password10', 'Bio for Alice'),
(11, 'bob@example.com', 'bob', 'password11', 'Bio for Bob'),
(12, 'carol@example.com', 'carol', 'password12', 'Bio for Carol');

-- Insert test data into post table
INSERT INTO post (id, parent_id, body, account_id, ratioed, timestamp) VALUES
(10, NULL, 'This is the first post by Alice', 10, FALSE, CURRENT_TIMESTAMP),
(11, 10, 'This is a comment on the first post by Bob', 11, FALSE, CURRENT_TIMESTAMP),
(12, NULL, 'This is another post by Carol', 12, FALSE, CURRENT_TIMESTAMP),
(13, 12, 'This is a comment on the second post by Alice', 10, FALSE, CURRENT_TIMESTAMP),
(14, NULL, 'This is a third post by Bob', 11, TRUE, CURRENT_TIMESTAMP);

-- Insert test data into hashtag table
INSERT INTO hashtag (parent_id, tag) VALUES
(10, 'firstpost'),
(10, 'introduction'),
(12, 'secondpost'),
(12, 'announcement'),
(14, 'thirdpost');

-- Insert test data into reaction table
INSERT INTO reaction (id, reaction_type, parent_id, account_id, timestamp) VALUES
(10, 'like', 10, 11, CURRENT_TIMESTAMP),
(11, 'dislike', 10, 12, CURRENT_TIMESTAMP),
(12, 'like', 12, 10, CURRENT_TIMESTAMP),
(13, 'dislike', 12, 11, CURRENT_TIMESTAMP),
(14, 'like', 14, 12, CURRENT_TIMESTAMP);

-- Select statements to retrieve data from each table
SELECT * FROM account;
SELECT * FROM post;
SELECT * FROM hashtag;
SELECT * FROM reaction;

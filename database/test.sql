-- CAREFUL: This script will delete all data in the tables and insert new test data
-- Delete all records from each table before inserting new data
DELETE FROM reaction;
DELETE FROM hashtag;
DELETE FROM post;
DELETE FROM account;
DELETE FROM attachment;

-- Insert test data into account table
INSERT INTO account (id, email, username, password, bio) VALUES
(1, 'alice@example.com', 'alice', 'password10', 'Bio for Alice'),
(2, 'bob@example.com', 'bob', 'password11', 'Bio for Bob'),
(3, 'carol@example.com', 'carol', 'password12', 'Bio for Carol');

-- Insert test data into post table
INSERT INTO post (id, parent_id, body, account_id, ratioed, timestamp) VALUES
(1, NULL, 'This is the first post by Alice #firstpost #introduction', 10, FALSE, CURRENT_TIMESTAMP),
(2, 1, 'This is a comment on Alice''s first post by Bob #secontpost', 11, FALSE, CURRENT_TIMESTAMP),
(3, NULL, 'This is another post by Carol', 12, FALSE, CURRENT_TIMESTAMP),
(4, 3, 'This is a comment on the Carol''s post by Alice #letsgo', 10, FALSE, CURRENT_TIMESTAMP),
(5, NULL, 'This is a third post by Bob, with an attachment.', 11, TRUE, CURRENT_TIMESTAMP);

INSERT INTO attachment (parent_id, src) VALUES
(14, 'https://preview.redd.it/the-sacabambaspis-an-extinct-genus-of-jawless-fish-v0-mn6hrm3yk5sa1.jpg?auto=webp&s=fff00db26e0b861d1a4beff64ae3398a11aa68a1', 'image');

-- Insert test data into hashtag table
INSERT INTO hashtag (parent_id, tag) VALUES
(1, 'firstpost'),
(1, 'introduction'),
(2, 'secondpost'),
(4, 'letsgo');

-- Insert test data into reaction table
INSERT INTO reaction (reaction_type, parent_id, account_id, timestamp) VALUES
('like', 10, 11, CURRENT_TIMESTAMP),
('dislike', 10, 12, CURRENT_TIMESTAMP),
('like', 12, 10, CURRENT_TIMESTAMP),
('dislike', 12, 11, CURRENT_TIMESTAMP),
('like', 14, 12, CURRENT_TIMESTAMP);

-- Select statements to retrieve data from each table
-- SELECT * FROM account;
-- SELECT * FROM post;
-- SELECT * FROM hashtag;
-- SELECT * FROM reaction;

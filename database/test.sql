-- CAREFUL: This script will delete all data in the tables and insert new test data
-- Delete all records from each table before inserting new data
DELETE FROM reaction;
DELETE FROM hashtag;
DELETE FROM post;
DELETE FROM account;
DELETE FROM attachment;

-- Insert test data into account table
INSERT INTO account (id, email, username, bio) VALUES
('11111111-1111-1111-1111-111111111111', 'alice@example.com', 'alice', 'Bio for Alice'),
('22222222-2222-2222-2222-222222222222', 'bob@example.com', 'bob', 'Bio for Bob'),
('33333333-3333-3333-3333-333333333333', 'carol@example.com', 'carol', 'Bio for Carol');

-- Insert test data into post table
INSERT INTO post (id, parent_id, body, account_id, ratioed, timestamp) VALUES
('44444444-4444-4444-4444-444444444444', NULL, 'This is the first post by Alice #firstpost #introduction', '11111111-1111-1111-1111-111111111111', FALSE, CURRENT_TIMESTAMP),
('55555555-5555-5555-5555-555555555555', '44444444-4444-4444-4444-444444444444', 'This is a comment on Alice''s first post by Bob #secondpost', '22222222-2222-2222-2222-222222222222', FALSE, CURRENT_TIMESTAMP),
('66666666-6666-6666-6666-666666666666', NULL, 'This is another post by Carol', '33333333-3333-3333-3333-333333333333', FALSE, CURRENT_TIMESTAMP),
('77777777-7777-7777-7777-777777777777', '66666666-6666-6666-6666-666666666666', 'This is a comment on the Carol''s post by Alice #letsgo', '11111111-1111-1111-1111-111111111111', FALSE, CURRENT_TIMESTAMP),
('88888888-8888-8888-8888-888888888888', NULL, 'This is a third post by Bob, with an attachment.', '22222222-2222-2222-2222-222222222222', TRUE, CURRENT_TIMESTAMP);

-- Insert test data into attachment table
INSERT INTO attachment (parent_id, src, type) VALUES
('88888888-8888-8888-8888-888888888888', 'https://preview.redd.it/the-sacabambaspis-an-extinct-genus-of-jawless-fish-v0-mn6hrm3yk5sa1.jpg?auto=webp&s=fff00db26e0b861d1a4beff64ae3398a11aa68a1', 'image');

-- Insert test data into hashtag table
INSERT INTO hashtag (parent_id, tag) VALUES
('44444444-4444-4444-4444-444444444444', 'firstpost'),
('44444444-4444-4444-4444-444444444444', 'introduction'),
('55555555-5555-5555-5555-555555555555', 'secondpost'),
('77777777-7777-7777-7777-777777777777', 'letsgo');

-- Insert test data into reaction table
INSERT INTO reaction (reaction_type, parent_id, account_id, timestamp) VALUES
('like', '44444444-4444-4444-4444-444444444444', '22222222-2222-2222-2222-222222222222', CURRENT_TIMESTAMP),
('dislike', '44444444-4444-4444-4444-444444444444', '33333333-3333-3333-3333-333333333333', CURRENT_TIMESTAMP),
('like', '66666666-6666-6666-6666-666666666666', '11111111-1111-1111-1111-111111111111', CURRENT_TIMESTAMP),
('dislike', '66666666-6666-6666-6666-666666666666', '22222222-2222-2222-2222-222222222222', CURRENT_TIMESTAMP),
('like', '88888888-8888-8888-8888-888888888888', '33333333-3333-3333-3333-333333333333', CURRENT_TIMESTAMP);

-- Select statements to retrieve data from each table
-- SELECT * FROM account;
-- SELECT * FROM post;
-- SELECT * FROM hashtag;
-- SELECT * FROM reaction;

# Ratioed

## Minimum Requirements
- Users can create accounts
- Users with accounts can make posts
- Users with accounts can like or dislike posts
- At a certain ratio of like to dislike, a post's body is automatically removed (it has been 'ratioed')
- A post can only be ratioed if it has more than 100 votes (otherwise one dislike would remove a post immediately)
- Posts can be filtered by hashtags

## Information Modeling
#### Base Facts
- Account has id
- Account has email
- Account has username
- Account has password
- Account has bio
- Post has id
- Post has body
- Post has author
- Post has hashtags
- Post has ratio
- Post has timestamp
- Reaction has an id
- Reaction has reaction type (like or dislike)
- Reaction has parent post or comment
- Reaction has author
- Reaction has timestamp
- Comment inherits from Post
- Comment has id
- Comment has parent post
- Parent post can be a Post or a Comment
- Comment has author
- Comment has body
- Comment has likes
- Comment has dislikes
- Comment has timestamp
#### Constraints
- Account has exactly one id
- Account ids must be unique
- Account has exactly one email
- Emails must be unique
- Account has exactly one username
- Usernames must be unique
- Account has exactly one passsword
- Account has exactly one bio
- Post is linked to exactly one author account
- Post has exactly one id
- Post ids must be unique
- Post has exactly one body
- Post body must have at least one character
- Post has 0 or more hashtags
- Post has ratioed status
- Post has exactly one timestamp
- Reaction has exactly one id
- Reaction ids must be unique
- Reaction is linked to exactly one parent post
- Reaction is linked to exactly one author account
- Reaction has exactly one timestamp
- Reaction has exactly one reaction type
- Reaction type must be 'like' or 'dislike'
- Comment is linked to exactly one author account
- Comment is linked to exactly one parent post
- Comment has exactly one id
- Comment ids must be unique
- Comment has exactly one body
- Comment body must have at least one character
- Comment has 0 or more likes
- Comment has 0 or more dislikes
- Comment has ratioed status
- Comment has exactly one timestamp
#### Derivation Rules
- Ratio = likes / dislikes
#### Schema
- `ratioed/init.sql`


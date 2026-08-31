ALTER TABLE posts
ADD COLUMN user_id INTEGER;

UPDATE posts
SET user_id = 1
WHERE user_id IS NULL;

ALTER TABLE posts
ALTER COLUMN user_id SET NOT NULL;

ALTER TABLE posts
ADD CONSTRAINT fk_posts_user
FOREIGN KEY (user_id)
REFERENCES users(id);
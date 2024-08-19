CREATE TABLE "PostUser" (
    userId VARCHAR(250),
    postId INTEGER,
    PRIMARY KEY (userId, postId),
    FOREIGN KEY (userId) REFERENCES User(userId),
    FOREIGN KEY (postId) REFERENCES Post(postId)
);
CREATE TABLE "CommentsUser" (
    userId VARCHAR(250),
    commentId INTEGER,
    PRIMARY KEY (userId, commentId),
    FOREIGN KEY (userId) REFERENCES User(userId),
    FOREIGN KEY (commentId) REFERENCES Comment(commentId)
);
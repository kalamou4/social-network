CREATE TABLE "CommentsPost" (
    postId INTEGER,
    commentId INTEGER,
    PRIMARY KEY (postId, commentId),
    FOREIGN KEY (postId) REFERENCES Post(postId),
    FOREIGN KEY (commentId) REFERENCES Comment(commentId)
);
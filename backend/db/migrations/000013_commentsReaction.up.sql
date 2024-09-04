CREATE TABLE "CommentsReaction" (
    commentId INTEGER,
    reactionId INTEGER,
    PRIMARY KEY (commentId, reactionId),
    FOREIGN KEY (commentId) REFERENCES Comment(commentId),
    FOREIGN KEY (reactionId) REFERENCES Reaction(reactionId)
);
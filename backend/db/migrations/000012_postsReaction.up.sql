CREATE TABLE "PostsReaction" (
    postId INTEGER,
    reactionId INTEGER,
    PRIMARY KEY (postId, reactionId),
    FOREIGN KEY (postId) REFERENCES Post(postId),
    FOREIGN KEY (reactionId) REFERENCES Reaction(reactionId)
);
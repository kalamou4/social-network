CREATE TABLE "Reactions" (
    reactionId INTEGER PRIMARY KEY,
    reactionType TEXT CHECK( reactionType IN ('like', 'dislike') ),
    createdAt TIMESTAMP
);
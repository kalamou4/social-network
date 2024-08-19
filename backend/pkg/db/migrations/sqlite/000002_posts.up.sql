CREATE TABLE "Posts" (
    postId INT PRIMARY KEY,
    title VARCHAR(255),
    postContent TEXT,
    imageUrl VARCHAR(100),
    categorie VARCHAR(100),
    createdAt TIMESTAMP
);


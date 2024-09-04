CREATE TABLE "Sessions" (
    sessionId VARCHAR(254) PRIMARY KEY,
    userId VARCHAR(254),
    FOREIGN KEY (userId) REFERENCES User(userId)
);
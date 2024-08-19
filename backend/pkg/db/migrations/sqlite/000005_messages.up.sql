CREATE TABLE "Messages" (
    messageId INTEGER PRIMARY KEY,
    senderId INTEGER,
    receiverId INTEGER,
    createdAt TIMESTAMP,
    messageContent TEXT,
    messageType VARCHAR(50),
    messageStatus VARCHAR(50)
);
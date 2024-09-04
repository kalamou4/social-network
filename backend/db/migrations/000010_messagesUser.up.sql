CREATE TABLE "MessagesUser" (
    userId VARCHAR(254),
    messageId INTEGER,
    PRIMARY KEY (userId, messageId),
    FOREIGN KEY (userId) REFERENCES User(userId),
    FOREIGN KEY (messageId) REFERENCES Message(messageId)
);

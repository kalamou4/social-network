CREATE TABLE "MessagesGroup" (
    groupeId INTEGER,
    messageId INTEGER,
    PRIMARY KEY (groupeId, messageId),
    FOREIGN KEY (groupeId) REFERENCES Groupe(groupeId),
    FOREIGN KEY (messageId) REFERENCES Message(messageId)
);
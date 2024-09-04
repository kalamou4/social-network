CREATE TABLE "Members" (
    userId VARCHAR(254),
    groupeId VARCHAR(254),
    PRIMARY KEY (userId, groupeId),
    FOREIGN KEY (userId) REFERENCES User(userId),
    FOREIGN KEY (groupeId) REFERENCES Groupe(groupeId)
);
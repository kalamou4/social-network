CREATE TABLE Notifications(
  notificationId INTEGER,
  receiverId         INTEGER,
  senderId           INTEGER,
  notifType TEXT CHECK( notifType IN ('following request', 'group invitation', 'requests to join the group', 'event')),
  notifContent   TEXT,
  creationTime   TEXT,
  PRIMARY KEY (notificationId AUTOINCREMENT),
  FOREIGN KEY (receiverId) REFERENCES users (userId),
  FOREIGN KEY (senderId) REFERENCES users (userId)
);
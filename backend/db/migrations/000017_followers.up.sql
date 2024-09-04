CREATE TABLE "Followers"(
  followId  INTEGER,
  followerId INTEGER,
  followingId INTEGER,
  followStatus TEXT CHECK( followStatus IN ('follower', 'pending')),
  PRIMARY KEY (followId AUTOINCREMENT),
  FOREIGN KEY (followerId) REFERENCES users (userId)
  FOREIGN KEY (followingId) REFERENCES users (userId)
);
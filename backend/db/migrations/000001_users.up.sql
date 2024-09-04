CREATE TABLE "Users" (
    userId VARCHAR(254) PRIMARY KEY,
    firstname VARCHAR(100),
    lastname VARCHAR(100),
    email VARCHAR5(100),
    password TEXT,
    dateofbirth DATE,
    gender TEXT CHECK( gender IN ('male', 'female') ) DEFAULT 'male',
    nickname VARCHAR(50),
    aboutme TEXT,
    visibility BOOLEAN 
);
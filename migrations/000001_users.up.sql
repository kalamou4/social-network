CREATE TABLE users (
    id STRING PRIMARY KEY,                     
    email VARCHAR(255) NOT NULL UNIQUE,        
    password VARCHAR(255) NOT NULL,            
    first_name VARCHAR(255) NOT NULL,          
    last_name VARCHAR(255) NOT NULL,           
    date_of_birth DATE NOT NULL,               
    avatar_url VARCHAR(255),                   
    nickname VARCHAR(100) NOT NULL,                    
    about_me TEXT,                              
    is_public BOOLEAN DEFAULT TRUE,            
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP  

);

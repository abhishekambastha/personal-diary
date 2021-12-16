CREATE TABLE entries(
    id id INT AUTO_INCREMENT,
    created_timestamp DATETIME,
    text TEXT NOT NULL,
    picture TEXT,
    tid TEXT NOT NULL,
    PRIMARY KEY(id)
);

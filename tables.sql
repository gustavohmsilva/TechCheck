CREATE TABLE techcheck.Genre (
	Id INT auto_increment NOT NULL,
	Name varchar(100) NOT NULL,
	CONSTRAINT Genre_PK PRIMARY KEY (Id),
	CONSTRAINT Genre_UN UNIQUE KEY (Name)
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb4
COLLATE=utf8mb4_general_ci;
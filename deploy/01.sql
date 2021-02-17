CREATE DATABASE IF NOT EXISTS `techcheck`;
GO
GRANT ALL ON `techcheck`.* TO 'techcheck'@'%';
GO
USE techcheck
GO
-- techcheck.Genre definition
CREATE TABLE Genre (
	Id int(10) unsigned NOT NULL AUTO_INCREMENT,
	Name varchar(100) NOT NULL,
	CONSTRAINT Genre_PK PRIMARY KEY (Id),
	CONSTRAINT GenreName_UN UNIQUE KEY (Name)
);
GO
-- techcheck.Book definition
CREATE TABLE Book (
	`Id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	`AuthorId` int(10) unsigned DEFAULT NULL,
	`ISBN` varchar(20) DEFAULT NULL,
	`Name` varchar(200) NOT NULL,
	`Edition` int(11) NOT NULL,
	`IsLatestEdition` tinyint(1) NOT NULL DEFAULT 0,
	`PredecessorEditionId` int(10) unsigned DEFAULT NULL,
	`Image` varchar(256) DEFAULT NULL,
	CONSTRAINT Book_PK PRIMARY KEY (Id),
	CONSTRAINT Book_PreviousEdition_FK FOREIGN KEY (PredecessorEditionId) REFERENCES techcheck.Book(Id)
);
GO
-- techcheck.BookGenre definition
CREATE TABLE BookGenre (
	BookId int(10) unsigned NOT NULL,
	GenreId int(10) unsigned NOT NULL,
	CONSTRAINT BookGenre_Book_FK FOREIGN KEY (BookId) REFERENCES techcheck.Book(Id),
	CONSTRAINT BookGenre_Genre_FK FOREIGN KEY (GenreId) REFERENCES techcheck.Genre(Id)
);
GO
-- techcheck.UserType definition
CREATE TABLE UserType (
	Id int(19) unsigned auto_increment NOT NULL,
	Name varchar(50) NOT NULL,
	CONSTRAINT UserType_PK PRIMARY KEY (Id),
	CONSTRAINT UserTypeName_UN UNIQUE KEY (Name)
);
GO
INSERT INTO UserType (Name) VALUES('Author');
GO
INSERT INTO UserType (Name) VALUES('Revisor');
GO
CREATE TABLE User (
	Id int(10) unsigned auto_increment NOT NULL,
	`User` varchar(100) NOT NULL,
	Pass varchar(100) NOT NULL,
	FullName varchar(100) NOT NULL,
	UserTypeId int(10) unsigned NULL,
	Description TEXT NULL,
	ImageColumn1 varchar(256) NULL,
	CONSTRAINT User_PK PRIMARY KEY (Id),
	CONSTRAINT User_UN UNIQUE KEY (`User`),
	CONSTRAINT User_UserType_FK FOREIGN KEY (UserTypeId) REFERENCES techcheck.UserType(Id)
);
GO
ALTER TABLE Book ADD CONSTRAINT Book_FK_1 FOREIGN KEY (AuthorId) REFERENCES techcheck.`User`(Id);
GO
CREATE TABLE Bank (
	Id int(10) unsigned auto_increment NOT NULL,
	Name varchar(100) NOT NULL,
	Image varchar(256) NULL,
	Since DATETIME DEFAULT current_timestamp() NOT NULL,
	FounderId int(10) unsigned NOT NULL,
	CONSTRAINT Bank_PK PRIMARY KEY (Id),
	CONSTRAINT BankName_UN UNIQUE KEY (Name),
	CONSTRAINT BankFounder_UN UNIQUE KEY (FounderId),
	CONSTRAINT Bank_Founder_FK FOREIGN KEY (FounderId) REFERENCES techcheck.`User`(Id)
);
GO
CREATE TABLE `Error` (
	Id int(10) unsigned auto_increment NOT NULL,
	BookId int(10) unsigned NOT NULL,
	Page INT UNSIGNED NOT NULL,
	ChapterNr INT UNSIGNED NOT NULL,
	ChapterName varchar(256) NOT NULL,
	Paragraph INT UNSIGNED NOT NULL,
	Line INT UNSIGNED NOT NULL,
	Description TEXT NOT NULL,
	ReportedById int(10) unsigned NOT NULL,
	ReportedAt TIME DEFAULT current_timestamp() NULL,
	Accepted BOOL DEFAULT false NOT NULL,
	AcceptedAt TIME NULL,
	Severity INT UNSIGNED DEFAULT 1 NOT NULL,
	CheckIssued BOOL DEFAULT false NOT NULL,
	CONSTRAINT Error_PK PRIMARY KEY (Id),
	CONSTRAINT Error_Book_FK FOREIGN KEY (BookId) REFERENCES techcheck.Book(Id),
	CONSTRAINT Error_Revisor_FK FOREIGN KEY (ReportedById) REFERENCES techcheck.`User`(Id)
);
GO
CREATE TABLE `Check` (
	Id int(10) unsigned auto_increment NOT NULL,
	BankId int(10) unsigned NOT NULL,
	IssueDate TIMESTAMP DEFAULT current_timestamp() NULL,
	BeneficiaryId int(10) unsigned NOT NULL,
	Value FLOAT DEFAULT 2.56 NOT NULL,
	OverErrorId int(10) unsigned NOT NULL,
	CONSTRAINT Check_PK PRIMARY KEY (Id),
	CONSTRAINT Check_Bank_FK FOREIGN KEY (BankId) REFERENCES techcheck.Bank(Id),
	CONSTRAINT Check_beneficiary_FK FOREIGN KEY (BeneficiaryId) REFERENCES techcheck.`User`(Id),
	CONSTRAINT Check_OverError_FK FOREIGN KEY (OverErrorId) REFERENCES techcheck.Error(Id)
);
GO
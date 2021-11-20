# Tutorial and sample codes for accessing relational database

This is to follow https://golang.org/doc/tutorial/database-access.

Go commands to create the mod file
- go mod init example/data-access

MySQL CLI commands to set up the database
- First, open cmd and go to "C:\Program Files\MySQL\MySQL Server 8.0\bin".
- mysql -u root -p
- create database recordings;
- use recordings;
- source \path\to\createtables.sql
- select * from album;

Go commands to get dependencies
- go get .

Go command to run the app
- go run .

MySQL CLI commands to remove the database
- mysql -u root -p
- SHOW DATABASES;
- SHOW SCHEMAS;
- DROP DATABASE recordings;

MySQL CLI commands to manage users
- CREATE USER 'newuser'@'localhost' IDENTIFIED BY 'password';
- GRANT ALL PRIVILEGES ON * . * TO 'newuser'@'localhost';
- FLUSH PRIVILEGES;
- SELECT User, Host  FROM mysql.user;
- DROP USER 'newuser'@'localhost';




# cmpe-132-project
---

## Webstack
---

- [Go](https://go.dev/doc/install)
- [Sqlite3](https://www.sqlite.org/download.html)
- [Javascript](https://www.javascript.com/)
- [JQuery](https://jquery.com/)
- [HTML](https://en.wikipedia.org/wiki/HTML)
- [BootStrap](https://getbootstrap.com/)

## How To Guide
---
1. To setup the tool, please download and setup all of the tools listed above
2. [sql/test.sql](./sql/test.sql) The way this tool works is through having admins add others to the database using the GUI, however the admin would have to be manually added using sql before the tool is deployed. Because of this, please naviagte to the sql file that is linked and create a new user for yourself. Example:
```
INSERT INTO users (roleId, firstName, lastName, studentId, email) VALUES (4, 'Corey', 'Kelley', '014294501', 'corey.kelley@sjsu.edu'); -- 4 indicates admin
```
3. In the sql folder there is a bash script to automatically setup the database, however, it is only tested on windows. So please manually setup the database if you run on a different OS.
4. Run 
```
go mod init;
go build .; ./cmpe-132-project.exe;
```
5. Now navigate to the following link [Website](http://localhost:8080/)

6. First all you need to enter is your studentId then it will take you to a page to enter your temp password

7. When you get to that page, an "email" will be sent to your email, in order to simulate this I just print the temp code to the console which you can copy and paste

## Examples of Working Product
---
### Registering
---
1. 
![alt text](pictures/bare-home-page.png)
2. 
![alt text](pictures/image.png)
3. 
![alt text](pictures/image-1.png) ![alt text](pictures/image-2.png)
4. 
![alt text](pictures/image-3.png)
### Signing In
---
1. 
![alt text](pictures/image-4.png)
### Checking Out Books
---
1. 
![alt text](pictures/image-5.png)
2. 
![alt text](pictures/image-7.png)

### Viewing Checked Out Books to Reserve For Other Students
---
1. 
![alt text](pictures/image-8.png)
2. 
![alt text](pictures/image-9.png)

### Adding New Book to Library
---
1. 
![alt text](pictures/image-10.png)

### Admin Page for Editing Roles
---
1. 
![alt text](pictures/image-11.png)
2. 
![alt text](pictures/image-12.png)
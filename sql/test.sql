INSERT INTO perms (checkoutBook, checkinBook, reserveRoom, reserveBooks, addBooks, createUser, updateUser, deleteUser) VALUES (true, true, true, false, false, false, false, false);
INSERT INTO perms (checkoutBook, checkinBook, reserveRoom, reserveBooks, addBooks, createUser, updateUser, deleteUser) VALUES (true, true, true, true, false, false, false, false);
INSERT INTO perms (checkoutBook, checkinBook, reserveRoom, reserveBooks, addBooks, createUser, updateUser, deleteUser) VALUES (true, true, true, true, true, false, false, false);
INSERT INTO perms (checkoutBook, checkinBook, reserveRoom, reserveBooks, addBooks, createUser, updateUser, deleteUser) VALUES (true, true, true, true, true, true, true, true);

INSERT INTO roles (permsId, roleName) VALUES (1, 'Student');
INSERT INTO roles (permsId, roleName) VALUES (2, 'Professor');
INSERT INTO roles (permsId, roleName) VALUES (3, 'Librarian');
INSERT INTO roles (permsId, roleName) VALUES (4, 'Admin');

INSERT INTO users (roleId, firstName, lastName, studentId, email) VALUES (4, 'Corey', 'Kelley', '014294501', 'corey.kelley@sjsu.edu');
INSERT INTO users (roleId, firstName, lastName, studentId, email) VALUES (4, 'Corey1', 'Kelley1', '014294502', 'corey1.kelley1@sjsu.edu');

INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 10, "1119376513");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376514");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376515");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376516");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376517");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376518");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376519");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376520");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376521");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376522");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376523");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376524");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376525");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376526");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376527");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376528");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376529");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376530");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376531");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376532");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376533");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376534");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376535");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376536");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376537");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376538");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376539");
INSERT INTO books (bookName, count, isbn) VALUES ("Creating a Web Site: Design and Build Your First Site!", 20, "1119376540");

INSERT INTO checkedOutBooks (userId, bookId, checkedOutDate, dueDate) VALUES (0, 0, "1713384597", "1716001797");

INSERT INTO libraryFunds (nameOfFund, funds) VALUES ("Book Fund", 100000);

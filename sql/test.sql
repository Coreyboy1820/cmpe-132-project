INSERT INTO perms (checkoutBook, checkinBook, reserveRoom, reserveBooks, addBooks, createUser, updateUser, deleteUser) VALUES (true, true, true, false, false, false, false, false);
INSERT INTO perms (checkoutBook, checkinBook, reserveRoom, reserveBooks, addBooks, createUser, updateUser, deleteUser) VALUES (true, true, true, true, false, false, false, false);
INSERT INTO perms (checkoutBook, checkinBook, reserveRoom, reserveBooks, addBooks, createUser, updateUser, deleteUser) VALUES (true, true, true, true, true, false, false, false);
INSERT INTO perms (checkoutBook, checkinBook, reserveRoom, reserveBooks, addBooks, createUser, updateUser, deleteUser) VALUES (true, true, true, true, true, true, true, true);

INSERT INTO roles (permsId, roleName) VALUES (0, 'Student');
INSERT INTO roles (permsId, roleName) VALUES (1, 'Professor');
INSERT INTO roles (permsId, roleName) VALUES (2, 'Librarian');
INSERT INTO roles (permsId, roleName) VALUES (3, 'Admin');

INSERT INTO users (roleId, firstName, lastName, studentId, email) VALUES (4, 'Corey', 'Kelley', '014294501', 'corey.kelley@sjsu.edu');
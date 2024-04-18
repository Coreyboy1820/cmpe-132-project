DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS checkedOutBooks;
DROP VIEW IF EXISTS cart;
DROP VIEW IF EXISTS booksInCart;
DROP VIEW IF EXISTS booksAndCheckedOut;

CREATE TABLE IF NOT EXISTS books (
    bookId INTEGER PRIMARY KEY AUTOINCREMENT,
    bookName TEXT,
    count INTEGER,
    isbn TEXT UNIQUE
);

CREATE TABLE IF NOT EXISTS checkedOutBooks (
    checkedOutBooksId INTEGER PRIMARY KEY AUTOINCREMENT,
    userId INTEGER REFERENCES users,
    bookId INTEGER REFERENCES books,
    checkedOutDate TEXT, -- unix timestamps
    dueDate TEXT
);

CREATE TABLE IF NOT EXISTS cart (
    CartId INTEGER PRIMARY KEY AUTOINCREMENT,
    userId INTEGER REFERENCES users,
    bookId INTEGER REFERENCES books
);

CREATE VIEW booksInCart 
AS SELECT 
    userId,
    cart.bookId,
    bookName,
    count,
    isbn
FROM
    books, cart
WHERE
    books.bookId = cart.bookId;

CREATE VIEW booksAndCheckedOut
AS SELECT 
    bookId,
    bookName,
    count,
    isbn,
    checkedOutDate,
    dueDate
FROM 
    checkedOutBooks, books
WHERE
    checkedOutBooks.bookId = books.bookId;

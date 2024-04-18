DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS checkedOutBooks;

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
    checkedOutBooks.bookId = books.bookId 

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS perms;

CREATE TABLE IF NOT EXISTS perms (
    permsId INTEGER PRIMARY KEY AUTOINCREMENT,
    checkoutBook BOOLEAN,
    checkinBook BOOLEAN,
    reserveRoom BOOLEAN,
    reserveBooks BOOLEAN,
    addBooks BOOLEAN,
    createUser BOOLEAN,
    updateUser BOOLEAN,
    deleteUser BOOLEAN
);

CREATE TABLE IF NOT EXISTS roles (
    roleId INTEGER PRIMARY KEY AUTOINCREMENT,
    permsId INTEGER REFERENCES perms,
    roleName TEXT
);

CREATE TABLE IF NOT EXISTS users (
    userId INTEGER PRIMARY KEY AUTOINCREMENT,
    roleId INTEGER REFERENCES roles,
    firstName TEXT,
    lastName TEXT,
    studentId TEXT UNIQUE,
    passwordHash TEXT DEFAULT 0, 
    salt TEXT DEFAULT "",
    passwordSet BOOLEAN DEFAULT false,
    email TEXT,
    loggedIn BOOLEAN DEFAULT false ,
    active BOOLEAN DEFAULT true
);

CREATE VIEW usersAndPerms 
AS SELECT 
    -- user
    userId,
    firstName,
    lastName,
    studentId,
    passwordHash,
    salt,
    email,
    passwordSet,
    loggedIn,
    active,

    -- roles
    roleName,

    -- perms
    checkoutBook,
    checkinBook,
    reserveRoom,
    reserveBooks,
    addBooks,
    createUser,
    updateUser,
    deleteUser
FROM 
    users, roles, perms
WHERE
    users.roleId = roles.roleId 
    AND
    roles.permsId = perms.permsId;

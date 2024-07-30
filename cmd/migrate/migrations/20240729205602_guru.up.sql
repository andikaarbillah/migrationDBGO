CREATE TABLE guru (
    id INTEGER PRIMARY KEY AUTO_INCREMENT NOT NULL,
    name VARCHAR(255) NOT NULL,
    role ENUM('penjas', 'mate-matika', 'bahasa inggris') NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255)
);

INSERT INTO guru ( name, role, email, password ) VALUES(
     "junaidi", 'penjas', 'junaidi@gmail.com', '123456'
);
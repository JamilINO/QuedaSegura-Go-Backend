CREATE DATABASE quedasegura;

CREATE TABLE IF NOT EXISTS Users (
    id uuid PRIMARY KEY,
    email text NOT NULL,
    pass  varchar(150) NOT NULL,
    real_name text NOT NULL
);


CREATE TABLE IF NOT EXISTS Devices (
    id uuid PRIMARY KEY,
    foreign_id uuid REFERENCES Users(id) NOT NULL,
    mac_adress macaddr NOT NULL
);

CREATE TABLE IF NOT EXISTS Contacts (
    id uuid PRIMARY KEY,
    foreign_id uuid REFERENCES Users(id) NOT NULL,
    email text NOT NULL
);


INSERT INTO Users(id, email, pass, real_name)
VALUES (
    gen_random_uuid (),
    'habdig7@gmail.com',
    'hello',
    'Mateus Vieira'
);

INSERT INTO Devices(id, foreign_id, mac_adress)
VALUES (
    gen_random_uuid (),
    'a7ef7d38-0a71-4064-89be-accd6e35e4c4',
    '9e:9d:17:56:60:b9'
);

INSERT INTO Devices(id, foreign_id, mac_adress)
VALUES (
    gen_random_uuid (),
    'a7ef7d38-0a71-4064-89be-accd6e35e4c4',
    'aa:aa:aa:aa:aa:aa'
);

INSERT INTO Contacts (id, foreign_id, email)
VALUES (
    gen_random_uuid(),
    'a7ef7d38-0a71-4064-89be-accd6e35e4c4',
    'mavieira60@yahoo.com.br' 
)

INSERT INTO Contacts (id, foreign_id, email)
VALUES (
    gen_random_uuid(),
    'a7ef7d38-0a71-4064-89be-accd6e35e4c4',
    '10723904@mackenzista.com.br' 
)

SELECT Users.id, mac_adress, Users.email, Contacts.email FROM Devices 
INNER JOIN Users ON Users.id = Devices.foreign_id 
INNER JOIN Contacts ON Users.id = Contacts.foreign_id
WHERE mac_adress='9e:9d:17:56:60:b9';
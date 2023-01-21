create table users
(
    username          varchar(255) not null
        primary key,
    password_hash     varchar(255),
    registration_date timestamp,
    password_salt     varchar(255),
    first_name        varchar(255),
    last_name         varchar(255)
);

INSERT INTO users (username, password_hash, registration_date, password_salt, first_name, last_name) VALUES ('test username', 'test hash', '1980-01-01 00:00:00.000000', 'test salt', 'test first name', 'test last name');
INSERT INTO users (username, password_hash, registration_date, password_salt, first_name, last_name) VALUES ('new username', '860fb51ba4efb65f5ad5ccbc448044fc1f308b466632d496896807f184fd8b29', '2023-01-06 13:14:29.389422', 'NMSOJOHPHRLpaBSnHLSD', '', '');

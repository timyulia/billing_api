CREATE TABLE accounts
(
    id  serial not null unique,
    amount int not null DEFAULT 0
);
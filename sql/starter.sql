CREATE DATABASE IF NOT EXISTS gamestash_billing;

CREATE TABLE IF NOT EXISTS products
(
    ID     int not null primary key,
    Name   varchar(10),
    Code varchar(3)
);

CREATE TABLE IF NOT EXISTS products
(
    ID             int not null primary key,
    Title          varchar(100),
    Description    text,
    Status         enum ('Not Payed', 'Declined', 'Disabled', 'Processing', 'On Hold', 'Complete'),
    QuantityTypeId int
)
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL
);


insert into users values (1,'太郎','taro@mail.com');
insert into users values (2,'次郎','jiro@mail.com');
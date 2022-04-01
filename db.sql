create table if not exists users (
    id int  auto_increment,
    username varchar(10),
    password varchar(10),
    createdAt datetime,
    updatedAt datetime,
    deletedAt datetime,
    primary key (id)
);
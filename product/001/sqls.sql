create database kaifang2 charset=utf8;
create table if not exists kfperson(
    id int primary key auto_increment,
    name varchar(20),
    idcard char(18)
);


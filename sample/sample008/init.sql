create database db01 character set utf8mb4;
create user 'app01'@'%' identified by 'app01';
grant all on db01.* to 'app01'@'%';
flush privileges;
use db01;

create table users(
 id int auto_increment not null primary key, 
 name varchar(255), 
 password varchar(255),
 created_at timestamp not null default current_timestamp, 
 updated_at timestamp not null default current_timestamp on update current_timestamp, 
 index(id)
);
 
insert into users(name,password) values ("user01","password01");
insert into users(name,password) values ("user02","password02");
insert into users(name,password) values ("user03","password03");


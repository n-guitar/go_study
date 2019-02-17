create database db01 character set utf8mb4;
create user 'app01'@'%' identified by 'app01';
grant all on db01.* to 'app01'@'%';
flush privileges;
use db01;

create table users(
 id int auto_increment not null primary key, 
 name varchar(255), 
 password varchar(255),
 hash varchar(255),
 created_at timestamp not null default current_timestamp, 
 updated_at timestamp not null default current_timestamp on update current_timestamp, 
 index(id)
);
 
insert into users(name,password,hash) values ("user01","password01","4de43d6847096454adf50cc4fed29969343af2e6bcd5e809e57df071d6197468");
insert into users(name,password,hash) values ("user02","password02","965c596433b550f418265f5dcbc61e8a4dcfb70c2c4f3f4f57811ad4429c9470");
insert into users(name,password,hash) values ("user03","password03","c8192233749fc1e067d21c51855a2ca5522cbed89808ef93806873e8be393582");


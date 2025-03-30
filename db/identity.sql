create database if not exists zero;
use zero;

create table if not exists user(
    id bigint primary key not null,
    username varchar(128) not null,
    email varchar(128) not null, 
    avatar varchar(128) default '' not null,
    password_hash varchar(128) not null,
    password_salt varchar(128) not null,
    is_deleted tinyint default 0 comment '0:no|1:yes',
    created_at timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null,
    deleted_at timestamp,
    index(username, is_deleted),
    index(email, is_deleted)
)engine='innodb' charset='utf8mb4';


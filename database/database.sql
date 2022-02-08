create table users (
id bigserial not null primary key,
name varchar(60) not null,
email varchar(255) not null,
password varchar(255) not null);


create table comments (
commentid bigserial primary key not null,
owner integer not null,
date timestamp default current_timestamp,
parentid integer,
content varchar(255),
foreign key (owner) references users(id));


create table posts (
postid bigserial primary key not null,
owner integer not null,
title varchar(255) not null,
date timestamp default current_timestamp,
content varchar(255) not null,
comments integer,
foreign key (owner) references users(id),
foreign key (comments) references comments(commentid) on delete cascade);


create table votehandler (
owner integer not null,
postid integer,
commentid integer,
upvote bool,
foreign key (owner) references users(id),
foreign key (postid) references posts(postid) on delete cascade,
foreign key (commentid) references comments(commentid)  on delete cascade);



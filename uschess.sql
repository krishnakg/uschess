drop database uschess;

create database uschess
  default character set utf8
  default collate utf8_general_ci;

use uschess;

create table event (
id varchar(25),
sections int,
state varchar(5),
city varchar(30),
event_date varchar(25),
players  int,
name varchar(100),
primary key (id)
);

create table section(
id varchar(25),
name varchar(50),
event_id varchar(25),
primary key (id),
foreign key (event_id)
  references event(id) 
  on delete cascade
);

create table player (
id int PRIMARY KEY,
name varchar(100),
state varchar(5));


create table game (
id int not null auto_increment,
event_id varchar(25),
section_id varchar(25) ,
round int,
player1 int,
player1_color tinyint,
player2 int,
player2_color tinyint,
result varchar(1),
primary key (id),
foreign key (event_id)
  references event(id) 
  on delete cascade,
foreign key (section_id)
  references section(id) 
  on delete cascade,
foreign key (player1)
  references player(id)
  on delete cascade,
foreign key (player2)
  references player(id)
  on delete cascade
);

create table tournament_history(
id int not null auto_increment,
uscf_id int,
event_id varchar(25),
section_id varchar(25),
rating_type varchar(15),
pre_rating int,
post_rating int,
score float,
primary key (id),
foreign key (event_id)
  references event(id) 
  on delete cascade,
foreign key (section_id)
  references section(id) 
  on delete cascade,
foreign key (uscf_id)
  references player(id)
  on delete cascade
 );

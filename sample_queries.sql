use uschess;

/* Number of wins, loses and draws */
select 
(select count(*) from game where  (player1='15375510' and result='L') or (player2 ='15375510' and result='W')) as No_Of_Loses,
(select count(*) Draw from game where  (player1='15375510' and result='D') or (player2 ='15375510' and result='D')) as No_Of_Draws,
(select count(*) Draw from game where  (player1='15375510' and result='W') or (player2 ='15375510' and result='L')) as No_Of_Wins
from dual;

/* Number of wins as white and black*/
select 
(select count(*) Draw from game where (player1_color='1' and (player1='15375510' and result='W')) or (player2_color='1' and (player2 ='15375510' and result='L'))) as No_Of_Wins_As_White, 
(select count(*) Draw from game where (player1_color='2' and (player1='15375510' and result='W')) or (player2_color='2' and (player2 ='15375510' and result='L'))) as No_Of_Wins_As_Black
from dual;

/* Latest rating of a player */
select event_id, rating_type, post_rating 
from tournament_history 
where uscf_id='15375510' and rating_type='R' order by event_id desc limit 1;

select g.id,g.event_id, g.section_id, g.round, g.player1, p1.name, g.player2, p2.name, g.result 
from player p1, game g, player p2 
where 
( g.player1=15375510 or g.player2=15375510 ) and 
p1.id=g.player1 and 
p2.id=g.player2 ;

/* Given player id and sectionid give the opponents and the rating and results*/
select g.event_id, g.section_id, g.round, g.player2, p2.name, g.player1_color, th.post_rating ,g.result, th.score 
from 
(select event_id, section_id, round,player1, player1_color, player2, player2_color, result from game where player1='15375510' and section_id='201801283052.1'
union
select event_id, section_id, round,player2, player2_color, player1, player1_color, 
case game.result
  when 'L' then 'W'
  when 'W' then 'L'
  When 'D' then 'D'
  end as result from game where player2='15375510' and section_id='201801156542.1'
order by round) g, 
player p1, 
player p2, 
tournament_history th
where 
( g.player1=15375510 or g.player2=15375510 ) and 
p1.id=g.player1 and 
p2.id=g.player2 and 
th.uscf_id=p2.id and 
th.uscf_id=g.player2 and 
th.rating_type='R' and 
th.section_id='201801156542.1';

/* Altering the table event and player to support longer state names */
alter table event
modify column state varchar(30);

alter table player
modify column state varchar(30);

/* Altering player table for adding fide_id column*/
alter table player
add column fide_id int;


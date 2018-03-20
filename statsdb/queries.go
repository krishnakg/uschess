package statsdb

const (
	// Insert queries
	queryInsertEvent             = "insert into event values (?, ?, ?, ?, ?, ?, ?)"
	queryInsertSection           = "insert into section values (?, ?, ?)"
	queryInsertGame              = "insert into game (event_id, section_id, round, player1, player1_color, player2, player2_color, result) values (?, ?, ?, ?, ?, ?, ?, ?)"
	queryInsertTournamentHistory = "insert into tournament_history (uscf_id, event_id, section_id, rating_type, pre_rating, post_rating, score) values (?, ?, ?, ?, ?, ?, ?)"
	queryInsertPlayer            = "insert into player (id, name, state) values (?, ?, ?) on duplicate key update name=?, state=?"
	queryInsertFideID            = "insert into player (id, fide_id) values (?, ?) on duplicate key update fide_id=?"

	// Select queries
	queryGetPlayer            = `select name, state from player where id=?`
	queryGetEvent             = `select name, state, city, players, sections from event where id=?`
	queryGetRecentTournaments = `select distinct e.id, e.name, e.city, e.state, e.players 
																		from event e, tournament_history th 
																		where th.event_id=e.id and th.rating_type='R' order by e.id desc limit ?`
	queryGetSectionInfo        = `select id, name from section where event_id=?`
	queryGetPlayerSearchResult = `select id, name, state from player where name like ? limit 10`
	queryGetEventsForPlayer    = `select e.id, e.name, th.section_id, th.uscf_id, th.pre_rating, th.post_rating, th.rating_type 
																		from event e, tournament_history th 
																		where th.event_id=e.id and th.rating_type='R' and th.uscf_id=? order by e.id desc`
	queryGetSectionResults = `select th.section_id, th.uscf_id, p.name, th.rating_type, th.pre_rating, th.post_rating, th.score 
																		from tournament_history th, player p 
																		where p.id=th.uscf_id and th.rating_type='R' and th.section_id=?`
	queryGetAllGamesInSection = `select g.id, g.section_id, g.round, g.result, g.player1, p1.name, g.player1_color, g.player2, p2.name, g.player2_color 
																		from player p1, game g, player p2 
																		where g.section_id=? and  p1.id=g.player1 and  p2.id=g.player2 and 
																		g.event_id in (select event_id from tournament_history where rating_type='R');`
	queryGetMutualGames = `select g.id, e.name,g.section_id, g.round, g.result, g.player1, p1.name, g.player1_color, g.player2, p2.name, g.player2_color
																		from player p1, game g, player p2, event e 
																		where (( g.player1=? and g.player2=?) or (g.player1=? and g.player2=?)) and 
																		p1.id=g.player1 and  p2.id=g.player2 and e.id=g.event_id;`
	queryPlayersInRatingRangeAndNoFide = `select distinct uscf_id 
																					from tournament_history th, player p 
																					where th.uscf_id=p.id and 
																					p.fide_id is null and 
																					post_rating >= ? and post_rating < ?`

	// Delete queries
	queryDeleteEvent = "delete from event where id like ?"
)

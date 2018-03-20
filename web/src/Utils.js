
export const getAbsolutePathForTournament = function(tournamentId) {
  return '/tournaments/' + tournamentId;
}

export const getAbsolutePathForSection = function(sectionId) {
  return '/tournaments/' + sectionId.replace(".", "/");
}

export const getAbsolutePathForPlayer = function(playerId) {
  return '/players/' + playerId;
}

export const getAbsolutePathForPlayerCompare = function(player1Id, player2Id) {
  return '/players/' + player1Id + '/vs/' + player2Id;
}

  // The tournament Id is of the form YYYYMMDD. So we convert it to the form YYYY-MM-DD.
export const tournamentIdToDateString = function(id) {
    return id.slice(0, 4) + '-' + id.slice(4,6) + '-' + id.slice(6,8);
  }

export const uscfPlayerURL = function(id) {
  return "http://www.uschess.org/assets/msa_joomla/MbrDtlMain.php?" + id
}

export const uscfTournamentURL = function(id) {
  return "http://www.uschess.org/msa/XtblMain.php?" + id;
}

export const chessDbPlayerURL = function(id) {
  return "https://chess-db.com/public/pinfo.jsp?id=" + id;
}

export default {
  getAbsolutePathForTournament : getAbsolutePathForTournament,
  getAbsolutePathForSection : getAbsolutePathForSection,
  getAbsolutePathForPlayer : getAbsolutePathForPlayer,
  getAbsolutePathForPlayerCompare : getAbsolutePathForPlayerCompare,
  tournamentIdToDateString : tournamentIdToDateString,
  uscfPlayerURL : uscfPlayerURL,
  uscfTournamentURL : uscfTournamentURL
}
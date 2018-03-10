
export const getAbsolutePathForTournament = function(tournamentId) {
  return '/tournaments/' + tournamentId;
}

export const getAbsolutePathForSection = function(sectionId) {
  return '/tournaments/' + sectionId.replace(".", "/");
}

export const getAbsolutePathForPlayer = function(playerId) {
  return '/players/' + playerId;
}

  // The tournament Id is of the form YYYYMMDD. So we convert it to the form YYYY-MM-DD.
export const tournamentIdToDateString = function(id) {
    return id.slice(0, 4) + '-' + id.slice(4,6) + '-' + id.slice(6,8);
  }

export default {
  getAbsolutePathForTournament : getAbsolutePathForTournament,
  getAbsolutePathForSection : getAbsolutePathForSection,
  getAbsolutePathForPlayer : getAbsolutePathForPlayer,
  tournamentIdToDateString : tournamentIdToDateString
}
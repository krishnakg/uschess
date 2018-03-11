const apiUrl = 'http://localhost:8080'
const Configs = {
  playerInfoUrl : apiUrl + "/players/",
  playerEventsUrl : apiUrl + "/events?uscf_id=",
  playerSearchUrl : apiUrl + "/playersearch/",
  tournamentInfoUrl : apiUrl + "/tournaments/",
  sectionResultUrl : apiUrl + "/sections/",
  sectionGamesUrl : apiUrl + "/games/"
};

export default Configs;
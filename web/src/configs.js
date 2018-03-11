var apiUrl = 'http://54.245.37.181:8080'
if (process.env.NODE_ENV != 'production') {
  apiUrl = 'http://localhost:8080'
}

const Configs = {
  playerInfoUrl : apiUrl + "/players/",
  playerEventsUrl : apiUrl + "/events?uscf_id=",
  playerSearchUrl : apiUrl + "/playersearch/",
  tournamentInfoUrl : apiUrl + "/tournaments/",
  sectionResultUrl : apiUrl + "/sections/",
  sectionGamesUrl : apiUrl + "/games/"
};

export default Configs;
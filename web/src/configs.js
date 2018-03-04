const apiServerUrl = "http://localhost:8080/";
const Configs = {
  playerInfoUrl : apiServerUrl + "players/",
  playerEventsUrl : apiServerUrl + "events?uscf_id=",
  playerSearchUrl : apiServerUrl + "playersearch/",
  tournamentInfoUrl : apiServerUrl + "tournaments/",
};

export default Configs;
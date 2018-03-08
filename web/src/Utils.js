
export const getAbsolutePathForSection = function(sectionId) {
  return '/tournaments/' + sectionId.replace(".", "/");
}

export const getAbsolutePathForPlayer = function(playerId) {
  return '/players/' + playerId;
}

export default {
  getAbsolutePathForSection : getAbsolutePathForSection,
  getAbsolutePathForPlayer : getAbsolutePathForPlayer
}
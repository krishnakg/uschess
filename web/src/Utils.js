
export const getAbsolutePathForSection = function(sectionId) {
  return '/tournaments/' + sectionId.replace(".", "/")
}

export default {
  getAbsolutePathForSection : getAbsolutePathForSection
}
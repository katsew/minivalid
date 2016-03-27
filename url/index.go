package url

import (
  "regexp"
)

/*
 *  ref. Angluar.js
 *  https://github.com/angular/angular.js/blob/master/src/ng/directive/input.js#L26
 */
var urlRegExp = `^[a-z][a-z\d.+-]*:\/*(?:[^:@]+(?::[^@]+)?@)?(?:[^\s:/?#]+|\[[a-f\d:]+\])(?::\d+)?(?:\/[^?#]*)?(?:\?[^#]*)?(?:#.*)?$`
func Validate(urlStr string) bool {
  r, e := regexp.Compile(urlRegExp)

  if e != nil {
    return false
  }
  res := r.FindAllString(urlStr, -1)
  return len(res) > 0
}

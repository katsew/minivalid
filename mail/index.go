package mail

import (
  "regexp"
)

/*
 *  ref. Angluar.js
 *  https://github.com/angular/angular.js/blob/master/src/ng/directive/input.js#L27
 *
 *  Because Go cannot escape backticks,
 *  combine as a string.
 *  http://stackoverflow.com/questions/21198980/golang-how-to-escape-back-ticks
 *
 */
var mailRegExp = `^[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~.-]+@[a-z0-9]([a-z0-9-]*[a-z0-9])?(.[a-z0-9]([a-z0-9-]*[a-z0-9])?)*$`
func Validate(mailStr string) bool {
  r, e := regexp.Compile(mailRegExp)

  if e != nil {
    return false
  }
  res := r.FindAllString(mailStr, -1)
  return len(res) > 0
}

package main

import "regexp"

/* subst -- substitute pattern with "sub" */
func subst(sub string) stcode {
  var line string
  var stat stcode
  stat = ERR
  for i := line1; i <= line2; i++ {
    line = buf[i].txt
    r, err := regexp.Compile(pat)
    if err != nil {
      stat = ERR
    } else {
      if r.MatchString(line) {
        curln = i
        buf[i].txt = r.ReplaceAllString(line, sub)
      }
    }
  }
  stat = OK
  return stat
}

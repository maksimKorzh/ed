package main

import "regexp"

/* patscan -- find next occurrence of pattern after line n */
func patscan (way rune, n *int) stcode {
  var done bool
  var line string
  var stat stcode
  *n = curln
  stat = ERR
  done = false;
  for {
    if way == SCAN {
      *n = nextln(*n)
    } else {
      *n = prevln(*n)
    }
    line = buf[*n].txt
    r, err := regexp.Compile(pat)
    if err != nil {
      stat = ERR
    } else {
      if r.MatchString(line) {
        stat = OK
        done = true
      }
    }  
    if *n == curln || done { break }
  }
  return stat
}

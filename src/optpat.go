package main

import "strings"

/* optpat -- get optional pattern from lin[i], increment i */
func optpat(lin string, i *int) stcode {
  if *i >= len(lin)-1 { return ERR }
  if lin[*i+1:] == "\n" {
    *i = 0
  } else if lin[*i+1] == lin[*i] {    /* repeated delimiter */
   *i++    /* leave existing pattern alone */
  } else {
    pat = strings.Split(lin[*i:], string(lin[*i]))[1]
    *i += len(pat)+1
    if pat == "" { *i = 0 }
    if *i == 0 { return ERR }
  }
  return OK
}

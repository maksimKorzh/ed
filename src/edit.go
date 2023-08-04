package main

import "fmt"

/* edit -- main routine for text editor */
func main() {
  setbuf()
  var cursave, i int
  var status stcode
  for {
    lin = getline()
    i = 0;
    cursave = curln;
    if getlist(lin, &i, &status) == OK {
      status = docmd(lin, &i, &status)
    }
    if status == ERR {
      fmt.Println("?")
      curln = min(cursave, lastln)
    } else if status == ENDDATA {
      break
    }
  }
}

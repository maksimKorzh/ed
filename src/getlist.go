package main

import "fmt"

/* getlist -- get list of line nums at lin[i], increment i */
func getlist(lin string, i *int, status *stcode) stcode {
  var num int
  var done bool
  line2 = 0;
  nlines = 0;
  done = getone(lin, i, &num, status) != OK
  for *i < len(lin) && done == false {
    line1 = line2
    line2 = num
    nlines++
    if lin[*i] == SEMICOL { curln = num }
    if lin[*i] == COMMA || lin[*i] == SEMICOL {
      *i++
      done = getone(lin, i, &num, status) != OK
      fmt.Println("num2:", num, "; nlines:", nlines)
    } else { done = true }
  }
  nlines = min(nlines, 2);
  fmt.Println("nlines:", nlines)
  if nlines == 0 { line2 = curln }
  if nlines <= 1 { line1 = line2 }
  if (*status != ERR) { *status = OK }

  fmt.Println("getlist line1:", line1, "; line2:", line2, "; curln:", curln)
  return *status
}

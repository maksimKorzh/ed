package main

/* getnum -- get single line number component */
func getnum(lin string, i, num *int, status *stcode) stcode {
  *status = OK
  skipbl(lin, i)
  if *i < len(lin) {
    if lin[*i] >= '0' && lin[*i] <= '9' {
        *num = ctoi(lin, i); //fmt.Println("getnum num:", *num, *i, string(lin[*i]))
        *i--    /* move back; to be advanced at end */
    } else if lin[*i] == CURLINE {
      *num = curln
    } else if lin[*i] == LASTLINE {
      *num = lastln
    } else if lin[*i] == SCAN || lin[*i] == BACKSCAN {
      //if (optpat(lin, i) == ERR) { status = ERR  /* build pattern */
      //} else { status = patscan(lin[*i], num) }
    } else { *status = ENDDATA }
  }
  if (*status == OK) { *i++ }    /* next character to be examined */
  return *status
}

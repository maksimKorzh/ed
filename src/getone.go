package main

/* getone -- get one line number expression */
func getone(lin string, i, num *int, status *stcode) stcode {
  var istart, mul, pnum int
  istart = *i;
  *num = 0;
  if getnum(lin, i, num, status) == OK {    /*  1st term */
    for {
      skipbl(lin, i)
      if lin[*i] != PLUS && lin[*i] != MINUS {
        *status = ENDDATA
      } else {
        if lin[*i] == PLUS { mul = +1 } else { mul = -1 }
        *i++
        if getnum(lin, i, &pnum, status) == OK {
          *num = *num + mul * pnum
        }
        if *status == ENDDATA { *status = ERR }
      }
      if *status != OK { break }
    }
  }
  if *num < 0 || *num > lastln { *status = ERR }
  if *status != ERR {
    if *i <= istart {
      *status = ENDDATA
    } else {
      *status = OK
    }
  }
  return *status
}

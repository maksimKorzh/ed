package main

/* getfn -- get file name from lin[i]... */
func getfn(lin string, i *int, fil *string) stcode {
  var k int
  var stat stcode
  stat = ERR
  if lin[*i+1] == BLANK {
    *i += 2
    k = getword(lin, i, fil)    /* get new filename */
    if (k > 0) { stat = OK }
  } else if lin[*i+1] == NEWLINE && savefile != "" {
    *fil = savefile
    stat = OK
  }
  if stat == OK && savefile == "" {
    savefile = *fil    /* save if no old one */
  }
  return stat
}

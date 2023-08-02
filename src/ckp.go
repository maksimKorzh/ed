package main

/* ckp -- check for "p" after command */
func ckp (lin string, i *int,  pflag *bool, status *stcode) stcode {
  skipbl(lin, i);
  if lin[*i] == PCMD {
    *i++
    *pflag = true
  } else {
    *pflag = false
  }

  if lin[*i] == NEWLINE {
    *status = OK
  } else {
    *status = ERR
  }
  return *status
}

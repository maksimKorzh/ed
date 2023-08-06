package main

/* lndelete -- delete lines n1 through n2 */
func lndelete(n1, n2 int, status *stcode) stcode {
  if n1 <= 0 {
    *status = ERR
  } else {
    newbuf := make([]buftype, len(buf) - (n2 - n1) - 1)
    copy(newbuf[:n1], buf[:n1])
    copy(newbuf[n1:], buf[n2+1:])
    buf = newbuf
    lastln = lastln - (n2 - n1 + 1)
    curln = prevln(n1)
    *status = OK
  }
  return *status
}

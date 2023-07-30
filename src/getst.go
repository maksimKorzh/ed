package main

/* getst -- get status code description */
func getst(status stcode) string {
  descr := ""
  switch status {
    case ENDDATA: descr = "ENDDATA"
    case ERR: descr = "ERR"
    case OK: descr = "OK"
  }
  return descr
}

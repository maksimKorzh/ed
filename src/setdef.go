package main

/* setdef -- set defaulted line numbers */
func setdef(def1, def2 int, status *stcode) stcode {
  if (nlines == 0) {
    line1 = def1;
    line2 = def2
  }
  if line1 > line2 || line1 <= 0 {
    *status = ERR
  } else {
    *status = OK;
  }
  return *status
}

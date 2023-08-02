package main

/* setbuf (in memory) -- initialize line storage buffer */
func setbuf() {
  buf = []buftype{buftype{txt: "", mark: false},}
  curln = 0;
  lastln = 0
}


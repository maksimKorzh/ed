package main

/* puttxt (to buffer) -- put text from lin after curln */
func puttxt (lin string) stcode {
  lastln++
  curln++
  newbuf := make([]buftype, lastln+1)
  copy(newbuf[:curln], buf[:curln])
  newbuf[curln] = buftype{ txt: lin, mark: false }
  copy(newbuf[curln+1:], buf[curln:])
  buf = newbuf
  return OK
}

package main

/* edittype -- types for in-memory version of edit */
type stcode int
type buftype struct {
  txt string   /* text of line */
  mark bool    /* mark of line */
}

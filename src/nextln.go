package main

/* nextln -- get line after n */
func nextln(n int) int {
  if n >= lastln {
    return 0
  } else {
    return n + 1
  }
}

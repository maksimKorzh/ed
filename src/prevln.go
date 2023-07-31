package main

/* prevln -- get line before n */
func prevln(n int) int {
  if n <= 0 {
    return lastln
  } else {
    return n - 1
  }
}

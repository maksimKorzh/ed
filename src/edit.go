package main

import "fmt"

/* edit -- main routine for text editor */
func main() {
  curln = 10
  lastln = 16
  lin = "3+3-$+.+$ , 123"
  i := 0; num := 0; status := OK
  getone(lin, &i, &num, &status)
  fmt.Println("num:", num, "; i:", i)
  fmt.Println("status:", getst(status))
}

package main

import "fmt"

/* edit -- main routine for text editor */
func main() {
  curln = 10
  lastln = 160
  //lin = ".,.+10"
  lin = ".+10;$-.  p"
  i := 0; status := OK
  getlist(lin, &i, &status)
  //num := 0; getone(lin, &i, &num, &status)
  fmt.Println("status:", getst(status))
}

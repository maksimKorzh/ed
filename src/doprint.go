package main

import "fmt"

/* doprint -- print lines n1 through n2 */
func doprint (n1, n2 int, c rune) stcode {
  if (n1 <= 0) {
    return ERR
  } else {
    for i := n1; i <= n2; i++ {
      if c == NCMD { fmt.Print(i, "\t") }
      fmt.Println(buf[i].txt)
    }
    curln = n2;
    return OK
  }
}

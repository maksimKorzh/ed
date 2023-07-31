package main

import "fmt"

/* doprint -- print lines n1 through n2 */
func doprint (n1, n2 int) stcode {
  if (n1 <= 0) {
    return ERR
  } else {
    for i := n1; i <= n2; i++ {
      fmt.Println(buf[i].txt)
    }
    curln = n2;
    return OK
  }
}

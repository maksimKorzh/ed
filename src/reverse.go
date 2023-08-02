package main

/* reverse -- reverse buf[n1]...buf[n2] */
func reverse(n1, n2 int) {
var temp buftype
  for n1 < n2 {
    temp = buf[n1]
    buf[n1] = buf[n2]
    buf[n2] = temp
    n1 = n1 + 1
    n2 = n2 - 1
  }
}

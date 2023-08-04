package main

/* blkcopy -- copy block of lines n1..n2 to after n3 */
func blkcopy(n1, n2, n3 int) {
  cp(n1,n2)
  curln = n3
  for i := n1; i <= n2; i++ {
    puttxt(cpb[i-n1].txt)
  }
}

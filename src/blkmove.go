package main

/* blkmove -- move block of lines n1..n2 to after n3 */
func blkmove(n1, n2, n3 int) {
  if n3 < n1-1 {
    reverse(n3+1, n1-1)
    reverse(n1, n2)
    reverse(n3+1, n2)
  } else if n3 > n2 {
    reverse(n1, n2)
    reverse(n2+1, n3)
    reverse(n1, n3)
  }
}

package main

/* cp -- copy lines n1,n2 to copy buffer */
func cp(n1, n2 int) {
  cpb = []buftype{}
  for i := n1; i <= n2; i++ {
    cpb = append(cpb, buftype{txt: buf[i].txt, mark: buf[i].mark })
  }
}

package main

import "os"
import "fmt"
import "bufio"

/* dowrite -- write lines n1..n2 into file */
func dowrite(n1, n2 int, fil string) stcode {
  file, err := os.Create(fil)
  if err != nil { return ERR }
  defer file.Close()
  writer := bufio.NewWriter(file)
  count := 0
  for line := n1; line <= n2; line++ {
    newline := "\n"
    if line == len(buf) { newline = "" }
    writeln := buf[line].txt + newline
    _, err = writer.WriteString(writeln)
    if err != nil { return ERR }
    count += len(writeln)
  }
  writer.Flush()
  fmt.Println(count)
  return OK
}


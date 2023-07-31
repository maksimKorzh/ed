package main

import "fmt"

/* edit -- main routine for text editor */
func main() {
  buf =[]buftype{
    buftype{ txt: "You should never see this", mark: false },
    buftype{ txt: "line1", mark: false },
    buftype{ txt: "line2", mark: false },
    buftype{ txt: "line3", mark: false },
  }
  //buf = append(buf, buftype{txt: "", mark: false})
  curln = 0
  lastln = len(buf)-1

  var cursave, i int
  var status stcode
  for {
    lin = getline()
    i = 0;
    cursave = curln;
    if getlist(lin, &i, &status) == OK {
      if false /*ckglob(lin, &i, &status) = OK*/ {
        //status = doglob(lin, &i, &cursave, &status)
      } else if status != ERR {
        status = docmd(lin, &i, false, &status)
      }
      /* else ERR, do nothing */
    }
    if status == ERR {
      fmt.Println("?")
      curln = min(cursave, lastln)
    } else if status == ENDDATA {
      break // just for now
    }
  }
}

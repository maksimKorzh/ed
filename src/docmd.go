package main

import "fmt"
import "strings"

/* docmd -- handle all commands except globals */
func docmd (lin string, i *int, status *stcode) stcode {
  var fil, sub string
  var line3 int
  var pflag bool
  pflag = false;    /* may be set by d, m, s */
  *status = ERR;
  if lin[*i] == COMMA {
    line1 = 1
    line2 = lastln
    nlines = 1
    *i++
  }
  if lin[*i] == PCMD || lin[*i] == NCMD {
    if lin[*i+1] == NEWLINE {
      if setdef(curln, curln, status) == OK {
        *status = doprint(line1, line2, rune(lin[*i]))
      }
    }
  } else if lin[*i] == NEWLINE {
    if nlines == 0 { line2 = nextln(curln) }
    *status = doprint(line2, line2, PCMD)
  } else if lin[*i] == QCMD {
    if lin[*i+1] == NEWLINE && nlines == 0 {
      *status = ENDDATA
    }
  } else if lin[*i] == ACMD {
    if lin[*i+1] == NEWLINE {
      *status = lnappend(line2)
    }
  } else if lin[*i] == CCMD {
    if lin[*i+1] == NEWLINE {
      if setdef(curln, curln, status) == OK {
        if lndelete(line1, line2, status) == OK {
          *status = lnappend(prevln(line1))
        }
      }
    }
  } else if lin[*i] == DCMD {
    *i++
    if ckp(lin, i, &pflag, status) == OK {
      if setdef(curln, curln, status) == OK {
        if lndelete(line1, line2, status) == OK {
          if nextln(curln) != 0 {
            curln = nextln(curln)
          }
        }
      }
    }
  } else if lin[*i] == ICMD {
    if lin[*i+1] == NEWLINE {
      if line2 == 0 {
        *status = lnappend(0)
      } else {
        *status = lnappend(prevln(line2))
      }
    }
  } else if lin[*i] == EQCMD {
    *i++
    if ckp(lin, i, &pflag, status) == OK {
      fmt.Println(line2)
    }
  } else if lin[*i] == MCMD {
    *i++
    if getone(lin, i, &line3, status) == ENDDATA { *status = ERR }
    if *status == OK {
      if ckp(lin, i, &pflag, status) == OK {
        if setdef(curln, curln, status) == OK {
          *status = move(&line3)
        }
      }
    }
  } else if lin[*i] == TCMD {
    *i++
    if getone(lin, i, &line3, status) == ENDDATA { *status = ERR }
    if *status == OK {
      if ckp(lin, i, &pflag, status) == OK {
        if setdef(curln, curln, status) == OK {
          *status = dup(&line3)
        }
      }
    }
  } else if lin[*i] == YCMD {
    *i++
    if ckp(lin, i, &pflag, status) == OK {
      if setdef(curln, curln, status) == OK {
        cp(line1, line2)
        *status = OK
      }
    }
  } else if lin[*i] == XCMD {
    *i++
    if ckp(lin, i, &pflag, status) == OK {
      curln = line2
      if len(cpb) > 0 {
        for i := 0; i < len(cpb); i++ {
          puttxt(cpb[i].txt)
        }
        *status = OK
      } else {
        *status = ERR
      }
    } 
  } else if lin[*i] == SCMD {
    *i++
    if optpat(lin, i) == OK {
      sub = strings.Split(lin[*i:], string(lin[*i]))[1]
      *i += len(sub)+2
      if *i < len(lin) {
        if ckp(lin, i, &pflag, status) == OK {
          if setdef(1, lastln, status) == OK {
            *status = subst(sub)
          }
        }
      }
    }
  } else if lin[*i] == ECMD {
    if nlines == 0 {
      if getfn(lin, i, &fil) == OK {
        savefile = fil
        setbuf()
        *status = doread(0, fil)
      }
    }
  } else if lin[*i] == FCMD {
    if nlines == 0 {
      if getfn(lin, i, &fil) == OK {
        savefile = fil
        fmt.Println(savefile)
        *status = OK
      }
    }
  } else if lin[*i] == RCMD {
    if getfn(lin, i, &fil) == OK {
      *status = doread(line2, fil)
    }
  } else if lin[*i] == WCMD {
    if getfn(lin, i, &fil) == OK {
      if setdef(1, lastln, status) == OK {
        *status = dowrite(line1, line2, fil)
      }
    }
  }
  /* else status is ERR */
  if *status == OK && pflag {
    *status = doprint(curln, curln, PCMD)
  }
  return *status
}

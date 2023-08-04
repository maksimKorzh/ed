package main

/* editcons -- const declarations for edit */
const (
  ENDDATA stcode = iota
  ERR
  OK
)

const (
  BLANK    = ' '
  TAB      = '\t'
  COMMA    = ','
  SEMICOL  = ';'
  CURLINE  = '.'
  PERIOD   = '.'
  LASTLINE = '$'
  SCAN     = '/'
  BACKSCAN = '\\'
  PLUS     = '+'
  MINUS    = '-'
  NEWLINE  = '\n'
  PCMD     = 'p'
  NCMD     = 'n'
  QCMD     = 'q'
  ACMD     = 'a'
  DCMD     = 'd'
  CCMD     = 'c'
  ICMD     = 'i'
  EQCMD    = '='
  MCMD     = 'm'
  TCMD     = 't'
  YCMD     = 'y'
  XCMD     = 'x'
  SCMD     = 's'
  ECMD     = 'e'
  FCMD     = 'f'
  RCMD     = 'r'
  WCMD     = 'w'
)

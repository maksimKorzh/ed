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
  QCMD     = 'q'
  ACMD     = 'a'
)

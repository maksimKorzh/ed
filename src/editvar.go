package main

/* editvar -- variables for edit */
var buf[]buftype      /* editor's buffer */

var line1 int         /* first line number */
var line2 int         /* second line number */
var nlines int        /* # of line numbers specified */
var curln int         /* current line -- value of dot */
var lastln int        /* last line -- value of $ */

var pat string        /* pattern */
var lin string        /* input line */
var savefile string;  /* remembered file name */

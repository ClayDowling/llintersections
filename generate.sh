#!/bin/sh

TEXTFILE=graph$1.txt
DOTFILE=graph$1.gv
PNGFILE=graph$1.png

generator/generator > ${TEXTFILE}
./llintersections -graph < ${TEXTFILE} > ${DOTFILE}
dot -Tpng -o${PNGFILE} ${DOTFILE}
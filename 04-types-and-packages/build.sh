#!/bin/sh
gccgo -c somepackage.go && \
gccgo -c somepackage2.go && \
gccgo main.go somepackage.o somepackage2.o && \
./a.out


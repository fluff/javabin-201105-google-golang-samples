#!/bin/sh
8g -I $HOME/devel/go/web.go/_obj main.go && \
8l -L $HOME/devel/go/web.go/_obj main.8 && \
./8.out


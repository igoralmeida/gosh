# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)
include $(GOROOT)/src/Make.cmd

TARG=gosh-bin

gosh-main: gosh
	$(GC) gosh-main.go
	$(LD) -o $(TARG) gosh-main.8 gosh.8

gosh:
	$(GC) gosh.go

all: gosh-main


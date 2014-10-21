MAINDIR	= ./
LIBDIR	= $(MAINDIR)/lib
BINDIR	= $(MAINDIR)
CC	= gcc -O2 -g
LINK.c	= gcc -O2 -lm



LIBOBJ1	= nrutil.o kstwo.o probks.o sort.o KS.o
 
all:	ksm
 
 
ksm:	$(LIBOBJ1)
	$(LINK.c) -o $(BINDIR)/ksm $(LIBOBJ1)
	rm -f $(LIBOBJ1)

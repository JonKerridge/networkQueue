networkQueue
============

A program that exercises a queue process
and the Example1 from pnp
demonstrates that there are problems with the channels in GO

iF two networks use the came channel for output the netwroks interleave on that channel!

Can be overcome by using two different output channels.

The real problem is that it is hard to terminate a set of go routines

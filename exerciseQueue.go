/*
this network of go routine creates a Producer go routine
that puts data values in to a circular Queue as quickly
as possible.

a Prompt go routine requests a value from the Queue and then
the retrieved value is printed.

a Queue go routine accepts inputs on its put channel
iff there is space in the queue data structure

a Queue go routine accepts inputs on its get channel
iff there is a value in the queue, the next value
from the queue data is then sent to the Prompt process

a Queue go routine accespt inputs on either of its input
channels provided the queue is not full and not empty

                      ________     get
                put  |        | <-------
  Producer --------> |  Queue |             Prompt ------->
                     |________|  ------->
                                    out

the names are those used in the definition of Queue

This is an example of a fundamental design pattern for
Concurrent systems called Client - Server

Queue is an example of a Server
	A Server accepts a request from a Client and guarantees
		to respond to that request in finite time

Producer and Prompt are examples of Clients
	A Client makes a request to a Server and guarantees
		to read any response immediately.
		A Server does not have to send a response (Producer)
			but if it does then it must be read by the Client
			immediately (Prompt)

A Server can make a request as a Client to another Server

PROVIDED the network of Clients and Servers do not contain
any cycles then the network is deadlock and livelock free!
*/
package main

import (
	"fmt"
	"github.com/JonKerridge/pnp"
)

func main() {

	fmt.Printf("Concurrent and Parallel Systems in GO\n")
	fmt.Printf("Jon Kerridge Edinburgh Napier Univerity\n")
	fmt.Printf("email : j dot kerridge at napier.ac.uk\n")
	fmt.Printf(" blog : http://jonkerridge-goparallel.blogspot.co.uk/\n")
	fmt.Printf(" code : https://bitbucket.org/jkerridge/gows\n\n\n")

	fmt.Printf("A Network that Manipulates a Circular Queue\n")

	p2q := make(chan int)
	pr2q := make(chan int)
	q2pr := make(chan int)
	out := make(chan int)

	go pnp.Producer(p2q, 100)
	go pnp.Queue(p2q, pr2q, q2pr, 5)
	go pnp.Prompt(pr2q, q2pr, out)
	var i int = 0
	var v int = 0
	for i < 50 {
		v = <-out
		fmt.Printf("%v\n", v)
		i = i + 1
	}
	fmt.Printf("Finished Queue Processing\n")
	oute := make(chan int)
	go pnp.Example1(oute)
	var ei int = 0
	var ev int = 0
	for ei < 10 {
		ev = <-oute
		fmt.Printf("\t%v\n", ev)
		ei = ei + 1
	}
	fmt.Printf("Finished Example Processing\n")
}

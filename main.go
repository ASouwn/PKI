package main

/*
#include <stdio.h>
#include <stdlib.h>

void printHello() {
    printf("Hello from C!\n");
}
*/
import "C"
import "fmt"

func main() {
	fmt.Println("hello PKI")
	C.printHello()
}

package main

import (
	"log"
	"reflect"
	"testing"
)

func TestInterface(t *testing.T) {
	f := func(args interface{}) {
		log.Printf("args type is %s\n", reflect.TypeOf(args))
		log.Printf("&args type is %s\n", reflect.TypeOf(&args))
	}
	arg := "string"
	log.Println("if arg")
	f(arg)
	log.Println("if &arg")
	f(&arg)

	funcb := func(args interface{}) interface{} {
		log.Printf("args tipe is %s\n", reflect.TypeOf(args))
		return args
	}
	log.Printf("string type arges type is %s\n", reflect.TypeOf(funcb("strings")))
}

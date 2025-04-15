package main

import (
	"log"
	"os"
	"reflect"
	"strings"
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

func TestFielMethods(t *testing.T) {
	var (
		path_1 = "./dri_1/fiel.txt"
	)

	testPath := func(path string) bool {
		spl := strings.Split(path, "/")
		filePath := strings.Join(spl[:len(spl)-1], "/")
		log.Printf("call testPath with %s\n", filePath)
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			log.Printf("failed to create directory: %v\n", err)
			return true
		}
		return false
	}
	if !testPath(path_1) {
		log.Printf("path %s is exist\n", path_1)
		file, err := os.Create(path_1)
		if err != nil {
			log.Printf("failed to create file: %v\n", err)
			return
		}
		defer file.Close()
		log.Printf("file %s is created\n", path_1)
		file.WriteString("hello world\n")
	} else {
		log.Printf("path %s is not exist\n", path_1)
	}
}

package main

import (
	"log"

	apiservercel "k8s.io/apiserver/pkg/cel"
)

func main() {
	// TestValidateValidatingAdmissionPolicy
	r := apiservercel.CompilationResult{}
	log.Println(r)
	log.Println("main")
}

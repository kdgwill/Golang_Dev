package main

import (
	"fmt"

	"github.com/kdgwill/golang_dev/JAPC_WIG/pascomp"
)

func main() {
	//Get Arguments from command line
	//TODO: Use flags isntead of bare arguments
	//args := os.Args
	//if len(arg) != 2 {
	//	log.Fatal("Usage: ", filepath.Base(arg[0]), " FileName.pas")
	//}
	//EXT := "pas"
	//fileExtension := string(filepath.Ext(arg[0]))
	//fmt.Println(fileExtension)
	//if strings.EqualFold(fileExtension, EXT) == false {
	//	log.Fatal("Extension must be ", EXT)
	//}
	var scanner pascomp.Scanner //Since their is nill arguments can do C++ style init
	//would use scanner:= new(Scanner) to allocate memory for fields set all members to respective zero and return a pointer
	defer scanner.DeinitScanner() //Kind of a hack to mimic destructors
	scanner.InitScanner()         //no constructor so do this way
	var x = 0
	for {
		tok, tokString := scanner.GetToken(&x)
		if tok == pascomp.Tok_Eof {
			break
		}
		fmt.Printf("%-9s %s\n", tokString, tok)
	}
}

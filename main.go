package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main(){

	fmt.Println("Please insert the name of the file you want to create:")
	buf := bufio.NewReader(os.Stdin)
	name, err := buf.ReadString('\n')
	
	if err != nil{
		log.Fatal(err)
	}

	name = name + ".txt"

	fmt.Println("Please insert the content of the file you want to create:")
	buf = bufio.NewReader(os.Stdin)
	content, err := buf.ReadString('\n')

	document := CreateDocument(name, content)

	err = CreateFile(document)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("File ", document.Name, "successfully created.\nContent: ", document.Content, "\nCreation date: ", document.CreatedAt)
	}
}

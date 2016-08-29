package main

import "fmt"

func main() {
    fmt.Printf("   go install github.com/Olkku/hello     <-- compiles the file\n")
    fmt.Printf("   git push -u origin master             <-- Commits to Github\n\n")

    fmt.Printf("   Currently in use:\n")
    fmt.Printf("   nano Hello.go                         <-- open file to text editor\n")
    fmt.Printf("   go run hello.go                      <-- Compiles and run the program\n")
    fmt.Printf("   git commit -a -m \"Comment here\"     <-- Commits to Github\n")
    fmt.Printf("   git push                              <-- Pushes to Github\n")

}


//open -a TextEdit Hello.go 
//go install github.com/Olkku/hello
//git push -u origin master

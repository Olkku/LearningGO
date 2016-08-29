package main


import (
    "fmt"
)

func main() {
    p := fmt.Println

    p("\n   go install github.com/Olkku/hello      <-- compiles the file")
    p("   git push -u origin master              <-- Commits to Github\n")

    p("   Currently in use:")
    p("    nano Hello.go                         <-- open file to text editor")
    p("    go run hello.go                       <-- Compiles and run the program")
    p("    git commit -a -m \"Comment here\"       <-- Commits to Github")
    p("    git push                              <-- Pushes to Github\n")
}

package main


import (
  "fmt"
//  "os"
//  "os/exec"
)


func main() {
    p := fmt.Println
    answer := "n"

    answer = ask()
    for i := 1; i < 3; i++ {
      if answer == "y" {
         p("\n   go install github.com/Olkku/hello      <-- compiles the file")
         p("   git push -u origin master              <-- Commits to Github\n")

         p("   Currently in use:")
         p("    nano Hello.go                         <-- open file to text editor")
         p("    go run hello.go                       <-- Compiles and run the program")
         p("    git commit -a -m \"Comment here\"       <-- Commits to Github")
         p("    git push                              <-- Pushes to Github\n")
         return
      } else if answer == "n" {
         p("Bye bye then!")
         return
      } else {
         p("Give either \"y\" or \"n\"")
         answer = ask()
      }
    }
    p("You tried too many times")
}

func ask() string {
       answer := "n"
       fmt.Println("\nDo you need Go or Git help? (y/n)")
       fmt.Scanf("%s", &answer)
       //CleanScreen()
       return answer
}

func CleanScreen() {
  //c := exec.Command("clear")
  //c.Stdout = os.Stdout
  //c.Run()
}

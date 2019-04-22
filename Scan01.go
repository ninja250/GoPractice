package main
import (
  "fmt"
  "os"
  "bufio"
)
main(){
  stdin := bufio.NewScanner(os.Stdin)
  stdin.Scan()
  text := stdin.Text()
  fmt.Printf(text\n)
}

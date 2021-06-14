package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "time"
)



func writeToDisk (eventType string, desc string, stamp time.Time) {
  var dateFormat string
  dateFormat = stamp.Format(time.UnixDate)

  file, err := os.OpenFile("database.db", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    panic(err)
  }

  if _, err := file.WriteString(eventType+" :: "+desc+ " || " + dateFormat + "\n"); err != nil {
    file.Close()
    panic(err)
  }

  if err := file.Close(); err != nil {
    panic(err)
  }
}

func readPrintFile () {
  file, err := os.Open("database.db")
	if err != nil {
		fmt.Println(err)
	}

  defer file.Close()

  scanner := bufio.NewScanner(file)
	for scanner.Scan() {
    fmt.Println(scanner.Text())
  }

  if err := scanner.Err(); err != nil {
    fmt.Println(err)
  }


}

func main () {
  fmt.Println("Welcome to Event Logger v 0.0.1 - BETA release")
  fmt.Println("Enter your event below and press return/enter to save it into the list")
  fmt.Println("or enter 'list' to see what you already have")
  fmt.Print(":")
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    input := scanner.Text();

    if err := scanner.Err(); err != nil {
		    fmt.Println(os.Stderr, "reading standard input:", err)
	  }
    /* Handle command line input and split if eventType::description*/
    splice := strings.Split(input, "::")
    //fmt.Println("Entered:", input, splice, len(splice) == 1)
    /* Handle commands (no ::) */
    if len(splice) == 1 {
      switch splice[0] {
      case "list":
        fmt.Println("list command", splice[0])
        readPrintFile()
        fmt.Print(":")
      default:
        fmt.Println("Weirdly", splice[0], splice[0] == "list")
        fmt.Print(":")
      }
    } else if len(splice) > 1 {
      var command = splice[0]
      var description = splice[1]
      var timestamp = time.Now()
      //fmt.Println("Data:", command, description, timestamp)
      writeToDisk(command, description, timestamp)
      fmt.Print(":")
    } else {
      fmt.Println("That didn't work")
      fmt.Print(":")
    }
  }
}

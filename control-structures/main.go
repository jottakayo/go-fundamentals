package main

import "fmt"

func main() {
  fmt.Printf("Welcome to Golang Bank!\n")
  fmt.Printf("What do you want to do?\n")
  fmt.Printf("1. Check Balance\n")
  fmt.Printf("2. Deposit Money\n")
  fmt.Printf("3. Withdraw Money\n")
  fmt.Printf("4. Exit\n")

  var choice int

  fmt.Scanln(&choice)

  checkBalance := choice ==1

  if checkBalance {
    fmt.Printf("Your balance is $1000\n")
  }

  

  fmt.Printf("You chose option %d\n", choice)
}


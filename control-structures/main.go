package main

import "fmt"

func main() {
  fmt.Printf("Welcome to Golang Bank!\n")
  fmt.Printf("What do you want to do?\n")
  fmt.Printf("1. Check Balance\n")
  fmt.Printf("2. Deposit Money\n")
  fmt.Printf("3. Withdraw Money\n")
  fmt.Printf("4. Exit\n")
  
  var yourBalance float32 = 1000
  var choice int
  fmt.Printf("which one do you choose?\n")
  fmt.Scanln(&choice)
  // fmt.Printf("You chose option %d\n", choice)
  // checkBalance := choice ==1

  if choice == 1 {
    fmt.Printf("Your balance is %f\n", yourBalance)
  } else if choice == 2 {
    fmt.Printf("Your deposit: ")
    var depositMount float32
    fmt.Scan(&depositMount)
    // Nested if Statements
    if depositMount <= 0 {
      fmt.Printf("Invalid mount. Most be greather than 0.\n")
      return
    }
    yourBalance += depositMount
    fmt.Printf("Balancer update is: %.2f\n", yourBalance)
  } else if choice == 3 {
    fmt.Printf("How much money do you takeout?\n")
    var withdrawMoney float32
    fmt.Scan(&withdrawMoney)
    // Nested if Statements
    if withdrawMoney <= 0 {
      fmt.Printf("Invalid mount. Most be greather than 0.\n")
      return
    }
    if withdrawMoney > yourBalance {
      fmt.Printf("Invalid mount. Most be less than value in your account.\n")
      return
    }

    yourBalance -= withdrawMoney
    fmt.Printf("Balancer update is: %.2f\n", yourBalance)
  } else if choice == 4 {
    fmt.Printf("Ok, thank for coming here!\n")
  } else{
    fmt.Printf("This values is different of options that you may choese!\n")
  }
}


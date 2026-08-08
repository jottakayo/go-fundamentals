package main

import "fmt"

func main() {  
  var yourBalance float32 = 1000
  var choice int
  fmt.Printf("Welcome to Golang Bank!\n")
  for {
    fmt.Printf("What do you want to do?\n")
    fmt.Printf("1. Check Balance\n")
    fmt.Printf("2. Deposit Money\n")
    fmt.Printf("3. Withdraw Money\n")
    fmt.Printf("4. Exit\n")
    fmt.Printf("which one do you choose?\n")
    fmt.Scanln(&choice)
    
    switch choice {
    case 1:
      fmt.Printf("\nYour balance is %.2f\n\n", yourBalance)
    case 2:
      fmt.Printf("Your deposit: ")
      var depositMount float32
      fmt.Scan(&depositMount)
      if depositMount <= 0 {
        fmt.Printf("Invalid mount. Most be greather than 0.\n")
        return
      }
      yourBalance += depositMount
      fmt.Printf("\nBalancer update is: %.2f\n\n", yourBalance)
    case 3:
      fmt.Printf("How much money do you takeout?\n")
      var withdrawMoney float32
      fmt.Scan(&withdrawMoney)
      if withdrawMoney <= 0 {
        fmt.Printf("Invalid mount. Most be greather than 0. try again.\n")
        continue
      }
      if withdrawMoney > yourBalance {
        fmt.Printf("Invalid mount. Most be less than value in your account. Check your balance.\n")
        continue
      }
      yourBalance -= withdrawMoney
      fmt.Printf("\nBalancer update is: %.2f\n\n", yourBalance)
    case 4:
      fmt.Printf("Good bye!\n")
      fmt.Printf("Thank for coming here!\n")
      return
    default:
      fmt.Printf("\nThis values is different of options that you may choese!\n\n")
      continue
    }
  }
}


package main

import "fmt"
import "os"
import "strconv"

const accountBalanceFile = "balancer.txt"

func balanceFromFile() float64{
  data, _ := os.ReadFile(accountBalanceFile)
  balanceText := string(data)
  balance, _ := strconv.ParseFloat(balanceText, 64)

  return balance
}

func balanceToFile(balance float64) {
  balancerTxt := fmt.Sprintf("%v", balance)
  os.WriteFile(accountBalanceFile, []byte(balancerTxt), 0644 )
}

func main() {  
  var yourBalance float64 = balanceFromFile()
  
  fmt.Printf("Welcome to Golang Bank!\n")
  for {
    fmt.Printf("What do you want to do?\n")
    fmt.Printf("1. Check Balance\n")
    fmt.Printf("2. Deposit Money\n")
    fmt.Printf("3. Withdraw Money\n")
    fmt.Printf("4. Exit\n")

    var choice int
    fmt.Printf("which one do you choose?\n")
    fmt.Scanln(&choice)

    if choice == 1 {
      fmt.Printf("\nYour balance is %.2f\n\n", yourBalance)
    } else if choice == 2 {
      fmt.Printf("Your deposit: ")
      var depositMount float64
      fmt.Scan(&depositMount)
      if depositMount <= 0 {
        fmt.Printf("Invalid mount. Most be greather than 0.\n")
        return
      }
      yourBalance += depositMount
      fmt.Printf("\nBalancer update is: %.2f\n\n", yourBalance)
      balanceToFile(yourBalance)
    } else if choice == 3 {
      fmt.Printf("How much money do you takeout?\n")
      var withdrawMoney float64
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
      balanceToFile(yourBalance)
    } else if choice == 4 {
      fmt.Printf("Good bye!")
      break
    } else{
      fmt.Printf("\nThis values is different of options that you may choese!\n\n")
    }
  }
  fmt.Printf("Thank for coming here!\n")
}


package bank

var deposits = make(chan int)
var balances = make(chan int)
var withdraw = make(chan int)
var chanSuccess = make(chan bool)

func Deposit(amount int) {
	deposits <- amount
}
func Balance() int {
	return <-balances
}
func Withdraw(amount int) bool {
	withdraw <- amount
	return <-chanSuccess
}

func teller() {
	var balance int
	for {
		select {
		case amountDeposit := <-deposits:
			balance += amountDeposit
		case balances <- balance:
		case amountWithdraw := <-withdraw:
			if amountWithdraw > balance {
				chanSuccess <- false
			} else {
				balance -= amountWithdraw
				chanSuccess <- true
			}
		}
	}
}

func init() {
	go teller()
}

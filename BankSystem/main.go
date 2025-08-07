package main

type SavingsAccount struct {
	balance int
}

type CheckingAccount struct {
	balance int
}

type InvestmentAccount struct {
	balance int
}

func (s *SavingsAccount) MonthlyInterest() int {
	return s.balance * 5 / 12 / 100
}

func (s *SavingsAccount) CheckBalance() int {
	return s.balance
}

func (s *SavingsAccount) Deposit(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}
	s.balance += amount
	return "Success"
}

func (s *SavingsAccount) Withdraw(amount int) string {
	if amount > 0 {
		if s.balance >= amount {
			s.balance -= amount
			return "Success"
		}
		return "Account balance is not enough"
	} else {
		return "Amount cannot be negative"
	}
}

func (s *SavingsAccount) Transfer(receiver Account, amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}

	switch receiver.(type) {
	case *SavingsAccount, *CheckingAccount, *InvestmentAccount: // Accept ALL types
		if s.balance >= amount {
			s.balance -= amount
			receiver.Deposit(amount)
			return "Success"
		}
		return "Account balance is not enough"
	default:
		return "Invalid receiver account"
	}
}

func (c *CheckingAccount) MonthlyInterest() int {
	return c.balance * 1 / 12 / 100
}

func (c *CheckingAccount) CheckBalance() int {
	return c.balance
}

func (c *CheckingAccount) Deposit(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}
	c.balance += amount
	return "Success"
}

func (c *CheckingAccount) Withdraw(amount int) string {
	if amount > 0 {
		if c.balance >= amount {
			c.balance -= amount
			return "Success"
		}
		return "Account balance is not enough"
	} else {
		return "Amount cannot be negative"
	}
}

func (c *CheckingAccount) Transfer(receiver Account, amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}

	switch receiver.(type) {
	case *SavingsAccount, *CheckingAccount, *InvestmentAccount: // Accept ALL types
		if c.balance >= amount {
			c.balance -= amount
			receiver.Deposit(amount)
			return "Success"
		}
		return "Account balance is not enough"
	default:
		return "Invalid account type"
	}
}

func (i *InvestmentAccount) MonthlyInterest() int {
	return i.balance * 2 / 12 / 100
}

func (i *InvestmentAccount) CheckBalance() int {
	return i.balance
}

func (i *InvestmentAccount) Deposit(amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}
	i.balance += amount
	return "Success"
}

func (i *InvestmentAccount) Withdraw(amount int) string {
	if amount > 0 {
		if i.balance >= amount {
			i.balance -= amount
			return "Success"
		}
		return "Account balance is not enough"
	} else {
		return "Amount cannot be negative"
	}
}

func (i *InvestmentAccount) Transfer(receiver Account, amount int) string {
	if amount <= 0 {
		return "Amount cannot be negative"
	}

	switch receiver.(type) {
	case *SavingsAccount, *CheckingAccount, *InvestmentAccount: // Accept ALL types
		if i.balance >= amount {
			i.balance -= amount
			receiver.Deposit(amount)
			return "Success"
		}
		return "Account balance is not enough"
	default:
		return "Invalid account type"
	}
}

func NewSavingsAccount() *SavingsAccount {
	return &SavingsAccount{balance: 0}
}

func NewCheckingAccount() *CheckingAccount {
	return &CheckingAccount{balance: 0}
}

func NewInvestmentAccount() *InvestmentAccount {
	return &InvestmentAccount{balance: 0}
}

func CheckBalance(account Account) int {
	return account.CheckBalance()
}

func Transfer(sender Account, receiver Account, amount int) string {
	return sender.Transfer(receiver, amount)
}

func Deposit(account Account, amount int) string {
	return account.Deposit(amount)
}

func Withdraw(account Account, amount int) string {
	return account.Withdraw(amount)
}

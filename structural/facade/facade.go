package facade

// it provides a simplified interface to the client, hiding all the complexities of the underlying system
// it provides a higher level of abstraction to a complicated system
// use case: simplified interface to a complex library or framework

// example -> digital wallet
// walletFacade is the facade exposed to the client
// it hides the complexity of the following subsystem: account, security, balance, notifications
type walletFacade struct {
	account      *account
	wallet       *wallet
	security     *security
	notification *notification
}

func newFacade(accountID string, code int) *walletFacade {
	return &walletFacade{
		account:      newAccount(accountID),
		wallet:       newWallet(),
		security:     newCode(code),
		notification: newNotifier(),
	}
}

func (w *walletFacade) addMoney(accountID string, code, amount int) error {
	if err := w.account.checkAccount(accountID); err != nil {
		return err
	}
	if err := w.security.checkCode(code); err != nil {
		return err
	}
	w.wallet.addBalance(amount)
	w.notification.sendAddNotification()
	return nil
}

func (w *walletFacade) deductMoney(accountID string, code, amount int) error {
	if err := w.account.checkAccount(accountID); err != nil {
		return err
	}
	if err := w.security.checkCode(code); err != nil {
		return err
	}
	if err := w.wallet.debitBalance(amount); err != nil {
		return err
	}
	w.notification.sendDebitNotification()
	return nil
}

// COMPARISON
// facade vs adapter
// facade -> defines a new interface for existing objects / works with a group of objects
// adapter -> tries to make the existing interface usable / wraps one object

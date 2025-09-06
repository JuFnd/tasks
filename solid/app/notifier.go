package app

import (
	"fmt"

	"github.com/samber/lo"
)

type Notifier struct {
	somethingSenderClient struct {
		ip   string
		port string
	}
}

func NewNotifier(ip string, port string) *Notifier {
	return &Notifier{
		somethingSenderClient: struct {
			ip   string
			port string
		}{ip: ip, port: port},
	}
}

type Person struct {
	Name    string
	Surname string
	About   string
}

type DebtInfo struct {
	Customer Person
	Debt     int64
}

type EmailMsg = Message[*DebtInfo]

func (n *Notifier) SendMessage(emailFrom string, emailTo string, msg EmailMsg) error {
	log := lo.ToPtr("")
	defer func() {
		fmt.Println("----- EMAIL SENDING LOG -----")
		fmt.Println(*log)
	}()

	generatedMsg, err := msg.RawBody()
	if err != nil {
		*log = fmt.Sprintf("Failed to generate message body: %v", err)
		return fmt.Errorf("error sending message from %s to %s: %w", emailFrom, emailTo, err)
	}

	*log = fmt.Sprintf("Message sent from %s to %s: %s", emailFrom, emailTo, string(generatedMsg))
	return nil
}

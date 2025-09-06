package app

import (
	"fmt"
	"math/rand"
	"sync"

	"github.com/google/uuid"
)

type App struct {
	n  *Notifier
	wg sync.WaitGroup
}

func New() *App {
	return &App{
		n: NewNotifier("localhost", "1111"),
	}
}

func (app *App) Run() {
	anyConnections := rand.Intn(100)
	app.wg.Add(anyConnections)

	for i := 0; i < anyConnections; i++ {
		go func() {
			defer app.wg.Done()

			debt := DebtInfo{
				Customer: Person{
					Name:    uuid.NewString(),
					Surname: uuid.NewString(),
					About:   "Some info",
				},
				Debt: rand.Int63n(100000),
			}

			msg := NewMessage(&debt)
			emailFrom := fmt.Sprintf("%s@example.com", debt.Customer.Surname)
			emailTo := fmt.Sprintf("%s@example.com", debt.Customer.Name)

			err := app.n.SendMessage(emailFrom, emailTo, *msg)
			if err != nil {
				fmt.Printf("Error sending message: %v\n", err)
			}
		}()

	}

	app.wg.Wait()
}

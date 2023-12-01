package examples

import (
	"fmt"
	"github.com/Sortren/go-deps"
	"log"
)

type PaymentMethodOption string

const (
	WithCash       PaymentMethodOption = "cash"
	WithCreditCard PaymentMethodOption = "card"
)

type PaymentMethod interface {
	Process()
}

type Cash struct{}

func (c Cash) Process() {
	fmt.Println("paying with cash")
}

type CreditCard struct{}

func (c CreditCard) Process() {
	fmt.Println("paying with credit card")
}

type PaymentService struct {
	paymentMethodResolver deps.Resolver[PaymentMethodOption, PaymentMethod]
}

func NewPaymentService(paymentMethodResolver deps.Resolver[PaymentMethodOption, PaymentMethod]) *PaymentService {
	return &PaymentService{paymentMethodResolver: paymentMethodResolver}
}

func (p PaymentService) Pay(method PaymentMethodOption) error {
	resolve, err := p.paymentMethodResolver.Resolve(method)
	if err != nil {
		return fmt.Errorf("can't resolve payment method: %w", err)
	}

	resolve.Process()

	return nil
}

func main() {
	paymentResolver := deps.NewGenericResolver(
		map[PaymentMethodOption]PaymentMethod{
			WithCash:       Cash{},
			WithCreditCard: CreditCard{},
		})

	service := NewPaymentService(paymentResolver)

	if err := service.Pay(WithCash); err != nil {
		log.Fatal(err)
	}
}

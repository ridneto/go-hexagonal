package application

import "errors"

type ProductInterface interface {
	isValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED = "enabled"
)

type Product struct {
	ID string
	Name string
	Price float64
	Status string
}

func(p *Product) IsValid() (bool, error) {
	return true, nil
}

func(p *Product) Enabled() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New("price must be greater than 0 to enable the product")
}

func(p *Product) Disable() error {
	return nil
}

func(p *Product) GetID() string {
	return p.ID
}

func(p *Product) GetName() string {
	return p.Name
}

func(p *Product) GetStatus() string {
	return p.Status
}

func(p *Product) GetPrice() float64 {
	return p.Price
}
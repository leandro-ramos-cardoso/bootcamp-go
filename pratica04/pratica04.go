package main

import "fmt" 

type Product struct {
	Name  string
	Price float64
}

type Cart struct {
	Items []Product
}

type Wallet struct {
	Balance float64
}

// Método para adicionar produto — precisa ser ponteiro
func (cart *Cart) AddProduct(product Product) {
	cart.Items = append(cart.Items, product)
}

// Método para calcular o total
func (cart Cart) Total() float64 {
	var total float64 = 0.0

	for _, item := range cart.Items {
		total += item.Price
	}

	return total
}

func (wallet *Wallet) Pay(amount float64) {
	if amount > wallet.Balance {
		fmt.Println("Faça um emprestimo!!!")
	} else {
		wallet.Balance -= amount
		fmt.Println("Pagamento realizado com sucesso!")
	}
}

func main() {
	
	// Criando 3 produtos
	product1 := Product{Name: "Teclado Mecanico", Price: 250.00}
	product2 := Product{Name: "Mouse", Price: 150.00}
	product3 := Product{Name: "Mousepad", Price: 90.00}

	var cart Cart

	cart.AddProduct(product1)
	cart.AddProduct(product2)
	cart.AddProduct(product3)

	wallet := &Wallet{
		Balance: 500.00,
	}

	total := cart.Total()
	fmt.Printf("Total da compra: R$ %.2f\n", total)

	wallet.Pay(total)

	fmt.Printf("Saldo final da carteira: R$ %.2f\n", wallet.Balance)
}
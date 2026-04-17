package main

import (
	"bufio"
	"fmt"
	"homework/internal/product"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	var p product.Product

	fmt.Println("--- Введите данные о товаре ---")

	fmt.Print("Product info: ")
	scanner.Scan()
	p.Name = scanner.Text()

	fmt.Print("Brand: ")
	scanner.Scan()
	p.Brand = scanner.Text()

	fmt.Print("Price: ")
	scanner.Scan()
	inputPrice := scanner.Text()
	cleanPrice := strings.ReplaceAll(inputPrice, " ", "")
	cleanPrice = strings.ReplaceAll(cleanPrice, "\u00a0", "")
	p.Price, _ = strconv.Atoi(cleanPrice)

	fmt.Print("InStock(true/false): ")
	scanner.Scan()
	inputStock := scanner.Text()
	p.InStock, _ = strconv.ParseBool(strings.TrimSpace(inputStock))

	monthly := product.Calculate(p.Price)
	card := product.Converter(p, monthly)

	fmt.Println(card)
}

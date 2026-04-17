package product

import (
	"strconv"
	"strings"
)

type Product struct {
	Name    string
	Price   int
	Brand   string
	InStock bool
}

const (
	template = `===== Alifshop =====
Товар:	{name}
Бренд:	{brand}
Цена:	{price}  сум
В наличии: {inStock}
Рассрочка: 12 мес → {firstAmount} сум/мес
====================`
)

func formatNumber(n int) string {
	s := strconv.Itoa(n)
	var result []string

	for i := len(s); i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		result = append([]string{s[start:i]}, result...)
	}
	return strings.Join(result, " ")
}

func Converter(product Product, firstAmount int) string {
	text := strings.ReplaceAll(template, "{name}", product.Name)
	text = strings.ReplaceAll(text, "{brand}", product.Brand)
	text = strings.ReplaceAll(text, "{price}", formatNumber(product.Price))
	text = strings.ReplaceAll(text, "{inStock}", strconv.FormatBool(product.InStock))
	text = strings.ReplaceAll(text, "{firstAmount}", strconv.Itoa(firstAmount))
	return text
}

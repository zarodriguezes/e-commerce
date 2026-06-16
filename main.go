package main

import (
 "fmt"
)

type Product struct {
 ID int
 Name string
 Price float64
}

type CartItem struct {
 Product Product
 Quantity int
}

func Filter[T any](collection []T, predicate func(T) bool) []T {
 var result []T
 for _, item := range collection {
  if predicate(item) {
   result = append(result, item)
  }
 }
 return result
}

func Map[T any, U any](collection []T, transform func(T) U) []U {
 result := make([]U, len(collection))
 for i, item := range collection {
  result[i] = transform(item)
 }
 return result
}

func Reduce[T any, U any](collection []T, initial U, accumulator func(U, T) U) U {
 result := initial
 for _, item := range collection {
  result = accumulator(result, item)
 }
 return result
}

func FilterByPriceRange(min, max float64) func(Product) bool {
 return func(p Product) bool {
  return p.Price >= min && p.Price <= max
 }
}

func ApplyDiscount(percentage float64) func(Product) Product {
 factor := 1.0 - (percentage / 100.0)
 return func(p Product) Product {
  return Product{
   ID: p.ID,
   Name: p.Name + " (Descuento)",
   Price: p.Price * factor,
  }
 }
}

func main() {
 catalog := []Product{
  {ID: 1, Name: "Teclado Mecánico", Price: 85.50},
  {ID: 2, Name: "Ratón Gamer", Price: 45.00},
  {ID: 3, Name: "Monitor 4K PRO", Price: 350.00},
  {ID: 4, Name: "Auriculares Inalámbricos", Price: 120.00},
 }

 fmt.Println("=== 1. CATÁLOGO INICIAL ===")
 for _, p := range catalog {
  fmt.Printf("- %s: $%.2f\n", p.Name, p.Price)
 }

 premiumProducts := Filter(catalog, func(p Product) bool {
  return p.Price > 50.0
 })

 fmt.Println("\n=== 2. FILTRADO (PRODUCTOS > $50) ===")
 for _, p := range premiumProducts {
  fmt.Printf("- %s: $%.2f\n", p.Name, p.Price)
 }

 discountedProducts := Map(premiumProducts, ApplyDiscount(10.0))

 fmt.Println("\n=== 3. MAP (DESCUENTO DEL 10% APLICADO) ===")
 for _, p := range discountedProducts {
  fmt.Printf("- %s: $%.2f\n", p.Name, p.Price)
 }

 cart := []CartItem{
  {Product: discountedProducts[0], Quantity: 1},
  {Product: discountedProducts[1], Quantity: 2},
 }

 totalCart := Reduce(cart, 0.0, func(accumulator float64, item CartItem) float64 {
  return accumulator + (item.Product.Price * float64(item.Quantity))
 })

 fmt.Println("\n=== 4. PROCESAMIENTO DE ORDEN (REDUCE) ===")
 fmt.Printf("Subtotal de la orden calculada de forma pura: $%.2f\n", totalCart)

 fmt.Println("\n=== 5. COMPROBACIÓN DE INMUTABILIDAD ===")
 fmt.Printf("¿El precio del Teclado original cambió?: $%.2f (Sigue intacto)\n", catalog[0].Price)
}
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/matalmeida/hashlab-service-1-golang/product"
	grpc "google.golang.org/grpc"
)

func main() {
	srv := grpc.NewServer()
	var products productServer
	product.RegisterProductsServer(srv, products)
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("could not listen to :8888: %v", err)
	}

	log.Fatal(srv.Serve(l))
}

func getProductByID(productID string) (*product.Product, error) {
	var products []*product.Product

	products = append(products, &product.Product{Id: "1", PriceInCents: 10000, Title: "Produto de 100 pila", Description: "Bem caro"})
	products = append(products, &product.Product{Id: "2", PriceInCents: 100000, Title: "Produto de 1000 pila", Description: "Caro pra krai"})

	for _, prod := range products {
		if prod.Id == productID {
			return prod, nil
		}
	}

	return nil, fmt.Errorf("the product with id %s doesnt exists", productID)
}

func getUserByID(userID string) (*product.User, error) {
	var users []*product.User

	users = append(users, &product.User{Id: "1", FirstName: "Matheus", LastName: "Anjos", DateOfBirth: 819590400})
	users = append(users, &product.User{Id: "2", FirstName: "João", LastName: "Anjos", DateOfBirth: 865296000})
	users = append(users, &product.User{Id: "3", FirstName: "João", LastName: "Anjos", DateOfBirth: 1549411200})

	for _, user := range users {
		if user.Id == userID {
			return user, nil
		}
	}

	return nil, fmt.Errorf("the user with id %s doesnt exists", userID)
}

func isBirthDay(userBD int64) bool {
	userBirth := time.Unix(userBD, 0)
	currentDate := time.Now()

	if userBirth.Month() == currentDate.Month() && userBirth.Day() == currentDate.Day() {
		return true
	}

	return false
}

func isBlackFriday() bool {
	currentDate := time.Now()
	if currentDate.Month() == 11 && currentDate.Day() == 25 {
		return true
	}

	return false
}

type productServer struct{}

func (s productServer) WithDiscount(ctx context.Context, ids *product.Ids) (*product.ProductWithDiscount, error) {
	foundedProduct, err := getProductByID(ids.ProductID)
	if err != nil {
		return nil, fmt.Errorf("could not found product: %s", err)
	}
	foundedUser, err := getUserByID(ids.UserID)
	if err != nil {
		return nil, fmt.Errorf("could not found user: %s", err)
	}

	totalDiscountPct := float32(0.0)
	if isBirthDay(foundedUser.DateOfBirth) {
		totalDiscountPct += 0.05
	}

	if isBlackFriday() {
		totalDiscountPct += 0.1
	}

	if totalDiscountPct > 0.1 {
		totalDiscountPct = 0.1
	}

	productDiscount := &product.Discount{
		Prc:          totalDiscountPct * 100,
		ValueInCents: int32(float32(foundedProduct.PriceInCents) * totalDiscountPct),
	}

	productWithDiscount := &product.ProductWithDiscount{
		Id:           foundedProduct.Id,
		PriceInCents: foundedProduct.PriceInCents,
		Title:        foundedProduct.Title,
		Description:  foundedProduct.Description,
		Discount:     productDiscount,
	}

	return productWithDiscount, nil
}

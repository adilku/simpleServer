package store

import validation "github.com/go-ozzo/ozzo-validation"
import _ "github.com/go-ozzo/ozzo-validation/is"

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type JwtToken struct {
    Token string `json:"token"`
}

type Exception struct {
    Message string `json:"message"`
}

// Product represents an e-comm item
type Product struct {
	ID     int 		 	 `bson:"_id"`
	Title  string        `json:"title"`
	Image  []string      `json:"image"`
	Price   uint64       `json:"price"`
	Description  string   `json:"Description"`
}

func(product *Product) Validate() error {
	return validation.ValidateStruct(product,
		validation.Field(&product.Image, validation.Length(0, 3)),
		validation.Field(&product.Title, validation.Length(0, 200)),
		validation.Field(&product.Description, validation.Length(0, 1000)))
}

// Products is an array of Product objects
type Products []Product
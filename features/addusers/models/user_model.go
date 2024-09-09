package models

type User struct {
	UserId      string   `json:"userId" bson:"userId"`
	Password    string   `json:"password" bson:"password"`
	Name        string   `json:"name" bson:"name"`
	Email       string   `json:"email" bson:"email"`
	PhoneNumber string   `json:"phone_number" bson:"phone_number"`
	Role        string   `json:"role" bson:"role"`
	Address     Address  `json:"address" bson:"address"`
	Orders      []Orders `json:"orders" bson:"orders"`
	CreatedAt   string   `json:"created_at" bson:"created_at"`
}

type Address struct {
	Street  string `json:"street" bson:"street"`
	City    string `json:"city" bson:"city"`
	State   string `json:"state" bson:"state"`
	PinCode string `json:"pincode" bson:"pincode"`
}

type Orders struct {
	OrderId string `json:"order_id" bson:"order_id"`
}

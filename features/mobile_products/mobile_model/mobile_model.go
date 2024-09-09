package mobilemodel

type MobileInfo struct {
	MobileID      string              `json:"mobile_id" bson:"mobile_id"`
	Brand         string              `json:"brand" bson:"brand"`
	Model         string              `json:"model" bson:"model"`
	Varients      []DeviceVariant     `json:"varients" bson:"varients"`
	Specification MobileSpecification `json:"specification" bson:"specification"`
	Features      []string            `json:"features" bson:"features"`
	AddedAt       string              `json:"added_at" bson:"added_at"`
}

type DeviceVariant struct {
	VariantID   string `json:"varient_id" bson:"varient_id"`
	Color       string `json:"color" bson:"color"`
	Storage     string `json:"storage" bson:"storage"`
	Price       string `json:"price" bson:"price"`
	Condition   string `json:"condition" bson:"condition"`
	IsAvailable bool   `json:"is_available" bson:"is_available"`
}

type MobileSpecification struct {
	ScreenSize string `json:"screen_size" bson:"screen_size"`
	Processor  string `json:"processor" bson:"processor"`
	Battery    string `json:"battery" bson:"battery"`
}

/*
{
  "mobileId": "unique_mobile_id",
  "brand": "Samsung",
  "model": "Galaxy S21",
  "variants": [
    {
      "variantId": "variant1",
      "color": "Phantom Gray",
      "storage": "128 GB",
      "price": 29999,
      "condition": "New",
      "available": true
    },
    {
      "variantId": "variant2",
      "color": "Phantom White",
      "storage": "256 GB",
      "price": 34999,
      "condition": "New",
      "available": true
    }
  ],
  "specifications": {
    "screenSize": "6.2 inches",
    "processor": "Exynos 2100",
    "ram": "8 GB",
    "battery": "4000 mAh"
  },
  "features": ["5G", "Water Resistant", "Fast Charging"],
  "addedAt": "2024-09-08T00:00:00Z"
}

*/

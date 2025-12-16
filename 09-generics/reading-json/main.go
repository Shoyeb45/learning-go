package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ContactInfo struct {
    Name  string `json:"name"`
    Email string `json:"email"`
}

type PurchaseInfo struct {
    Name  string  `json:"name"`
    Price float32 `json:"price"`
    Info  string  `json:"info"`
}


func main() {
    contacts := loadJson[ContactInfo]("./contactInfo.json")
    fmt.Printf("\n%+v\n", contacts)

    purchases := loadJson[PurchaseInfo]("./purchaseInfo.json")
    fmt.Printf("\n%+v\n", purchases)
}


func loadJson[T ContactInfo | PurchaseInfo](filePath string) []T {
    data, err := os.ReadFile(filePath)
    if err != nil {
        panic(err)
    }

    var loaded []T
    if err := json.Unmarshal(data, &loaded); err != nil {
        panic(err)
    }

    return loaded
}

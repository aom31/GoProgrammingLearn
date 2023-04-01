package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	_ mysql.Open()
	_ gorm.Open()
}

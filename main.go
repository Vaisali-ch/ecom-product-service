// 2025 file
package main

import (
	"github.com/vaisali-ch/ecom-product-service/config"
	"github.com/vaisali-ch/ecom-product-service/routes"
)

func main() {
	config.SetupProductDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}

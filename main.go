package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	router.POST("/", func(c *gin.Context) {
		var berat, jarak float64

		beratStr := c.PostForm("berat")
		jarakStr := c.PostForm("jarak")

		fmt.Sscanf(beratStr, "%f", &berat)
		fmt.Sscanf(jarakStr, "%f", &jarak)

		var layanan string

		if berat < 5 {
			if jarak < 200 {
				layanan = "Pengiriman menggunakan layanan ekspres"
			} else {
				layanan = "Pengiriman menggunakan layanan reguler"
			}
		} else {
			if jarak < 200 {
				layanan = "Pengiriman menggunakan layanan reguler"
			} else {
				layanan = "Pengiriman menggunakan layanan kargo"
			}
		}

		c.HTML(200, "hasil.html", gin.H{
			"layanan": layanan,
		})
	})

	router.Run(":8080")
}

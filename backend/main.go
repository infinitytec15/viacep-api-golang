package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Next()
	})

	router.GET("/cep/:cep", func(c *gin.Context) {
		cep := c.Param("cep")

		response, err := http.Get(fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar o CEP"})
			return
		}
		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler a resposta"})
			return
		}

		if string(body) == `{"erro":true}` {
			c.JSON(http.StatusNotFound, gin.H{"error": "CEP não encontrado"})
			return
		}

		var endereco ViaCEP
		err = json.Unmarshal(body, &endereco)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao decodificar o JSON"})
			return
		}

		c.JSON(http.StatusOK, endereco)
	})

	router.Run(":8080")
}

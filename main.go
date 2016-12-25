package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	// "fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	// "net/http/httputil"
)

type GithubRepos struct {
	Id          int    `json:"id"`
	FullName    string `json:"name"`
	Description string `json:"description"`
	HtmlUrl     string `json:"html_url"`
	UpdatedAt   string `json:"updated_at"`
}

func errorResponse(c *gin.Context, resultMessage string, statusCode int) {
	c.JSON(statusCode, gin.H{
		"result":  "error",
		"message": resultMessage,
	})
}

func main() {

	router := gin.Default()

	router.Use(cors.Middleware(cors.Config{
		Origins:        "*",
		Methods:        "GET, PUT, POST, DELETE",
		RequestHeaders: "Origin, Authorization, Content-Type",
		ExposedHeaders: "",
		// MaxAge: 50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.GET("/github/repos", func(c *gin.Context) {
		client := &http.Client{}
		request, error := http.NewRequest("GET", "https://api.github.com/users/masa69/repos?sort=updated", nil)
		// dump, _ := httputil.DumpRequestOut(request, true)
		// fmt.Printf("%s", dump)
		response, error := client.Do(request)
		// fmt.Println(response.Header["Content-Type"])
		// fmt.Println(response.StatusCode)
		// fmt.Println(response.Status)
		// dumpResp, _ := httputil.DumpResponse(response, true)
		// fmt.Printf("%s", dumpResp)
		if error != nil {
			errorResponse(c, response.Status, response.StatusCode)
			return
		}
		defer response.Body.Close()

		body, error := ioutil.ReadAll(response.Body)

		if error != nil {
			// fmt.Println("error:", error)
			errorResponse(c, "error", 500)
			return
		}

		bytes := []byte(body)

		var g []GithubRepos

		jsonError := json.Unmarshal(bytes, &g)

		if jsonError != nil {
			// fmt.Println("error:", jsonError)
			errorResponse(c, "error", 500)
			return
		}

		c.JSON(200, gin.H{
			"result":  "success",
			"message": "success",
			"body":    g,
		})
	})

	router.Run(":8088")
}

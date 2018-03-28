// To compile: GOOS=linux go build
package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	sleep(os.Getenv("SLEEP"))
	app := gin.Default()
	app.GET("/kill", kill)
	app.GET("/poison", poison)
	app.GET("/cure", cure)
	app.GET("/health", health)
	app.GET("/burn/:n", burn)

	app.Run(":8000")
}

var healthy bool = true

func poison(c *gin.Context) {
	healthy = false
}

func cure(c *gin.Context) {
	healthy = true
}

func kill(c *gin.Context) {
	os.Exit(1)
}

func health(c *gin.Context) {
	var code int
	if healthy {
		code = 200
	} else {
		code = 500
	}
	c.JSON(code, gin.H{
		"status":     healthy,
		"sleepiness": os.Getenv("SLEEP"),
		"host":       os.Getenv("HOSTNAME"),
	})
}

func burn(c *gin.Context) {
	n := c.Param("n")

	i, err := strconv.ParseInt(n, 10, 64)
	if err != nil {
		c.JSON(500, gin.H{"error": "Not a number: " + n})
		return
	}
	c.JSON(200, gin.H{"result": GetNthPrime(i)})
}

func sleep(env string) {
	if i, err := strconv.ParseInt(env, 10, 64); err == nil {
		fmt.Printf("Sleeping for " + env)
		time.Sleep(time.Duration(i) * time.Second)
	} else {
		fmt.Printf("Failed sleeping for " + env)
	}
}

func isPrime(n int64, primes *list.List) bool {
	if primes == nil {
		return false
	}

	for e := primes.Front(); e != nil; e = e.Next() {
		if n%e.Value.(int64) == 0 {
			return false
		}
	}

	return true
}

func GetNthPrime(n int64) int64 {
	primes := list.New()
	primes.PushBack(int64(2))

	for i := int64(3); int64(primes.Len()) < n; i += int64(2) {
		if isPrime(i, primes) {
			primes.PushBack(i)
		}
	}

	return primes.Back().Value.(int64)
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

type Person struct {
	FirstName string `form:"firstname" json:"firstname" binding:"required"`
	LastName  string `form:"lastname" json:"lastname" binding:"required"`
	Age       string `form:"age" json:"age" binding:"required"`
	Job       string `form:"job" json:"job" binding:"required"`
}

var (
	store = make(map[string]Person)
	mutex = &sync.Mutex{}
)

func (p Person) GetHash() string {
	return strings.ToLower(p.FirstName + "," + p.LastName)
}

func NewPerson(firstname, lastname, job, age string) *Person {
	person := Person{
		FirstName: firstname,
		LastName:  lastname,
		Job:       job,
		Age:       age,
	}

	return &person
}

func hello(c *gin.Context) {
	firstname := c.Param("firstname")
	lastname := c.Param("lastname")

	key := NewPerson(firstname, lastname, "", "").GetHash()
	if person, ok := store[key]; ok {
		c.String(http.StatusOK, "Hello %s %s; Job: %s; Age: %s", person.FirstName, person.LastName, person.Job, person.Age)
	} else {
		c.String(http.StatusNotFound, "%s %s is not registered", firstname, lastname)
	}
}

func register(c *gin.Context) {
	firstname := c.PostForm("firstname")
	if firstname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "firtname is required"})
		return
	}

	lastname := c.PostForm("lastname")
	if lastname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "lastname is required"})
		return
	}

	job := c.PostForm("job")
	if job == "" {
		job = "Unknown"
	}

	age := c.PostForm("age")
	if age == "" {
		age = "18"
	} else {
		match, _ := regexp.MatchString("^\\d+$", age)
		if !match {
			c.JSON(http.StatusBadRequest, gin.H{"message": "age should be integer"})
			return
		}
	}

	mutex.Lock()
	defer mutex.Unlock()

	key := NewPerson(firstname, lastname, "", "").GetHash()
	if _, exists := store[key]; exists {
		c.JSON(http.StatusConflict, gin.H{"message": fmt.Sprintf("%s %s registered before", firstname, lastname)})
		return
	}

	store[key] = *NewPerson(firstname, lastname, job, age)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s %s registered successfully", firstname, lastname)})
}

package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    // "strconv"
)

func GetPeople(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": people})
}

func GetPerson(c *gin.Context) {
    var param_id string
    param_id = c.Param("id")
    // param_id = strconv.Atoi(c.PostForm("id"))
    for _, item := range people {
        if item.ID == param_id {
            c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": item})
            return
        }
    }
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": &Person{}})
}

func CreatePerson(c *gin.Context) {
    var person Person
    param_id := c.Param("id")
    param_firstname := c.DefaultQuery("firstname", "Guest")
    param_lastname := c.DefaultQuery("lastname", "")
    // param_lastname := c.Query("lastname")
    // param_id = strconv.Atoi(c.PostForm("id"))
    person.ID = param_id
    person.Firstname = param_firstname
    person.Lastname = param_lastname
    people = append(people, person)
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": people})
}

func DeletePerson(c *gin.Context) {
    var param_id string
    param_id = c.Param("id")
    // param_id = strconv.Atoi(c.PostForm("id"))
    for index, item := range people {
        if item.ID == param_id {
            people = append(people[:index], people[index+1:]...)
            break
        }
        c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": people})
    }
}

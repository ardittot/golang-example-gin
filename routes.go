package main

func initializeRoutes() {

  // Handle the index route
  router.GET("/people", GetPeople)
  router.GET("/people/:id", GetPerson)
  router.POST("/people/:id", CreatePerson)
  router.DELETE("/people/:id", DeletePerson)

}

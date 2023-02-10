package main

import (
	"DB/psql"
	"DB/verify"
)
func main()  {
  psql.ConnectDBB() // Evolution DB Connection
  psql.ConnectDBC() // PHM DB Connection
  
  verify.Check()
}

/*

--------------- GIN ----------------
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {

	//set a router
	router := gin.Default()
	router.GET("/ping", PingGet)

	return router
}

//define the context of the router
func PingGet(c *gin.Context) {
	c.JSON(http.StatusOK, map[int]string{
		10: "Hello",
		11: "World",
	})
}

//run the router
func main() {
	router := createRouter()
	router.Run()
}
*/

package main

import (
	"fmt"
)

type michael struct {
	name string
}

func main() {
	mike := michael{
		name: "Michael Lazuardi",
	}

	mike.changeName("mike")
	fmt.Println(mike)

	mikePointer := &mike
	mikePointer.changeNamePtr("kontol")
	fmt.Println(mikePointer)

}

func (m michael) changeName(newName string) {
	fmt.Println(newName + " is your new name")
}

func (mikePointer *michael) changeNamePtr(newPointerName string) {
	(*mikePointer).name = newPointerName
}

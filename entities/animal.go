package entities

/* Use case layer only contain business logic of application
this layer will communicate with datastore layer with interface
*/

// animal entities accessible in other package because its uppercased!
type Animal struct {
	ID   int
	Name int
	Age  int
}

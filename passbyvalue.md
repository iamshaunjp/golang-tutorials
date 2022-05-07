## Pass by Value
In this section We Will Discuss What Does Mean "Pass by Value"
  We have 2 Type of Values
  
  Pointer Wrapper Values
 - Map
 - Slices
 - Functions

Non-Pointer Values

- Strings
- Int
- Floats
- Booleans
- Arrays
- Stuctures

For **non-pointers**, we need to redefine separately, for example

Whenever we try to change it with a function parameter, a variable copy of a new parameter is created in the computer memory for that function, so it does not change even if we try to change the value of the variable in our next codes. due to scope range
 
 ```go
 package main
 import "fmt"
 
func changevalue(i int){
i = 20
}

func main(){
i := 5
changevalue(i)

fmt.Println("Value:",i) // Value: 5
}
```
If you ask how this is for the **pointer wrapper value**, they update the memory address directly instead of creating a copy for the variable.

```go
 package main
 import "fmt"
 
func changevalue(y map[string]float64){
y["baklava"]: 15.5
}

func main(){

menu := [string]float64 {
"pie": 35.6
"ice cream": 20.4
}

fmt.Println(menu) // map[ice cream:20.4 pie:35.6]
changevalue(menu)
fmt.Println(menu) // map[baklava:15.5 ice cream:20.4 pie:35.6]
}
```
Thanks for Reading 
*WinneRose*

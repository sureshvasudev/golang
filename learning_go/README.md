## GO Notes

#### About Package "main"
The package main must hold the main method, which is the entry point to the application.

#### The Zero Value
**Go** assigns a default **zero value** to any variable that is declared but not assigned a value

#### Explicit Type Conversion
Go doesn’t allow automatic type promotion between variables

#### Literals Are Untyped
 Literals in Go are untyped. Go is a practical language, and it makes sense to avoid forcing a type until the developer specifies one
 Example: var x float64 = **10**
 var x float64 = **10 * 3**
 
#### Variable Declaration
*   var x, y = 10, "hello" 
*   x, y := 10, "hello"
*   var x, y int = 10, 20
*   var (
    x    int
    y        = 20
    z    int = 30
    d, e     = 40, "hello"
    f, g string
)

#### Contant
*   const z = 20 * 10
*   const zz int = 20 * 10
Go doesn’t provide a way to specify that a value calculated at runtime is immutable. so it cannot be assigned to a constant 
Example: var x,y =5
const c = x* y **Compilation Error**

## Commands
**go mod init {module_name}** :  To create a module
**go build {*.go}**: To compile the file
**file.exe or file** : To run the exe file created by go build
**go build -o hello** : To create the exe in a custom name
**go fmt** : To format the source file automatically
**go fmt ./...**: To apply formatting in all files in the directory and sub directories
**go vet**: To check for syntax errors




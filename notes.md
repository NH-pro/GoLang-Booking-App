1. Make sure to have Go Extension installed in VS Code.

2. Initialize your Go project or "module"
    - In Terminal, type: go mod init [name of project]

3. All our code must belong to a "package"
    - The first statement in a Go file must be "package"

4. Go projects need to know where to start executing the code.
    - aka, the "Main" function of the program
    - you can only have ONE main function

5. Go uses packages that contain built in functions
    - to print something out, you need to import the "fmt" package and then use the "Print()" function.

6. To execute project,
    - In Terminal, type: go run main.go
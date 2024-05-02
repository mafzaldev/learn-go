### Learn GO

#### Install GO

1. Download the latest version of GO from [here](https://go.dev/dl/).
2. Install the downloaded file.
3. Check if GO is installed by running `go version` in the terminal.

#### Initialize a GO module

1. Create a new directory for your project.
2. Run `go mod init <module-name>` in the terminal.

#### Initialize a GO package

1. Create a new directory for your package.

   ```
   mkdir <package-name>
   cd <package-name>
   ```

2. Create a new file with the name of the package.

   ```
   touch <package-name>.go
   ```

3. Add the package name at the top of the file.

   ```
   package <package-name>
   ```

4. Write the code for the package.

   ```
   package <package-name>
   import "fmt"

   func main() {
      fmt.Println("Hello World!")
   }
   ```

#### Build/Run a GO program

1. Run `go build <complete-file-path>.go` in the terminal.
2. Run `go run <complete-file-path>.go` in the terminal.

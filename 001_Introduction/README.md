# GoLang Introduction 

Go is a statically typed, compiled programming language. It is popular for cloud-based or server-side applications.



## Introduction.pdf 
An introduction to the GO programming language, its charactersticks, use cases and applications.



## Installation 

Download your platform specific binary from - https://go.dev/dl/ and install it.

![goins_1.png](/.pics/001/goins_1.png)
![goins_2.png](/.pics/001/goins_2.png)

To verify the installation, go to your terminal and type and enter "go".

You should see and output if the installation was done correctly.

![verify.png](/.pics/001/verify.png)

Also, i am using VSCode as my code editor, you can install an extension named "go" (by - Go Team at Google) to make life a little easy.


![ext.png](/.pics/001/ext.png)

Click on install if any pop-up shows up  



## Hello World 

Writing your first GO Program 

- Create a folder in which you will keep your code (i made one named GO-Hello)

- In golang we have to first initilaize our project. To do so open your integrated terminal in VS Code and use the command

  - **go mod init module-name**

    - It will create a file named go.mod


```
go mod init GO-Hello
```
![init.png](/.pics/001/init.png)

![struc.png](/.pics/001/struc.png)

- All our code must belong to a package (main).

- The **package main** tells the Go compiler that the package should compile as an executable program instead of a shared library


```
package main

import "fmt"

func main() {
	fmt.Print("Hello World")

}
```

- for one go application / program, you can have only one main function as it is the entrypoint to your code 

- Go code needs to oragnized in packages 

- **import "fmt"**  - a go package for standard IO.

- Run your go file using the command 
  - **go run filename.go**

```
go run main.go
```

![run.png](/.pics/001/run.png)

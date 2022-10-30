# GoLang Introduction 

Go is a statically typed, compiled programming language. It is popular for cloud-based or server-side applications.



## Introduction.pdf 
An introduction to the GO programming language, its charactersticks, use cases and applications.



## Installation 

download your platform specific binary from - https://go.dev/dl/ and install it.

![goins_1.png](/.pics/001/goins_1.png)
![goins_2.png](/.pics/001/goins_2.png)

to verify the installation, go to your terminal and type and enter "go".

You should see and output if the installation was done correctly.

![verify.png](/.pics/001/verify.png)

also, i am using VSCode as my code editor, you can install an extension named "go" (by - Go Team at Google) to make life a little easy.


![ext.png](/.pics/001/ext.png)

click on install if any popup shows up  



## Hello World 

Writing your first GO Program 

- create a folder in which you will keep your code (i made one named GO-Hello)

- in golang we have to first initilaize our project. To do so open your integrated terminal in the code editor and use the command

  - **go mod init module-name**

    - it will create a file named go.mod


```
go mod init GO-Hello
```
![init.png](/.pics/001/init.png)

![struc.png](/.pics/001/struc.png)

- all our code must belong to a package.

![code.png](/.pics/001/code.png)

- for one go application, you can have only one main function as it is the entrypoint to your code 

- go code needs to oragnized in packages 

- **import "fmt"**  - a go package for standard IO.

- run your go file using the command 
  - **go run filename.go**

```
go run main.go
```

![run.png](/.pics/001/run.png)

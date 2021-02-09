



### Process:
##### Init Go module & download required dependencies:
* [How to init go module & dependencies](https://medium.com/@adiach3nko/package-management-with-go-modules-the-pragmatic-guide-c831b4eaaf31)
* this is basically `go mod init` & `go get` command

##### Add CLI :
* `go get github.com/spf13/cobra/cobra`  & `cobra init --pkg-name client-go` . Here client-go is our repository name
*  after running the previous command, a root.go file will be created inside cmd folder. No need to modify the root.go file
* create another go file inside the cmd folder. This is for our sub-command.
* Now implement your sub-command according to your need , Ex: to create sub-command `create-deploy`
    * first create variable `createCMD`
    * add the function inside the `RUN` section, you want to run for your sub-command `createCMD`  . Such as `createDeployment()` 
    * inside init() function , add `createCMD` to your rootCMD `rootCmd.AddCommand(createCMD)`

##### Build:
*`go build .` this will create a binary file named your project . such as `client-go`
* Or `go build -o ./app` this will create a binary file name `app`

##### RUN:
* `./client-go create-deploy` or  using flag `./client-go create-deploy -r=5 -i=pranganmajumder/go-basic-restapi:1.0.3`
* `./client-go get-deploy`
* `./client-go update-deploy` it'll create 1 replica and update the previous image. to check run again `./client-go get-deploy`
    * using flag`./client-go update-deploy -r=3 -i=pranganmajumder/go-basic-restapi:1.0.0` 
* `./client-go delete-deploy`
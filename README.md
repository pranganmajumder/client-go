



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
    * add the function inside the `RUN` section, you want to run for your sub-command `createCMD` ` 
    * inside init() function , add `createCMD` to your rootCMD `rootCmd.AddCommand(createCMD)`

##### Build:
* `go build .` this will create a binary file named your project . such as `client-go`
*  Or `go build -o ./app` this will create a binary file name `app`

##### RUN:
* `./app create new-deploy -r=5 -i=pranganmajumder/go-basic-restapi:1.0.3`
* `./app get list-deploy` will display all the running deployment
* `./app update update-deploy -r=3 -i=pranganmajumder/go-basic-restapi:1.0.0` 
* `./app delete delete-deploy`
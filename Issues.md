### 1. Go Module an Alternative for GOPATH:
  - https://insujang.github.io/2020-04-04/go-modules/
  - If our project has multiple packages then , the project should be put in GOPATH and should contain src folder , otherwise it won't run.
  - An alternative to this problem is to use GO Modules from Go 1.11


### 2. Setting GOPATH 
- post 1.13+
  - go env -w GOPATH=$HOME/go
- https://github.com/golang/go/wiki/SettingGOPATH
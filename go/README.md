# golang
## go mod init
```go
go mod init {pkg}
```

## go test
```go
go test -v ./...
```
```go
go test -v ./{file}
```

# choco
https://marcus116.blogspot.com/2019/02/chocolatey-chocolatey.html

## search
```java
// search command : <filter>
choco search git 
 
@example 
// search exact name 
choco search git -e
 
// search id only
choco search git --by-id-only
 
// search approved only
choco search git --approved-only
 
// search local only 
choco search git -lo
```

## list
```java
// list command : <filter>
choco list git 
 
@example 
// list local package 
choco list -lo
choco list --local-only
 
// list already package on windows
choco list -li
choco list -lai
```

## install
```java
// install command : <filter>
choco install git 
 
@example 
// confirm all prompts
choco install -y
 
// specific version
choco install git --version 2.20.1
 
// prerelease
choco install git --pre
```

## upgrade
```java
// upgrade command 
choco upgrade git 
 
@example 
// upgrade all 
choco upgrade all
 
// upgrade all expect some
choco upgrade all --except="'skype'"
```

## uninstall
```java
// uninstall command 
choco uninstall git 
 
@example 
// uninstall specific version 
choco uninstall git --version 2.20.1
```

## config
```java
// config command 
choco config 
 
@example 
// config list 
choco config list  
 
// get config 
choco config get cacheLocation
 
// set config
choco config set cacheLocation d:\temp\choco
 
// config help
choco config -h
```

## help
```java
// help command 
choco -h 
```
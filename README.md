# Project Overview
This project is a DSA assignment for TI'23 Student at YARSI University. It involves making RESTful API in Golang about a specific feature in an app, this time is "Buy Thing"
## Golang Installation
You can find the relevant installation files at https://golang.org/dl/.

Follow the instructions related to your operating system. To check if Go was installed successfully, you can run the following command in a terminal window:

```bash
go version
```
Which should show the version of your Go installation.
## Run Locally

Assuming you have cloned the repo and installed [Go](https://golang.org/dl/), you can run the the code via a command prompt/terminal. But first you had to download and install [gin (Web framework for Golang)](github.com/gin-gonic/gin), [gorm (ORM library for Golang)](https://gorm.io/), and gorm's mysql driver, by running:
```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

Then start the app by running:
```bash
go run .\server.go
```

# Go-Pandora
---
## Introduction
Go-Pandora is a Golang library designed to simplify interactions with the Pandora Secrets Management System. It provides a lightweight client interface for securely managing users and secrets. With Go-Pandora, developers can seamlessly integrate Pandora's services into their Go applications.

## User Guide
Go-Pandora enables interaction with Pandora's secrets management services through simple function calls. This guide briefly introduces the main concepts of working with the library.

### Installation
To install Go-Pandora in your project, use the following command:
```shell
go get github.com/MustafaAbdulazizHamza/go-pandora
```
### Pandora Client
To initialize a Pandora client, use the NewPandoraClient() function provided by the library and pass the required parameters:
```go
package main

import goPandora "github.com/MustafaAbdulazizHamza/go-pandora"

func main() {
	pandora := goPandora.NewPandoraClient("https://IP:Port", "username", "password", "private_key.pem", "public_key.pem")
}
```
### User Management
User management is a crucial part of any system. In Pandora, user management is primarily handled by the root user, with some exceptions allowing individual users to update their own passwords.
1. To add a new user:
```go
	if err := pandora.AddUser("usernam1", "password1"); err != nil {
		os.Exit((1))
	}
```
2. To update user credentials:
```go
	if err := pandora.UpdateUserCredentials("usernam1", "password1"); err != nil {
		os.Exit((1))
	}
```
3. To delete a user:

```go
	if err := pandora.DeleteUser("usernam1"); err != nil {
		os.Exit((1))
	}
```
### The Secrets Management
Pandora was initially developed as a secrets management system, with its primary responsibility being to allow users to securely add, retrieve, update, and delete their secrets in a centralized manner through simple API calls.
1. To add a secret:
```go
	if err := pandora.PostSecret("secret_id", "secret"); err != nil {
		os.Exit(1)
	}
```
2. To retrieve a secret:
```go
	if secret, err := pandora.GetSecret("secret_id"); err != nil {
		os.Exit(1)
	} else {
		fmt.Println(secret)
	}
```
3. To update a secret:
```go
	if err := pandora.UpdateSecret("secret_id", "secret1"); err != nil {
		os.Exit(1)
	}
```
4. To delete a secret:
```go
	if err := pandora.DeleteSecret("secret_id"); err != nil {
		os.Exit(1)
	}
```

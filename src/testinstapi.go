package main

import (
    "fmt"
    "github.com/yanatan16/golang-instagram/instagram"
//    "net/url"
)

// ExampleNew sets up the whole instagram API
func main() {
    client_key := "cf8da68c641f48e2a5135c820c0a2b51"
    secret := "c7751614139940ada91f1593e9f96a54"
    auth_token := "6326791597.cf8da68.9a4f55542f124443a264eab3d61f7e83"
    apiAuthenticatedUser := instagram.New(client_key, secret, auth_token, true)
    if ok, err := apiAuthenticatedUser.VerifyCredentials(); !ok {
        panic(err)
    }
    fmt.Println("Successfully created instagram.Api with user credentials")
}

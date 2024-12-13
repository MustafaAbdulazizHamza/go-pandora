package goPandora

import "fmt"

func (c *PandoraClient) AddUser(username, password string) error {
	err := userOperation(c.url, c.username, c.password, username, password, "POST")
	return err
}
func (c *PandoraClient) DeleteUser(username string) error {
	err := userOperation(c.url, c.username, c.password, username, " ", "DELETE")
	return err

}
func (c *PandoraClient) UpdateUserCredentials(username, password string) error {
	err := userOperation(c.url, c.username, c.password, username, password, "PATCH")
	return err

}
func userOperation(url, authusername, authpassword, username, password, operation string) error {
	user := user{
		Username: username,
		Password: password,
	}
	body, err := toJson(user)
	if err != nil {
		return err
	}
	res, err := sendHTTPRequest(operation, url+"user", authusername, authpassword, string(body))
	if err != nil {
		return err
	}
	response, err := ParseResponse(res)
	if err != nil {
		fmt.Println("f")
		return err
	}
	if response.Status == "200" {
		return nil
	}
	return fmt.Errorf("user operation failed with status code: %s", response.Status)

}

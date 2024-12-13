package goPandora

import "fmt"

func secretGetDelete(url, username, password, secretID, privateKey string, isDelete bool) (string, error) {
	req := RequestedSecret{
		SecretID: secretID,
	}
	body, err := toJson(req)
	if err != nil {
		return "", err
	}
	method := "GET"
	if isDelete {
		method = "DELETE"
	}
	res, err := sendHTTPRequest(method, url+"secret", username, password, string(body))
	if err != nil {
		return "", err
	}
	response, err := ParseResponse(res)
	if err != nil {
		return "", err
	}
	if !isDelete && response.Status == "200" {
		response.Text, err = decryptWithPrivateKey(response.Text, privateKey)
		if err != nil {
			return "", err
		}
	}
	return response.Text, nil

}

func secretPostUpdate(url, username, password, secret, secretID, publicKey string, isUpdate bool) error {
	secret, err := encryptWithPublicKey(secret, publicKey)
	if err != nil {
		return err
	}
	sec := Secret{SecretID: secretID, Secret: secret}
	body, err := toJson(sec)
	if err != nil {
		return err
	}
	method := "POST"
	if isUpdate {
		method = "PATCH"
	}
	res, err := sendHTTPRequest(method, url+"secret", username, password, string(body))
	if err != nil {
		return err
	}
	response, err := ParseResponse(res)
	if err != nil {
		return err
	}
	if response.Status == "200" {
		return nil
	}
	return fmt.Errorf("secret update failed with status code: %s", response.Status)
}

func (c *PandoraClient) GetSecret(secretID string) (secret string, err error) {
	secret, err = secretGetDelete(c.url, c.username, c.password, secretID, c.privateKey, false)
	return secret, err
}

func (c *PandoraClient) PostSecret(secretID, secret string) error {
	err := secretPostUpdate(c.url, c.username, c.password, secret, secretID, c.publicKey, false)
	return err

}
func (c *PandoraClient) DeleteSecret(secretID string) error {
	_, err := secretGetDelete(c.url, c.username, c.password, secretID, "", true)
	return err
}
func (c *PandoraClient) UpdateSecret(secretID, secret string) error {
	err := secretPostUpdate(c.url, c.username, c.password, secret, secretID, c.publicKey, true)
	return err
}

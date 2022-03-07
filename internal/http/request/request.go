package request

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	clickupError "github.com/End313234/clickup-wrapper/internal/errors"
)

func MakeRequest(req CustomRequest) error {
	var client http.Client

	request, _ := http.NewRequest(req.Method, req.URL, nil)
	request.Header.Add("Authorization", req.AccessToken)
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	data, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode != 200 {
		var err clickupError.Error
		json.Unmarshal(data, &err)
		return errors.New(err.Err)
	}

	err = json.Unmarshal(data, req.Value)
	if err != nil {
		return err
	}

	return nil
}

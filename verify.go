package kickbox

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Verify ...
func (c *Client) Verify(email string) (*Result, *Error) {
	req, err := http.NewRequest("GET", "https://api.kickbox.com/v2/verify", nil)

	if err != nil {
		return nil, &Error{err: err}
	}

	// Add query to the ur;
	query := req.URL.Query()
	query.Add("email", email)
	query.Add("api_key", c.apiKey)
	req.URL.RawQuery = query.Encode()

	req.Header.Set("User-Agent", "kickbox-go/1.0.0 (https://github.com/kickboxio/kickbox-go)")

	response, err := c.client.Do(req)

	if err != nil {
		return nil, &Error{
			err:  err,
			Code: -1,
		}
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, &Error{
			err:     err,
			Message: "Failed to read response",
			Code:    -2,
		}
	}

	result := &Result{}

	if err := json.Unmarshal(body, result); err != nil {
		return nil, &Error{
			err:     err,
			Message: "Failed to convert response body",
			Code:    -3,
		}
	}

	if response.StatusCode != 200 {
		err := &Error{
			err:     err,
			Code:    response.StatusCode,
			Message: result.Message,
			Success: result.Success,
		}

		return nil, err
	}

	return result, nil
}

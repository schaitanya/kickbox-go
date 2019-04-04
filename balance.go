package kickbox

// func (c *Client) Balance() *Error {
// 	req, err := http.NewRequest("GET", "https://api.kickbox.com/v2/balance", nil)

// 	if err != nil {
// 		return nil, &Error{err: err}
// 	}

// 	// Add query to the ur;
// 	query := req.URL.Query()
// 	query.Add("email", email)
// 	query.Add("api_key", c.apiKey)
// 	req.URL.RawQuery = query.Encode()

// 	req.Header.Set("User-Agent", "kickbox-go/1.0.0 (https://github.com/kickboxio/kickbox-go)")

// 	response, err := c.client.Do(req)
// }

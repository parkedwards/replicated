package kotsclient

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c *VendorV3Client) RemoveKOTSRegistry(endpoint string) error {
	err := c.DoJSON("DELETE", fmt.Sprintf(`/v3/external_registry/%s`, url.QueryEscape(endpoint)), http.StatusNoContent, nil, nil)
	if err != nil {
		return err
	}
	return nil
}

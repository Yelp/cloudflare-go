package cloudflare

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

type CustomPage struct {
	URL            string    `json:"url"`
	State          string    `json:"state"`
	Description    bool      `json:"description"`
	PreviewTarget  string    `json:"preview_target"`
	RequiredTokens []string  `json:"required_tokens"`
	CreatedOn      time.Time `json:"created_on"`
	ModifiedOn     time.Time `json:"modified_on"`
}

// CustomPagesResponse is the API response, containing an array of CustomPages.
type CustomPagesResponse struct {
	Success  bool         `json:"success"`
	Errors   []string     `json:"errors"`
	Messages []string     `json:"messages"`
	Result   []CustomPage `json:"result"`
}

// CustomPagesDetailResponse is the API response, containing an single CustomPage.
type CustomPageDetailResponse struct {
	Success  bool       `json:"success"`
	Errors   []string   `json:"errors"`
	Messages []string   `json:"messages"`
	Result   CustomPage `json:"result"`
}

// ListCustomPages lists custom pages associated with a zone
// API reference:
//  https://api.cloudflare.com/#custom-pages-for-a-zone-available-custom-pages
//  GET /zones/:zone_identifier/custom_pages
func (api *API) ListCustomPages(zoneID string) ([]CustomPage, error) {
	uri := "/zones/" + zoneID + "/custom_pages"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []CustomPage{}, errors.Wrap(err, errMakeRequestError)
	}
	var r CustomPagesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []CustomPage{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// CustomPageDetails returns the details for a custom page.
// API reference:
//  https://api.cloudflare.com/#custom-pages-for-a-zone-custom-page-details
//  GET /zones/:zone_identifier/custom_pages/:identifier
func (api *API) CustomPageDetails(zoneID, pageID string) (CustomPage, error) {
	uri := "/zones/" + zoneID + "/custom_pages/" + pageID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return CustomPage{}, errors.Wrap(err, errMakeRequestError)
	}
	var r CustomPageDetailResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return CustomPage{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// Update a custom page.
// API reference:
//  https://api.cloudflare.com/#custom-pages-for-a-zone-update-custom-page-url
//  PUT /zones/:zone_identifier/custom_pages/:identifier
func (api *API) UpdateCustomPage(zoneID, pageID string, page CustomPage) (CustomPage, error) {
	uri := "/zones/" + zoneID + "/custom_pages/" + pageID
	res, err := api.makeRequest("PUT", uri, page)
	if err != nil {
		return CustomPage{}, errors.Wrap(err, errMakeRequestError)
	}
	var r CustomPageDetailResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return CustomPage{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

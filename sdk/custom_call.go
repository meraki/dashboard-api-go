package meraki

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type CustomCallService service

func (s *CustomCallService) GetCustomCall(ResourcePath string, QueryParms *map[string]string) (*resty.Response, error) {
	path := ResourcePath
	var response *resty.Response
	var err error
	request := s.client.R()
	request.SetHeader("Content-Type", "application/json").
		SetHeader("Accept", "application/json").
		SetError(&Error)
	if QueryParms != nil {
		request.SetQueryParams(*QueryParms)
	}
	response, err = request.Get(path)
	if err != nil {
		return nil, err

	}

	if response.IsError() {
		return response, fmt.Errorf("error with custom operation %s", ResourcePath)
	}

	return response, err
}

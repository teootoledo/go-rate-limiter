package controller

import (
	"errors"
	"net/http"
)

// AssertSuccess - Assert function to use with baloo library. It asserts that the response of a query has a http.StatusOK
func AssertSuccess(res *http.Response, req *http.Request) error {
	if res.StatusCode != http.StatusOK {
		return errors.New("error")
	}

	return nil
}

// AssertBadRequest - Assert function to use with baloo library. It asserts that the response of a query has a http.StatusBadRequest
func AssertBadRequest(res *http.Response, req *http.Request) error {
	if res.StatusCode != http.StatusBadRequest {
		return errors.New("error")
	}

	return nil
}

// AssertTooManyRequests - Assert function to use with baloo library. It asserts that the response of a query has a http.StatusTooManyRequests
func AssertTooManyRequests(res *http.Response, req *http.Request) error {
	if res.StatusCode != http.StatusTooManyRequests {
		return errors.New("error")
	}

	return nil
}

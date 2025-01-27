package api

import "github.com/go-resty/resty/v2"

var Client = resty.New()
var Fetch = Client.R
// Copyright 2016 JesusSlim. All Rights Reserved.
//
// Request class provided encapsulation of http.Request

package epackage

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type request struct {
	*http.Request                              // base http request
	paramMaxMemory int64                       // used in multipart/form-data form for file upload
	stash          map[interface{}]interface{} //some data u want to stash here
}

func Request(r *http.Request, paramMaxMemory ...int64) *request {
	if len(paramMaxMemory) > 0 {
		return &request{
			r,
			paramMaxMemory[0],
			map[interface{}]interface{}{},
		}
	} else {
		return &request{
			r,
			64 << 20,
			map[interface{}]interface{}{},
		}
	}
}

// the method of request is get or not
func (this *request) IsGet() bool {
	return this.Method == "GET"
}

// the method of request is post or not
func (this *request) IsPost() bool {
	return this.Method == "POST"
}

// is an ajax request or not
func (this *request) IsAjax() bool {
	return this.Header.Get("X-Requested-With") == "XMLHttpRequest"
}

// is an websocket request or not
func (this *request) IsWebsocket() bool {
	return this.Header.Get("Upgrade") == "websocket"
}

// return the params in request
func (this *request) Params() url.Values {
	if this.Form == nil {
		this.ParseFormAuto()
	}
	return this.Form
}

// return an string param by key,if it's nil,return the default value
func (this *request) GetString(key string, defaultValue ...string) string {
	v := this.Params().Get(key)
	if len(v) == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return v
}

// return an int param by key,if it's nil,return the default value
func (this *request) GetInt(key string, defaultValue ...int) int {
	v := this.Params().Get(key)
	vi, err := strconv.Atoi(v)
	if err != nil && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return vi
}

// return an int64 param by key,if it's nil,return the default value
func (this *request) GetInt64(key string, defaultValue ...int64) int64 {
	v := this.Params().Get(key)
	if len(v) == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	vi, _ := strconv.ParseInt(v, 10, 64)
	return vi
}

// return an float64 param by key,if it's nil,return the default value
func (this *request) GetFloat(key string, defaultValue ...float64) float64 {
	v := this.Params().Get(key)
	if len(v) == 0 && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	vi, _ := strconv.ParseFloat(v, 64)
	return vi
}

// parse form by func ParseMultipartForm or ParseForm
func (this *request) ParseFormAuto() error {
	if strings.Contains(this.Header.Get("Content-Type"), "multipart/form-data") {
		return this.ParseMultipartForm(this.paramMaxMemory)
	} else {
		return this.ParseForm()
	}
}

// return an string value by key in cookie,if it's nil,return the default value
func (this *request) GetCookie(key string, defaultValue ...string) string {
	v, err := this.Cookie(key)
	if err != nil || len(v.Value) == 0 {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}
	return v.Value
}

// stash a k-v:Stash(k,v)
// get a value by key in stash:Stash(k)
func (this *request) Stash(key interface{}, value ...interface{}) interface{} {
	if len(value) > 0 {
		this.stash[key] = value[0]
		return nil
	} else {
		return this.stash[key]
	}
}

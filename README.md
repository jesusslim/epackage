# epackage
Epackage provided some encapsulation of golang.为了更方便地使用,提供一些golang的简单封装.

## Request

Request.go提供针对http.request的一些封装 使得使用更为简便 并规避了一些post取参数的坑

### Get it:

	go get -u github.com/jesusslim/epackage
	
### Use it:

	//in your ServeHTTP,req is the param name of *http.Request
	mr := epackage.Request(req)
	fmt.Println("get:", mr.GetFloat("sid"))
	fmt.Println("cookie:", mr.GetCookie("test"))
	fmt.Println("stash", mr.Stash("sid"))
	fmt.Println("stash set", mr.Stash("sid", 123))
	fmt.Println("stash get", mr.Stash("sid"))
	if mr.IsGet() {
		//...		
	}

### Functions:

<b>IsGet</b>:the method of request is get or not

<b>IsPost</b>:the method of request is post or not

<b>IsAjax</b>:the method of request is post or not

<b>IsWebsocket</b>:is an websocket request or not

<b>Params</b>:return the params in request

<b>GetString</b>:return an string param by key,if it's nil,return the default value

<b>GetInt</b>:return an int param by key,if it's nil,return the default value

<b>GetInt64</b>:return an int64 param by key,if it's nil,return the default value

<b>GetFloat</b>:return an float64 param by key,if it's nil,return the default value

<b>ParseFormAuto</b>:parse form by func ParseMultipartForm or ParseForm

<b>GetCookie</b>:return an string value by key in cookie,if it's nil,return the default value

<b>Stash</b>:stash a k-v:Stash(k,v);get a value by key in stash:Stash(k)
	
# Overview 

Sample of consuming an authenticator plugin. Run the `Makefile` and you should see similar output:

```
$ make
Data received: map[pass:blahblahblah user:user]
The response was: Basic dXNlcjpibGFoYmxhaGJsYWg= 
```

The [main.go](main.go) file sends the basic auth plugin a username and password:

```
	// Prepare to pass some data
	dataIn := make(map[string]string)
	dataIn["user"] = "user"
	dataIn["pass"] = "blahblahblah"

    // Invoke authenticator plugin and output response
	headers := basicAuth.Authenticate(dataIn)
```

The [plugin.go](plugin.go) responds with an HTTP authentication header for the given username and password:


```
	// Create auth header and add to dataOut
	authHeader := "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
	dataOut = make(map[string]string)
	dataOut["Authorization"] = authHeader

```

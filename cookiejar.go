package fasthttp

// CookieJar is container of cookies
//
// This object is used to handle multiple cookies
type CookieJar map[string]*Cookie

// Set sets cookie using key-value
//
// This function can replace an existent cookie
func (cj *CookieJar) Set(key, value string) {
	setCookie(cj, key, value)
}

// SetBytesK sets cookie using key=value
//
// This function can replace an existent cookie.
func (cj *CookieJar) SetBytesK(key []byte, value string) {
	setCookie(cj, b2s(key), value)
}

// SetBytesV sets cookie using key=value
//
// This function can replace an existent cookie.
func (cj *CookieJar) SetBytesV(key string, value []byte) {
	setCookie(cj, key, b2s(value))
}

// SetBytesKV sets cookie using key=value
//
// This function can replace an existent cookie.
func (cj *CookieJar) SetBytesKV(key, value []byte) {
	setCookie(cj, b2s(key), b2s(value))
}

func setCookie(cj *CookieJar, key, value string) {
	c, ok := (*cj)[key]
	if !ok {
		c = fasthttp.AcquireCookie()
	}
	c.SetKey(key)
	c.SetValue(value)
	(*cj)[key] = c
}

// SetCookie sets cookie using its key.
//
// After that you can use Peek function to get cookie value.
func (cj *CookieJar) SetCookie(cookie *Cookie) {
	(*cj)[b2s(c.Key())] = cookie
}

// Peek peeks cookie value using key.
//
// This function does not delete cookie
func (cj *CookieJar) Peek(key string) *Cookie {
	return (*cj)[key]
}

// Release releases all cookie values.
func (cj *CookieJar) Release() {
	for k, v := range *cj {
		fasthttp.ReleaseCookie(v)
		delete(*cj, k)
	}
}

// ReleaseCookie releases a cookie specified by parsed key.
func (cj *CookieJar) ReleaseCookie(key string) {
	c, ok := (*cj)[key]
	if ok {
		fasthttp.ReleaseCookie(c)
		delete(*cj, key)
	}
}

// PeekValue returns value of specified cookie-key.
func (cj *CookieJar) PeekValue(key string) []byte {
	c, ok := (*cj)[key]
	if ok {
		return c.Value()
	}
	return nil
}

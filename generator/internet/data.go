package internet

var (
	freeEmailDomain = []string{"gmail.com", "yahoo.com", "hotmail.com"}

	tld = []string{"com", "com", "com", "com", "com", "com", "biz", "info", "net", "org"}

	userFormats = []string{"{{lastName}}.{{firstName}}",
		"{{firstName}}.{{lastName}}",
		"{{firstName}}",
		"{{lastName}}"}

	emailFormats = []string{"{{user}}@{{domain}}", "{{user}}@{{freeEmailDomain}}"}

	urlFormats = []string{"http://www.{{domain}}/",
		"http://{{domain}}/",
		"http://www.{{domain}}/{{slug}}",
		"http://www.{{domain}}/{{slug}}",
		"https://www.{{domain}}/{{slug}}",
		"http://www.{{domain}}/{{slug}}.html",
		"http://{{domain}}/{{slug}}",
		"http://{{domain}}/{{slug}}",
		"http://{{domain}}/{{slug}}.html",
		"https://{{domain}}/{{slug}}.html",
	}

	statusCodes        = []string{"100", "101", "102", "200", "201", "202", "203", "204", "205", "206", "207", "208", "226", "300", "301", "302", "303", "304", "305", "306", "307", "308", "400", "401", "402", "403", "404", "405", "406", "407", "408", "409", "410", "411", "412", "413", "414", "415", "416", "417", "418", "420", "422", "423", "424", "425", "426", "428", "429", "431", "444", "449", "450", "451", "499", "500", "501", "502", "503", "504", "505", "506", "507", "508", "509", "510", "511", "598", "599"}
	statusCodeMessages = []string{"Continue", "Switching Protocols", "Processing (WebDAV)", "OK", "Created", "Accepted", "Non-Authoritative Information", "No Content", "Reset Content", "Partial Content", "Multi-Status (WebDAV)", "Already Reported (WebDAV)", "IM Used", "Multiple Choices", "Moved Permanently", "Found", "See Other", "Not Modified", "Use Proxy", "(Unused)", "Temporary Redirect", "Permanent Redirect (experimental)", "Bad Request", "Unauthorized", "Payment Required", "Forbidden", "Not Found", "Method Not Allowed", "Not Acceptable", "Proxy Authentication Required", "Request Timeout", "Conflict", "Gone", "Length Required", "Precondition Failed", "Request Entity Too Large", "Request-URI Too Long", "Unsupported Media Type", "Requested Range Not Satisfiable", "Expectation Failed", "I'm a teapot (RFC 2324)", "Enhance Your Calm (Twitter)", "Unprocessable Entity (WebDAV)", "Locked (WebDAV)", "Failed Dependency (WebDAV)", "Reserved for WebDAV", "Upgrade Required", "Precondition Required", "Too Many Requests", "Request Header Fields Too Large", "No Response (Nginx)", "Retry With (Microsoft)", "Blocked by Windows Parental Controls (Microsoft)", "Unavailable For Legal Reasons", "Client Closed Request (Nginx)", "Internal Server Error", "Not Implemented", "Bad Gateway", "Service Unavailable", "Gateway Timeout", "HTTP Version Not Supported", "Variant Also Negotiates (Experimental)", "Insufficient Storage (WebDAV)", "Loop Detected (WebDAV)", "Bandwidth Limit Exceeded (Apache)", "Not Extended", "Network Authentication Required", "Network read timeout error", "Network connect timeout error"}
)

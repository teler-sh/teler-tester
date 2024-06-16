teler-tester

	This page serves as a demonstration of the
	teler-waf[1] package implementation, which is
	configured with default settings to test the
	resilience of web applications against common
	web attack threats.

	If you are able to successfully execute such
	as cross-site scripting (XSS) and pop-up an
	alert, kindly report it to us via the
	vulnerability report form[2] under its GitHub
	repository.

Demo

	Run from this directory

		$ go run .

	To build

		$ make build

	Live site: https://test.teler.sh


[1]: https://github.com/teler-sh/teler-waf
[2]: https://github.com/teler-sh/teler-waf/security/advisories/new
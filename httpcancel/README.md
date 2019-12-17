# httpcancel

This demonstrates how the context which is passed to an HTTP handler can be used
to detect that the remote HTTP client has gone away.

Use as follows: 

* start httpcancel demo
* curl http://127.0.0.1:8006/slow and hit *CTRL+C* within 10 seconds
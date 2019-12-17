# values

*context.WithValue* used for passing additional information via an API that cannot be modified.

In this case the example implements an HTTP middleware which retrieves authentication 
information and provides this information to HTTP handlers down the call chain.

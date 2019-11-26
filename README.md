# SWAN ACCESS CONTROL LIST FOR GIN GOLANG 

## How To Use

Install Swan

```
go get github.com/sofyan48/swan
```

Add ACL_ADDR in Your Dot Environment or Execute From Terminal

from dot env 
```
ACL_ADDR=172.19.0.0/24, 127.0.0.0/24
```
from os env
```
export export ACL_ADDR=172.19.0.0/24
or 
export export ACL_ADDR=172.19.0.0/24
```

Then Add SwanACL To gin Routes
```
api := r.Group("api")
{
    api.GET("/ping", swan.SwanACL(),ping.Ping)
}
```
example

```
u110:~/go/src/github.com/u110/gotour/daemon (pipeline *%) $ curl -X POST -H "Content-Type: application/json" -d '{"Name":"sensuikan1973", "Age":100}' localhost:8080
{"name":"sensuikan1973","age":1}
u110:~/go/src/github.com/u110/gotour/daemon (pipeline *%) $ curl -X POST -H "Content-Type: application/json" -d '{"name":"sensuikan1973", "Age":100}' localhost:8080
{"name":"sensuikan1973","age":2}
u110:~/go/src/github.com/u110/gotour/daemon (pipeline *%) $ curl -X POST -H "Content-Type: application/json" -d '{"name":"sensuikan1973x", "Age":100}' localhost:8080
{"name":"sensuikan1973x","age":3}
```

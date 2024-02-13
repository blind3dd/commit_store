# Simple Memory K/V Storage

---

go build cmd/commit_store/main.go\
./main

---

`message00=$(echo "message00" | base64);\`\
`curl --silent -XPOST http://localhost:8080/v1/commit -d '{"commit": {"value": "'$message00'"}}'\`
---
`message01=$(echo "message01" | base64);\`\
`curl --silent -XPOST http://localhost:8080/v1/commit -d '{"commit": {"value": "'$message01'"}}'\`
---
`message02=$(echo "message02" | base64);\`\
`curl --silent -XPOST http://localhost:8080/v1/commit -d '{"commit": {"value": "'$message02'"}}'\`
---

`curl --silent http://localhost:8080/v1/commit -d '{"offset": 0}'\`\
`{"commit":{"value":"bWVzc2FnZTAwCg==","offset":0}}`
---
`curl --silent http://localhost:8080/v1/commit -d '{"offset": 1}'\`\
`{"commit":{"value":"bWVzc2FnZTAxCg==","offset":1}}`
---
`curl --silent http://localhost:8080/v1/commit -d '{"offset": 2}'\`\
`{"commit":{"value":"bWVzc2FnZTAyCg==","offset":2}}`
---

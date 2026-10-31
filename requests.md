# Фактический прогон curl

```text
$ curl -sS -i --max-time 15 -X POST http://127.0.0.1:18011/api/v1/notes -H Content-Type: application/json -d {"title": "Первая заметка", "content": "Go REST"}
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Location: /api/v1/notes/1
Date: Sat, 31 Oct 2026 08:44:20 GMT
Content-Length: 149

{"id":1,"title":"Первая заметка","content":"Go REST","createdAt":"2026-10-31T08:44:20.36139Z","updatedAt":"2026-10-31T08:44:20.36139Z"}


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18011/api/v1/notes/1
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 31 Oct 2026 08:44:20 GMT
Content-Length: 149

{"id":1,"title":"Первая заметка","content":"Go REST","createdAt":"2026-10-31T08:44:20.36139Z","updatedAt":"2026-10-31T08:44:20.36139Z"}


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18011/api/v1/notes?q=Go
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 31 Oct 2026 08:44:20 GMT
Content-Length: 3

[]


$ curl -sS -i --max-time 15 -X PATCH http://127.0.0.1:18011/api/v1/notes/1 -H Content-Type: application/json -d {"title": "Go заметка"}
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 31 Oct 2026 08:44:20 GMT
Content-Length: 140

{"id":1,"title":"Go заметка","content":"Go REST","createdAt":"2026-10-31T08:44:20.36139Z","updatedAt":"2026-10-31T08:44:20.396395Z"}


$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18011/api/v1/notes?q=Go&limit=1&page=1
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sat, 31 Oct 2026 08:44:20 GMT
Content-Length: 142

[{"id":1,"title":"Go заметка","content":"Go REST","createdAt":"2026-10-31T08:44:20.36139Z","updatedAt":"2026-10-31T08:44:20.396395Z"}]


$ curl -sS -i --max-time 15 -X POST http://127.0.0.1:18011/api/v1/notes -H Content-Type: application/json -d {"title": " "}
HTTP/1.1 400 Bad Request
Content-Type: application/json; charset=utf-8
Date: Sat, 31 Oct 2026 08:44:20 GMT
Content-Length: 26

{"error":"invalid_input"}


$ curl -sS -i --max-time 15 -X DELETE http://127.0.0.1:18011/api/v1/notes/1
HTTP/1.1 204 No Content
Content-Type: application/json; charset=utf-8
Date: Sat, 31 Oct 2026 08:44:20 GMT



$ curl -sS -i --max-time 15 -X GET http://127.0.0.1:18011/api/v1/notes/1
HTTP/1.1 404 Not Found
Content-Type: application/json; charset=utf-8
Date: Sat, 31 Oct 2026 08:44:20 GMT
Content-Length: 22

{"error":"not_found"}

```

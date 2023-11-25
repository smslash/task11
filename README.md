# Задание 11

Описание задачи можете найти [здесь](https://handsomely-owl-b53.notion.site/11-5d4416745c7947ee951f879668615d14)

## 1-post-form

```console
POST /login HTTP/1.1
Host: www.example.com
Content-Type: application/x-www-form-urlencoded
Content-Length: 31

username=admin&password=secret
```

## 2-post-json

```console
POST /api/users HTTP/1.1
Host: www.example.com
Content-Type: application/json
Content-Length: 55

{
  "name": "John Doe",
  "email": "john.doe@example.com"
}
```

## 3-get-query-string

```console
GET /search?query=blue%20shoes HTTP/1.1
Host: www.example.com
```

## 💎 4-transfer-encoding

```console
POST /api/chunked HTTP/1.1
Host: www.example.com
Transfer-Encoding: chunked

7\r\n
Hello, \r\n
6\r\n
world!\r\n
0\r\n
\r\n
```
# 1 - Encode
```bash
curl localhost:8000/encode
```

# 2 - Decode

```bash
curl -v -XPOST -H "Content-Type: application/json" -d "[{\"FirstName\":\"Luiz Henrique\"},{\"FirstName\":\"Amora\"},{\"FirstName\":\"Rafael\"},{\"FirstName\":\"Daniel\"},{\"FirstName\":\"Todd\"}]" localhost:8000/decode
```
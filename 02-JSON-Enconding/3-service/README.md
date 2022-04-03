# Encode
```json
[
  {
    "FirstName": "Luiz Henrique"
  },
  {
    "FirstName": "Amora"
  },
  {
    "FirstName": "Rafael"
  },
  {
    "FirstName": "Daniel"
  },
  {
    "FirstName": "Todd"
  }
]  

```

```bash
curl -v localhost:8000/encode?name=Amora
```
# Decode

Linux
```bash
curl -v -XPOST -H "Content-Type: application/json"  -d '{"FirstName": "Rafael"}' localhost:8000/decode
```
Windows
```
curl -v -XPOST -H "Content-Type: application/json" -d "{\"FirstName\": \"Rafael\"}" localhost:8000/decode
```
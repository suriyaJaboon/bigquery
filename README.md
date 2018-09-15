# **Golang**

```
### 1
$ openssl genrsa -out private_key 2048
$ openssl rsa -in private_key -pubout > public_key.pub
```

```
openssl genpkey -algorithm RSA -out private_key -pkeyopt rsa_keygen_bits:4096
openssl rsa -pubout -in private_key -out public_key.pub
```
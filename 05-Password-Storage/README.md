# Storing Passwords

* never store passwords
* store one-way encryption "hash" MD5 (weak), SHA1 (weak), SHA2 (good), Bcrypt(good), scrypt(good)
* hash on the client and hash again on the server
* you should salt the password. It means to add a random string when hashing the password
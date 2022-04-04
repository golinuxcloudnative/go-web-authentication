# Storing Passwords

* never store passwords
* store one-way encryption "hash" MD5 (weak), SHA1 (weak), SHA2 (good), Bcrypt(good), scrypt(good)
* hash on the client and hash again on the server
* you should salt the password. It means to add a random string when hashing the password


Considerations:

* sha
    * always returns the same hash for the same input. 
    * similar input produces a completly diferent hash. So someone cannot associate
    * It has fixed hash size.
    * Different inputs could generate the same hash.

* bcrypt
    * add a salt to protect against rainbow table attacks. It means the same input does not return the same hash
    * default password hash for OpenBSD and some Linux distros
    * cost is the computation power to give the hash. Higher the cost more time to generate it.
        * MinCost     int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	    * MaxCost     int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	    * DefaultCost int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
    * structure
        ```
            $2a$04$Ltd/XUCCAAA9tGZOcq9WaureRQm8wxTW4U7q.2BvzKwIc/KdBYnP6
            \__/\/ \__________________________________/\_______________/
            Alg Cost               Salt                        Hash
        ```
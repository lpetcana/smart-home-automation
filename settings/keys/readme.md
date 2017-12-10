Generate pair keys
https://en.wikibooks.org/wiki/Cryptography/Generate_a_keypair_using_OpenSSL

1)openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits:2048
2)openssl rsa -pubout -in private_key.pem -out public_key.pem
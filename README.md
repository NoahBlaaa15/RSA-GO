# RSA asymmetrical encryption algorithm implemented in GO
## key generation
    find 2 prime numbers p, q
    calculate n = p * q and f(n) = (p-1)(q-1)
    find e between 1 and f(n), while e and f(n) share only 1 as their greates common devisor
    find d, so that (e * d) mod f(n) = 1
## encryption
    encrypted message = message^e mod n
## decryption
    message = encrypted message^d mod n
## digital signature
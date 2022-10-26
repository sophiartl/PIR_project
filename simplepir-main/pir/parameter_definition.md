# Parameter definition in the simple and double PIR protocol 

We define a configuration accroding to the database we are using to simulate the client.

## Parameters definition

### Database size N
Database size, also expressed as N = $2^n$



### Lattice dimension n
LWE secret dimension


### Plaintext modulus p 
We compute the largest plaintext modulus p, such that the chosen LWE parameters support at least $\sqrt{N}$ homomorphic additions, with correctness error probability 𝛿.



### Ciphertext modulus q
We first fix the ciphertext modulus 𝑞 to be one of ${2^{16}, 2^{32}, 2^{64}}$, as modern hardware natively supports operations over Z𝑞 with these moduli.


### Standard deviation $\sigma$ sigma 
The error distribution  follows a discrete gaussian with fixed stadard deviation at 6.4 

### l x m
DB height x DB width


# Parameter definition in the spiral PIR protocol 

We define a configuration accroding to the databse we are using to simulate the client. In the paper, we chose to 

## Parameters definition

### Plaintext dimension n
Let $\lambda$ be a security parameter, the size of the plaintext is denoted by n = n($\lambda$) 

### Dimensions $\nu_1$ and $\nu_2$ 

The number of records in the database is N = $2^{\nu_1+\nu_2}$
The database is represented as a  $(\nu_2+1)$-dimensional hypercube with dimensions $2^{\nu_1}×2×2×···×2$. We index elements of the database using either the tuple (i, $𝑗_1$,...,$𝑗_{\nu_2}$) where i ∈ [0, $2^{\nu_1}$ − 1] and
$𝑗_1$,...,$𝑗_{\nu_2}$ ∈ {0,1}, or the tuple (𝑖,𝑗) where 𝑖 ∈ [0,$2^{\nu_1}$ − 1] and j ∈ [0, $2^{\nu_2}$ − 1]


### Plaintext modulus p 
Let $\lambda$ be a security parameter, the plaintext moulus is denoted by p = p($\lambda$) 


### s_e



### q2_bits 
Let 𝑞 = 𝑞(𝜆) be an encoding modulus for the query and 𝑞1 = 𝑞1(𝜆), 𝑞2 = 𝑞2(𝜆) be the smaller moduli associated with the PIR response. We require that 𝑞 is odd.
This can be used for modulus switching to achieve further compression by scaling some components to q1 and remaining components to modulus q2. We can use vers small values of q1 e.g. q1 = 4p. We need a different recovering algorithm that is parametrised by a pair of moduli q1 and q2 ∈ N

### t_gsw
Converting Regev Encodings into GSW Encodings with decomposition base $z_{GSW}$ from a collection of scalar Regev encodings of $\mu g_{𝑧_{GSW}}$ = [𝜇, 𝜇 · 𝑧_{GSW}, . . . , 𝜇 · $𝑧_{GSW}^{t_{GSW}-1}$ ] with $t_{GSW}$ = $\lfloor{log_{𝑧_{GSW}}(q)}\rfloor + 1$

$z_{GSW}$ is the decomposition base used in GSW encodings
### t_conv
For control over noise growth, additional decomposition base $z_{conv}$, this parameter does not need to match the decomposition base $𝑧_{conv}$. We have as above $t_{conv} = \lfloor{log_{𝑧_{conv}}(q)}\rfloor + 1$

 $𝑧_{conv}$ is the decomposition base used to translate scalar Regev encodings into matrix Regev encodings

### t_exp_left
Coefficient expansion algortihm 

### t_exp_right
Coefficient expansion algortihm 

### instances

### db_item_size
Size of each item in the database

## Default parameters




## Our parameter selection 

 {
        "n": 2,
        "nu_1": 6,
        "nu_2": 4,
        "p": 4,
        's_e': 87.62938774292914,
        "q2_bits": 13,
        "t_gsw": 4,
        "t_conv": 4,
        "t_exp_left": 4,
        "t_exp_right": 56,
        "instances": 1,
        "db_item_size": 2048
    }


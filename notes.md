# The size of Poker

http://arxiv.org/abs/1302.7008v2

Measuring the Size of Large No-Limit Poker Games, Johanson, 2013.

### 2007-2008: $1-$2 with $1000 (500-blind) stacks

* GS: 7159379256300503000014733539416250494206634292391071646899171132778113414200 ~ 7e75
* IS: 7231696218395692677395045408177846358424267196938605536692771479904913016 ~ 7e72
* IA: 937575457443070937268150407671117224976700640913137221641272121424098561 ~ 9e71

### 2009: $1-$2 with $400 (200-blind) stacks

* GS: 1375203442350500983963565602824903351778252845259200 ~ 1e51
* IS: 1389094358906842392181537788403345780331801813952 ~ 1e48
* IA: 180091019297791288982204479657796281550065385037 ~ 1e47

### 2010-Present: $50-$100 with $20000 (200-blind) stacks

* GS: 631143875439997536762421500982349491523134755009560867161754754138543071866492234040692467854187671526019435023155654264055463548134458792123919483147215176128484600 ~ 6e164
* IS: 637519066101007550690301496238244324920475418719042634144396116764136550474559674075887513367166011522983983431697050644965107911879207553424525286198175080441144 ~ 6e161
* IA: 82653117189901827068203416669319641326155549963289335994852924537125934134924844970514122385645557438192782454335992412716935898684703899327697523295834972572001 ~ 8e160



### Refs

- Poker Game States $20,000 stacks = 6e164
- GO number of legal possitions = |{black,white,empty}|^(19*19) = 3^361 = 1e172

- 2^512 ~ 1e154
- 2^600 ~ 4e180 -> (deberíamos poder contar hasta 600 ceros 000000..00 seguidos)
- 2^1024 ~ 1e308



### impl

como partimos el hash en uint64 (cada ui64 son 8 bytes y cada byte son 8 bits) 
entonces la cantidad total de bits en el hash debe ser multiplo de 8*8 = 64

entonces las opciones de longitud del hash final son:

`[(8*8)*i for i in range(20)]`

`[0, 64, 128, 192, 256, 320, 384, 448, 512, 576, 640, 704, 768, 832, 896, 960, 1024, 1088, 1152, 1216]`



### Hashes

- ui64 ~ 2^64 ~ 1e19
- sha512 ~ 2^512 ~ 1e154
- Skein-1024 ~ 2^1024 ~ 1e308



# Duplicate hash output length

```go
import (
    "crypto/sha512"
    "encoding/hex"
)

func main() {
    // Concatenate two 512-bit hashes together
    h1 := sha512.Sum512([]byte("hello"))
    h2 := sha512.Sum512([]byte("world"))
    h := append(h1[:], h2[:]...)

    // Convert the hash to a hexadecimal string
    hashString := hex.EncodeToString(h)

    fmt.Println(hashString)
}
```



# Hash function with variable output length

```go
import (
    "crypto/sha3" // or "golang.org/x/crypto/sha3"
    "fmt"
)

func main() {
    data := []byte("your data here")
    // 1 = 1 byte / 8 bits
    hash := make([]byte, 75) // 600 bits / (8b/1B) = 600/8 B = 75 B:bytes
    sha3.ShakeSum256(hash, data)
    fmt.Printf("%x\n", hash)
}
```



# Hash output as string-bin in Python

```py
bin(int(hashlib.sha256(str("123").encode('utf-8')).hexdigest(), 16))
```



# Sobre axiom

- Los register usan `type reg uint8`.
- El máximo valor de un `uint8` es 255.
- 8e160 < 1e255 por lo que estamos bien.

- cuando se usa 16, se usa `2^15 = 32,768` (32,767 + 1) buckets de `ui8` (0..255)
- `1,548,954,453` <- max valor


# recordar

- `byte // alias for uint8`


# axiom vs clarkduvall

```
➜  hll-truco git:(main) ✗ go run cmd/count-infosets-ronda-hll-axiom/main.go -deck=7 -abs=null -report=5 -limit=60
----------------- initing
{"time":"2024-05-12T17:48:33.624888-03:00","level":"INFO","msg":"START","deckSize":7,"absId":"null","infoset":"InfosetRondaBase","hash":"sha160","limitFlag":60,"reportFlag":5}
{"time":"2024-05-12T17:48:38.62556-03:00","level":"INFO","msg":"REPORT","delta":5.000486041,"estimate":4579}
{"time":"2024-05-12T17:48:43.626034-03:00","level":"INFO","msg":"REPORT","delta":10.001007,"estimate":4647}
{"time":"2024-05-12T17:48:48.626473-03:00","level":"INFO","msg":"REPORT","delta":15.001494083,"estimate":4663}
{"time":"2024-05-12T17:48:53.626925-03:00","level":"INFO","msg":"REPORT","delta":20.001993541,"estimate":4666}
{"time":"2024-05-12T17:48:58.627364-03:00","level":"INFO","msg":"REPORT","delta":25.002481,"estimate":4670}
{"time":"2024-05-12T17:49:03.62781-03:00","level":"INFO","msg":"REPORT","delta":30.002973333,"estimate":4672}
{"time":"2024-05-12T17:49:08.628252-03:00","level":"INFO","msg":"REPORT","delta":35.003463625,"estimate":4676}
{"time":"2024-05-12T17:49:13.628701-03:00","level":"INFO","msg":"REPORT","delta":40.0039595,"estimate":4677}
{"time":"2024-05-12T17:49:18.62915-03:00","level":"INFO","msg":"REPORT","delta":45.004457,"estimate":4677}
{"time":"2024-05-12T17:49:23.629603-03:00","level":"INFO","msg":"REPORT","delta":50.004954541,"estimate":4677}
{"time":"2024-05-12T17:49:28.630046-03:00","level":"INFO","msg":"REPORT","delta":55.005449625,"estimate":4678}
{"time":"2024-05-12T17:49:33.625061-03:00","level":"INFO","msg":"RESULTS","finalEstimate":4678,"terminals:":1986229,"finished":60.000511125}
➜  hll-truco git:(main) ✗ 
➜  hll-truco git:(main) ✗ 
➜  hll-truco git:(main) ✗ go run cmd/count-infosets-ronda-hll-clarkduvall/main.go -deck=7 -abs=null -report=5 -limit=60 
{"time":"2024-05-12T17:49:46.60378-03:00","level":"INFO","msg":"START","deckSize":7,"absId":"null","infoset":"InfosetRondaBase","hash":"sha160","limitFlag":60,"reportFlag":5}
{"time":"2024-05-12T17:49:51.604214-03:00","level":"INFO","msg":"REPORT","delta":5.000145791,"estimate":4557}
{"time":"2024-05-12T17:49:56.604312-03:00","level":"INFO","msg":"REPORT","delta":10.00028775,"estimate":4614}
{"time":"2024-05-12T17:50:01.604412-03:00","level":"INFO","msg":"REPORT","delta":15.000431791,"estimate":4638}
{"time":"2024-05-12T17:50:06.604508-03:00","level":"INFO","msg":"REPORT","delta":20.000571958,"estimate":4645}
{"time":"2024-05-12T17:50:11.604605-03:00","level":"INFO","msg":"REPORT","delta":25.000715583,"estimate":4652}
{"time":"2024-05-12T17:50:16.604703-03:00","level":"INFO","msg":"REPORT","delta":30.00086025,"estimate":4654}
{"time":"2024-05-12T17:50:21.604801-03:00","level":"INFO","msg":"REPORT","delta":35.001003833,"estimate":4657}
{"time":"2024-05-12T17:50:26.60489-03:00","level":"INFO","msg":"REPORT","delta":40.00114,"estimate":4657}
{"time":"2024-05-12T17:50:31.604999-03:00","level":"INFO","msg":"REPORT","delta":45.001296791,"estimate":4657}
{"time":"2024-05-12T17:50:36.605095-03:00","level":"INFO","msg":"REPORT","delta":50.001439166,"estimate":4657}
{"time":"2024-05-12T17:50:41.605186-03:00","level":"INFO","msg":"REPORT","delta":55.001580208,"estimate":4657}
{"time":"2024-05-12T17:50:46.603698-03:00","level":"INFO","msg":"RESULTS","finalEstimate":4658,"terminals:":1971671,"finished":60.000136125}
➜  hll-truco git:(main) ✗ 
➜  hll-truco git:(main) ✗ 
➜  hll-truco git:(main) ✗ go run cmd/count-infosets-ronda-hll-clarkduvall/main.go -deck=7 -abs=null -report=5 -limit=60
{"time":"2024-05-12T19:25:13.894195-03:00","level":"INFO","msg":"START","deckSize":7,"absId":"null","infoset":"InfosetRondaBase","hash":"sha160","limitFlag":60,"reportFlag":5}
{"time":"2024-05-12T19:25:18.894636-03:00","level":"INFO","msg":"REPORT","delta":5.000122709,"estimate":4591}
{"time":"2024-05-12T19:25:23.894757-03:00","level":"INFO","msg":"REPORT","delta":10.000278875,"estimate":4646}
{"time":"2024-05-12T19:25:28.894996-03:00","level":"INFO","msg":"REPORT","delta":15.0005525,"estimate":4664}
{"time":"2024-05-12T19:25:33.895204-03:00","level":"INFO","msg":"REPORT","delta":20.000784334,"estimate":4678}
{"time":"2024-05-12T19:25:38.895349-03:00","level":"INFO","msg":"REPORT","delta":25.000960542,"estimate":4682}
{"time":"2024-05-12T19:25:43.895422-03:00","level":"INFO","msg":"REPORT","delta":30.001082125,"estimate":4684}
{"time":"2024-05-12T19:25:48.89556-03:00","level":"INFO","msg":"REPORT","delta":35.001255084,"estimate":4688}
{"time":"2024-05-12T19:25:53.895669-03:00","level":"INFO","msg":"REPORT","delta":40.001398,"estimate":4688}
{"time":"2024-05-12T19:25:58.895782-03:00","level":"INFO","msg":"REPORT","delta":45.001545709,"estimate":4689}
{"time":"2024-05-12T19:26:03.895857-03:00","level":"INFO","msg":"REPORT","delta":50.001655084,"estimate":4689}
{"time":"2024-05-12T19:26:08.895935-03:00","level":"INFO","msg":"REPORT","delta":55.001768709,"estimate":4689}
{"time":"2024-05-12T19:26:13.894317-03:00","level":"INFO","msg":"RESULTS","finalEstimate":4689,"terminals:":1934101,"finished":60.000183625}
```

# configs

jugadores: 2p, 4p, 6p
envido: 2, -1
pts: 20, 40

3 * 2 * 2 = 12 configs en total

{p:{2,4,6}, e:2, pts:{20,40}} = 3*1*2= 6 casos
{p:{2,4,6}, e:-1, pts:{20,40}} = 3*1*2= 6 casos

# road

- track number of nodes visited
- por ronda y por partida ?

## core-hours

- 9500 core-hours per day ~ 395 cores per day

- 265 cores ~ 6144 core-hours per day
### Run as

`go run cmd/ronda-infosets-count/*.go -deck=7 -abs=a3 -info=InfosetRondaBase -hash=sha160 -report=10 -track=true`

## Results

### -deck=7 -info=InfosetRondaBase -hash=sha160

| abstraction | terminals | infosets | finished |
|-------------|-----------|----------|----------|
| a1          | 688,936   | 818      | 1m14     |
| b           | 688,936   | 1,849    | 1m11     |
| a2          | 688,936   | 1,849    | 1m12     |
| a3          | 688,936   | 4,006    | 1m12     |
| none        | 688,936   | 4,690    | 1m12     |

### -deck=14 -info=InfosetRondaBase -hash=sha160

| abstraction | terminals     | infosets | finished |
|-------------|---------------|----------|----------|
| a1          | 2,446,543,800 | 3,388    | 71h12m   |
| b           | ?             | ?        | ?        |
| a2          | 2,446,543,800 | 35,448   | 76h7m    |
| a3          | 2,446,543,800 | 184,316  | 73h30m   |
| none        | ?             | ?        | ?        |
# tagli
tagli is cli tool that handles git tags.


## list
```bash
❯ tagli list --json | jq .                                                 
{
  "v0.0.1": "bf29f494fd40a69153460cd0572cc72532fb4a5b",
  "v0.0.2": "ebde28d0196866480908670970b49074b5050612",
  "v0.0.3": "22552524b04b1c1c4606e1bdb8e1fa39525db1d0",
  "v0.0.4": "d997070dd22e6734b4ca6b2d59548795d35119dc"
}
```

## increment
```bash
tagli increment -t v0.0.6 -v major
```
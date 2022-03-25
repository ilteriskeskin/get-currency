# get-currency
Get currency in Terminal

## Installation

```Bash
git clone https://github.com/ilteriskeskin/get-currency.git
cd get-currency
go build main.go
echo "alias getcurrency='/home/ilteriskeskin/Belgeler/get-currency/main'" >> ~/.bashrc # or ~/.zshrc
getcurrency
```

output:
```
{
	"usd": {
		"code": "USD",
		"name": "U.S. Dollar",
		"inverseRate": 14.834884336573
	},
	"eur": {
		"code": "EUR",
		"name": "Euro",
		"inverseRate": 16.333187498828
	},
	"gbp": {
		"code": "GBP",
		"name": "U.K. Pound Sterling",
		"inverseRate": 19.553679515153
	}
}
```

---

## Screenshot

![screenshot](/screenshot/getcurrency.png)

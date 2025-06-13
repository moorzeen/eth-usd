# ETH-USD Balance Checker

Сервис для проверки баланса токенов на заданном Ethereum адресе с конвертацией в USD. Использует только on-chain данные через Chainlink price feeds.

## Поддерживаемые токены

Сервис проверяет баланс следующих токенов:
- ETH (Ethereum)
- WETH (Wrapped Ethereum)
- USDC (USD Coin)
- USDT (Tether)
- DAI (Dai Stablecoin)
- LINK (Chainlink)
- WBTC (Wrapped Bitcoin)
- AAVE (Aave)
- UNI (Uniswap)
- SNX (Synthetix)

## Требования

- Go 1.16 или выше

## Установка

1. Клонируйте репозиторий:
```bash
git clone https://github.com/yourusername/eth-usd.git
cd eth-usd
```

2. Установите зависимости:
```bash
go mod download
```

3. Соберите приложение:
```bash
go build -o ethusd cmd/ethusd/main.go
```

## Использование

Запустите приложение:
```bash
./ethusd -address 0x123...abc
```

Где `0x123...abc` - это Ethereum адрес, баланс которого вы хотите проверить.

## Функциональность

- Проверка баланса ETH и WETH
- Проверка баланса основных ERC20 токенов
- Конвертация всех балансов в USD через Chainlink price feeds (on-chain)
- Вывод общей стоимости в USD
- Показ только токенов с ненулевым балансом

## Технические детали

Сервис использует:
- Chainlink price feeds для получения актуальных цен в USD
- ERC20 ABI для работы с токенами
- Публичную Ethereum ноду для доступа к блокчейну
- On-chain данные без использования внешних API

### Пример вывода
```
Address: 0x123...abc

Token Balances:
ETH: 1.5 (≈ $3000.00)
WETH: 0.5 (≈ $1000.00)
USDC: 1000.0 (≈ $1000.00)
LINK: 100.0 (≈ $1500.00)
WBTC: 0.1 (≈ $4000.00)

Total USD Value: $10500.00
```

## Структура проекта

```
.
├── cmd
│   └── ethusd
│       └── main.go
├── balance
│   ├── erc20.go
│   └── token.go
├── price
│   ├── abi.go
│   ├── chainlink.go
│   └── tokens.go
├── go.mod
├── go.sum
└── README.md
```

## Лицензия

MIT 
# ETH-USD Balance Checker

Сервис для проверки баланса ETH и WETH на заданном Ethereum адресе с конвертацией в USD. Использует только on-chain данные через Chainlink price feeds.

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

- Проверка баланса ETH
- Проверка баланса WETH
- Конвертация балансов в USD через Chainlink price feed (on-chain)
- Вывод общей стоимости в USD

## Технические детали

Сервис использует:
- Chainlink ETH/USD price feed (0x5f4eC3Df9cbd43714FE2740f5E3616155c5b8419)
- Публичную Ethereum ноду для доступа к блокчейну
- ERC20 ABI для работы с WETH токеном

## Структура проекта

```
.
├── cmd
│   └── ethusd
│       └── main.go
├── pkg
│   ├── balance
│   │   ├── erc20.go
│   │   └── token.go
│   └── price
│       ├── abi.go
│       └── chainlink.go
├── go.mod
├── go.sum
└── README.md
```

## Лицензия

MIT 
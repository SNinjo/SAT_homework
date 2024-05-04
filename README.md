# multi-threading

## Python &middot; ![python version](https://img.shields.io/badge/Python-3.12.0-blue) ![code coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen)

### Build

進入資料夾

```shell
cd python
```

安裝所需套件

```shell
pip install -r requirements.txt
```

### Run

```shell
python main.py
```

在 `main.py` 可以自定義員工數量與肉品

Ex. 員工 Jo、Alan 一起處理兩片牛肉

```python
main(['Jo', 'Alan'], [Beef(), Beef()])
```

### Test

```shell
pytest --cov=. --cov-report term-missing
```

## Go &middot; ![go version](https://img.shields.io/badge/Python-1.22.2-blue) ![code coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen)

### Build

```shell
cd go
```

### Run

```shell
go run main.go
```

### Test

```shell
go test -cover ./...
```

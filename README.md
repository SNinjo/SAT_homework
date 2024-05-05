# multi-threading

- [multi-threading](#multi-threading)
  - [情境模擬](#情境模擬)
  - [Python](#python)
    - [Build](#build)
    - [Run](#run)
    - [Test](#test)
  - [Go](#go)
    - [Build](#build-1)
    - [Run](#run-1)
    - [Test](#test-1)

## 情境模擬

一家肉品加工廠有三種牛肉 (牛、豬、雞)，每個員工擁有不同的代號，且處理肉品的速度如下

1. 牛肉: 1s
2. 豬肉: 2s
3. 雞肉: 3s

每個員工互不干涉且隨機取肉品，並在「取得肉品」與「處理完肉品」時要輸出 Log  
範例:

```shell
A 在 2024-01-01 00:00:00 取得牛肉
B 在 2024-01-01 00:00:00 取得豬肉
A 在 2024-01-01 00:00:01 處理完牛肉
B 在 2024-01-01 00:00:02 處理完豬肉
...
```

## Python
![python version](https://img.shields.io/badge/Python-3.12.0-blue) ![code coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen)

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
...

if __name__ == '__main__':
    main(['Jo', 'Alan'], [Beef(), Beef()])
```

### Test

```shell
pytest --cov=. --cov-report term-missing
```

## Go
![go version](https://img.shields.io/badge/Go-1.22.2-blue) ![code coverage](https://img.shields.io/badge/Coverage-100%25-brightgreen)

### Build

進入資料夾

```shell
cd go
```

安裝所需套件

```shell
go mod download
```

### Run

```shell
go run main.go
```

在 `main.go` 可以自定義員工數量與肉品

Ex. 員工 Jo、Alan 一起處理兩片牛肉

```go
...

func main() {
    employees := employee.Employees{
        All: []employee.Employee{{Id: "Jo"}, {Id: "Alan"}},
    }
    employees.Process([]meat.Meat{meat.Beef{}, meat.Beef{}})
}
```

### Test

```shell
go test -coverprofile cover.out ./...
go tool cover -html cover.out
```

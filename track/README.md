# Track data API 

Get data for scandinavian trot racing tracks

## Usage
```bash
# fetch the CSV trackdata file needed by the service
./cmd/gendata.sh

# start service 
go run cmd/main.go

# call service 
curl localhost:8080?code=6
{"location":{"latitude":6,"longitude":12.002267},"name":"Åby"}
```

## API
Takes GET parameters:
  - `code`: ATG track code

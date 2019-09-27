# Name Rhymer

Silly rhymes containing a requested name

## Usage

```bash
# generate the TXT wordlist file needed by the service
./cmd/gendata.sh

# start service 
go run cmd/main.go

# call service 
curl localhost:8080?name=buffallo+bill
buffallo bill iakttar lipsillb

curl localhost:8080?name=buffallo+bill
buffallo bill marknadsför pastill
```

## API

GET parameters:
  - `name`: name to get a rhyme for



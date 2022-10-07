echo "Shikimori is growing from pixels"
cd ../ && go mod tidy && go mod vendor &&  gofmt ./ && golangci-lint run && cd scripts ;

echo "Checking dependencies and formatting..."
go build -o bin/ ./../cmd/shikimori/;

echo "Shikimori is ready to service you!";
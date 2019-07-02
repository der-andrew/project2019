## Thesaurus
[Project documentation](http://wiki.diacare-soft.ru/bin/view/%D0%9F%D1%80%D0%BE%D0%B5%D0%BA%D1%82%D1%8B/Maximus%3A%20next/Thesaurus/)

## Run
Run this command in project root:
```commandLine
docker-compose up -d
```

## Testing

Firstly up mongo:
```commandLine
$ docker-compose up -d mongodb
```

### By `go test`
To run all tests recursively, run this in project root:
```commandLine
$ docker-compose run --rm --name thesaurus_test thesaurus sh -c 'CGO_ENABLED=0 go test -cover -covermode atomic ./...'
```


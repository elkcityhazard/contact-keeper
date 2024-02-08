DSN="<your-database-source-name>"

build:
    go build -ldflags="-X main.DSN=${DSN}" -o bin/contact-keeper cmd/api/*.go

clean:
    if [ -f bin/contact-keeper ]; then rm bin/contact-keeper; fi

run:
    ./bin/contact-keeper

start:
    @echo "Starting contact keeper..."
    @${MAKE} clean
    @${MAKE} build
    @${MAKE} run
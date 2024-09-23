# <SERVICE_NAME> REST service

Made during DEV Challenge <YAER> competition (backend nomination).

## Test

The tests are executed during container build as a part of the `Dockerfile`.

To execute tests separately run:
```
docker build --target test --progress plain --no-cache .
```

Also you may execute tests locally by running:

```
go test -v ./...
```

## Run

To start application simply run:

```
docker compose up --build
```

## REST operations

```
# Get the README
curl -X POST localhost:8080/

# Get the API root
curl -X POST localhost:8080/v1/api

## Corner cases

TBD

# Copyright

Copyright (C) 2023-2024 Serhii Zasenko

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.

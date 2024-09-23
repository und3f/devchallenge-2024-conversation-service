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

### /category ENDPOINT

Categories represent the topics of conversation. Each conversation may cover
multiple topics simultaneously. The conversation topics must be assigned
correctly, as specialists will evaluate and assess the quality of the calls
based on these topics from their respective fields.

When adding or updating a category, it is necessary to check if the conversations still belong to this category.

GET /category – Returns a list of all conversation topics.

POST /category – Creates a new conversation topic.

PUT /category/{category_id} – Updates an existing conversation topic.

DELETE /category/{category_id} – Deletes a conversation topic by the specified identifier.

Validation Rules:

    title is required for POST, optional for PUT.
    points must be an array of strings if provided.

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

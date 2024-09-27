# Conversation REST service

Made during DEV Challenge 2024 competition (backend nomination).

## Run

To start application simply run:

```
docker compose up
```

## Tests

BDD tests by feature:
  * [Category](./tests/features/category.feature)
  * [Call](./tests/features/call.feature)

The fasts tests are executed on docker compose start.

To execute all integration tests (including long-running tests, requires
Internet access) run:
```
docker compose up test
```

Fast tests that do not require audio processing could be run with:
```
docker compose up test-fast
```

Also you may execute tests locally by running:

```
npm run test --prefix tests
```

## Corner cases

### CPU only

Since I don't have GPU all LLM model are executed on CPU only.

### System requirements

CPU: 3.6 GHz, 8 Threads
RAM: 16Gb 

### Concurrent requests

Due to the limitation of the developer laptop only one concurent analyze is possible.

### Audio file limitation

Max allowed file size is 2.6MB.

## REST operations

### /api/category ENDPOINT

Categories represent the topics of conversation. Each conversation may cover
multiple topics simultaneously. The conversation topics must be assigned
correctly, as specialists will evaluate and assess the quality of the calls
based on these topics from their respective fields.

When adding or updating a category, it is necessary to check if the conversations still belong to this category.

GET /api/category -- Returns a list of all conversation topics.

POST /api/category -- Creates a new conversation topic.

PUT /api/category/{category_id} -- Updates an existing conversation topic.

DELETE /api/category/{category_id} -- Deletes a conversation topic by the specified identifier.

Validation Rules:
- title is required for POST, optional for PUT.
- points must be an array of strings if provided.

### /api/call ENDPOINT

This API provides functionality for processing and analyzing audio calls. It
allows users to submit audio files via a URL, where the service will handle the
download, transcription, and analysis of the content. The system extracts key
information such as the call name and location, determines the emotional
tone of the conversation, and stores the results.

POST /api/call -- Creates a new call based on the provided audio file URL. Supported file formats are wav and mp3.

GET /api/call/{id} -- Retrieves details of a call by the specified identifier.
The emotional tone must be one of the following values: Neutral, Positive,
Negative, Angry.

import { fullURL } from './configuration.js'

import { When, Then, wrapPromiseWithTimeout } from '@cucumber/cucumber'
import assert from 'assert'

const callRerequestInterval = 1 * 1000
const callRerequestLongInterval = 17 * 1000

const sampleCreateCall = {
  'audio_url': 'https://github.com/ggerganov/whisper.cpp/raw/refs/heads/master/samples/jfk.wav'
}

let createResponse, getResponse
let createdCallId, processedCallId = 1

When('I make a request to create a call', async function (datatable) {
  createResponse = await fetch(
    fullURL('/api/call'), {
      'method': 'POST',
      'body': JSON.stringify(datatable.rowsHash())
    }
  )
});

Then('I should receive call created success response', async function () {
  assert.equal(200, createResponse.status)

  createdCallId = (await createResponse.json()).id
  assert.ok(createdCallId != null, 'Got created call id')
});

Then('get call should return success response', async function () {
  const response = await fetch(fullURL(`/api/call/${createdCallId}`))
  assert.equal(200, response.status)

  const call = await response.json()
});

const checkCallProcessed = async function (callId, interval) {
  await new Promise(r => {
    const i = setInterval(async () => {
      const response = await fetch(fullURL(`/api/call/${callId}`))
      if (response.status != 202) {
        clearInterval(i)
        r()
      }
    }, interval)
  })
}

Then('I wait till the call is processed', {timeout: -1}, async function() {
  await checkCallProcessed(createdCallId, callRerequestInterval)
})

Then('I wait till the call is processed using long poll', {timeout: -1}, async function() {
  await checkCallProcessed(createdCallId, callRerequestLongInterval)
})

Then('get call should return unprocessable entity', async function() {
  const response = await fetch(fullURL(`/api/call/${createdCallId}`))
  assert.equal(422, response.status)
})

Then('get call should return accepted response', async function () {
  const response = await fetch(fullURL(`/api/call/${createdCallId}`))
  assert.equal(202, response.status)
});

When('I make a request to get sample processed call', async function () {
  getResponse = await fetch(fullURL(`/api/call/${processedCallId}`))
});

Then('I should receive call processed response', async function () {
  assert.equal(200, getResponse.status)
  const call = await getResponse.json()

  assert.deepEqual({
    "id": 1,
    "name": "Sample Call",
    "text": "Hello and welcome to out call in Kyiv. I am happy to talk about visa and diplomatic inquries!",
    "location": "Kyiv",
    "emotional_tone": "Neutral",
    "categories": ['Diplomatic Inquiries', 'Visa and Passport Services']
  }, call)
});

When('I make a request to get non-existing call id', async function () {
  getResponse = await fetch(fullURL(`/api/call/${processedCallId + 100}`))
});

Then('I should receive not found error', function () {
  assert.equal(404, getResponse.status)
});

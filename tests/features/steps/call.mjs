import { fullURL } from './configuration.js'

import { When, Then } from '@cucumber/cucumber'
import assert from 'assert'

const PROCESSING_TIME = 5 * 1000
const sampleCreateCall = {
  'audio_url': 'http://example.com/audiofile.wav'
}

let createResponse, getResponse
let createdCallId, processedCallId = 1

When('I make a request to create sample call', async function () {
  createResponse = await fetch(
    fullURL('/api/call'), {
      'method': 'POST',
      'body': JSON.stringify(sampleCreateCall)
    }
  )
});

Then('I should receive call created success response', async function () {
  assert.equal(200, createResponse.status)

  createdCallId = (await createResponse.json()).id
  assert.ok(createdCallId != null, 'Got created call id')
});

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
    "text": "TRANSCRIBED TEXT",
    "location": "Kyiv",
    "emotional_tone": "Neutral",
    "categories": ['Diplomatic Inquiries', 'Visa and Passport Services']
  }, call)
});

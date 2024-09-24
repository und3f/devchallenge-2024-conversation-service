import { fullURL } from './configuration.js'

import { When, Then } from '@cucumber/cucumber'
import assert from 'assert'

const PROCESSING_TIME = 5 * 1000
const sampleCreateCall = {
  'audio_url': 'http://example.com/audiofile.wav'
}

let createResponse
let createdCallId

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

When('I make a request to get created sample call', async function () {
  const getResponse = await fetch(fullURL(`/api/call/${createdCallId}`))
  assert.equal(202, getResponse.status)
});

const sleep = ms => new Promise(res => setTimeout(res, ms));

Then('I wait for call process', async function () {
  await sleep(PROCESSING_TIME)
})

Then('I should receive call processed response', async function () {
  const getResponse = await fetch(fullURL(`/api/call/${createdCallId}`))
  assert.equal(200, getResponse.status)
});

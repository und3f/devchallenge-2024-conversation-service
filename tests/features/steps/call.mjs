import { fullURL } from './configuration.js'

import { Given, When, Then, wrapPromiseWithTimeout } from '@cucumber/cucumber'
import assert from 'assert'

const callRerequestInterval = 1 * 1000
const callRerequestLongInterval = 17 * 1000

let createResponse, getResponse
let createdCallId, processedCallId

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
  assert.ok(createdCallId, 'Got created call id')
});

Given('I created call with id {int}', function(id) {
  createdCallId = id
})

Then('get call should return success response:', async function (datatable) {
  const expectedCall = ConvertDatatableToCall(datatable)

  assert.ok(createdCallId)
  const response = await fetch(fullURL(`/api/call/${createdCallId}`))
  assert.equal(200, response.status)

  const call = await response.json()
  for (const category in expectedCall) {
    assert.deepEqual(call[category], expectedCall[category])
  }
});

const checkCallProcessed = async function (callId, interval) {
  assert.ok(callId)
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
  assert.ok(createdCallId)
  await checkCallProcessed(createdCallId, callRerequestInterval)
})

Then('I wait till the call is processed using long poll', {timeout: -1}, async function() {
  assert.ok(createdCallId)
  await checkCallProcessed(createdCallId, callRerequestLongInterval)
})

Then('get call should return unprocessable entity', async function() {
  assert.ok(createdCallId)
  const response = await fetch(fullURL(`/api/call/${createdCallId}`))
  assert.equal(422, response.status)
})

Then('get call should return accepted response', async function () {
  assert.ok(createdCallId)
  const response = await fetch(fullURL(`/api/call/${createdCallId}`))
  assert.equal(202, response.status)
});

When('I make a request to get a sample processed call', async function () {
  getResponse = await fetch(fullURL(`/api/call/1000`))
});

function ConvertDatatableToCall(datatable) {
  let categories = datatable.rows().filter(a => a[0] === 'categories').map(a => a[1])
  if (categories.length == 0) {
    categories = null
  }
  return {...datatable.rowsHash(), categories}
}

Then('I should receive call processed response:', async function (datatable) {
  assert.equal(200, getResponse.status)
  const call = await getResponse.json()

  assert.deepEqual(ConvertDatatableToCall(datatable), call)
});

When('I make a request to get non-existing call id', async function () {
  getResponse = await fetch(fullURL('/api/call/1'))
});

Then('I should receive not found error', function () {
  assert.equal(404, getResponse.status)
});

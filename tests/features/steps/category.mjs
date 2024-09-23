import { fullURL } from './configuration.js'

import { When, Then } from '@cucumber/cucumber'
import assert from 'assert'
import { response as getResponse } from './api-root.mjs'

const sampleCreateCategory = {
  'title': 'Topic Title',
  'points': [
    'Key Point 1',
    'Key Point 2'
  ]
}

let createResponse

function createPostHeaders() {
  const h = new Headers();
  h.append('Content-Type', 'application/json')

  return h
}

When('I make a request to create sample category', async function () {
  createResponse = await fetch(
    '/category', {
      'method': 'POST',
      'headers': createPostHeaders(),
      'body': JSON.stringify(sampleCreateCategory)
    }
  )
});

Then('I should receive category created success response', async function () {
  const response = createResponse
  assert.equal(201, response.status)

  const category = await response.json()
  assert.ok(category.id, 'Category id exists')
})

Then('I should receive list of all conversation topics', async function () {
  const response = getResponse
  assert.equal(200, response.status)

  const categories = await response.json()
  assert.ok(categories.length > 0, 'List is not empty')
})

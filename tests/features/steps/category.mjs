import { fullURL } from './configuration.js'

import { When, Then } from '@cucumber/cucumber'
import assert from 'assert'
import { response as getResponse } from './api-root.mjs'

const sampleCreateCategory = {
  "title": "Topic Title",
  "points": [
    "Key Point 1",
    "Key Point 2"
  ]
}

const sampleUpdateCategory = {
  "title": "Updated Topic Title",
  "points": [
    "Key Point 1 (new)",
    "Key Point 2"
  ]
}

let createResponse, deleteReponse
let createdCategoryId
let categories

function createPostHeaders() {
  const h = new Headers();
  h.append("Content-Type", "application/json")

  return h
}

When('I make a request to create sample category', async function () {
  createResponse = await fetch(
    fullURL('/category'), {
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
  assert.ok(category.id, "Category id exists")
  createdCategoryId = category.id
})

Then('I should receive list of all conversation topics', async function () {
  const response = getResponse
  assert.equal(200, response.status)

  categories = await response.json()
  assert.ok(categories.length > 0, 'List is not empty')
})


Then('I should see default conversation topics', async function () {
  const expectedCategories = [
    "Visa and Passport Services",
    "Diplomatic Inquiries",
    "Travel Advisories",
    "Consular Assistance",
    "Trade and Economic Cooperation"
  ]

  for (const expected of expectedCategories) {
    assert.ok(categories.find(c => c.title == expected), `Category "${expected}" found.`)
  }
});

When('I request to delete previously created sample category', async function () {
  deleteReponse = await fetch(
    fullURL(`/category/${createdCategoryId}`), {
      'method': 'DELETE',
    }
  )
});

Then('category should be unavailable', async function () {
  const response = await fetch(fullURL('/category'))
  const category = (await response.json())
    .find(c => c.id == createdCategoryId);

  assert.ok(category == null)
});

When('I request to update previously created sample category', async function () {

  const updateResponse = await fetch(
    fullURL(`/category/${createdCategoryId}`), {
      'method': 'PUT',
      'headers': createPostHeaders(),
      'body': JSON.stringify(sampleUpdateCategory)
    }
  )

  assert.equal(200, updateResponse.status)
});

Then('category should be updated', async function () {
  const response = await fetch(
    fullURL('/category')
  )

  assert.ok(response.ok)

  const category = (await response.json())
    .find(c => c.id == createdCategoryId);

  assert.ok(category, 'Category found')

  assert.deepEqual({
    ...sampleUpdateCategory,
    id: createdCategoryId,
  }, category)
});

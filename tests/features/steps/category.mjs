import { fullURL } from './configuration.js'

import { When, Then } from '@cucumber/cucumber'
import assert from 'assert'
import { response as getResponse } from './api-root.mjs'

let createResponse, deleteReponse
let createdCategoryId
let categories

function createPostHeaders() {
  const h = new Headers();
  h.append("Content-Type", "application/json")

  return h
}

function categoryDatatableToHash(datatable) {
  const points = datatable.rows().filter(a => a[0] === 'points').map(a => a[1])
  return {...datatable.rowsHash(), points}
}

When('I make a request to create a category:', async function (datatable) {
  createResponse = await fetch(
    fullURL('/api/category'), {
      'method': 'POST',
      'headers': createPostHeaders(),
      'body': JSON.stringify(categoryDatatableToHash(datatable))
    }
  )
});

Then('I should receive category created success response:', async function (datatable) {
  const response = createResponse

  assert.equal(response.status, 201)

  const category = await response.json()
  assert.ok(category.id, "Category id exists")
  createdCategoryId = category.id

  assert.deepEqual(category, {id: category.id, ...categoryDatatableToHash(datatable)})
})

Then('API returns: category create error: unprocessable entity', async function() {
  assert.equal(createResponse.status, 422)
})

Then('I should receive list of all conversation topics', async function () {
  const response = getResponse
  assert.equal(response.status, 200)

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
    const category = categories.find(c => c.title == expected)
    assert.ok(category, `Category "${expected}" found.`)
    assert.ok(category.points.length > 0, 'Category contains points.')
  }
});

When('I request to delete previously created category', async function () {
  deleteReponse = await fetch(
    fullURL(`/api/category/${createdCategoryId}`), {
      'method': 'DELETE',
    }
  )
});

Then('category should be unavailable', async function () {
  const response = await fetch(fullURL('/api/category'))
  const category = (await response.json())
    .find(c => c.id == createdCategoryId);

  assert.ok(category == null)
});

When('I request to update previously created category with:', async function (datatable) {
  const updateResponse = await fetch(
    fullURL(`/api/category/${createdCategoryId}`), {
      'method': 'PUT',
      'headers': createPostHeaders(),
      'body': JSON.stringify(categoryDatatableToHash(datatable))
    }
  )

  assert.equal(200, updateResponse.status)
});

When('I request to update previously created category using invalid data:', async function (datatable) {
  const updateResponse = await fetch(
    fullURL(`/api/category/${createdCategoryId}`), {
      'method': 'PUT',
      'headers': createPostHeaders(),
      'body': JSON.stringify(categoryDatatableToHash(datatable))
    }
  )

  assert.equal(422, updateResponse.status)
});

Then('category should match:', async function (datatable) {
  const response = await fetch(
    fullURL('/api/category')
  )

  assert.ok(response.ok)

  const category = (await response.json())
    .find(c => c.id == createdCategoryId);

  assert.ok(category, 'Category found')

  assert.deepEqual(category, {
    ...categoryDatatableToHash(datatable),
    id: createdCategoryId,
  })
});

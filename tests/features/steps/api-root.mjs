import { fullURL } from './configuration.js'

import { When, Then } from '@cucumber/cucumber'
import assert from 'assert'

export let response;
When('I request {string}', async function (path) {
  response = await fetch(fullURL(path))
})

Then('I should receive Readme document', async function () {
  assert.ok(response.ok)

  assert.match(await response.text(), /DEV Challenge/)
})

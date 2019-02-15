const _ = require('lodash')
const request = require('request-promise-native')
const constants = require('./constants')

async function getPkey(body) {
  const options = {
    uri: `https://${constants.GIMMEURL}`,
    method: 'GET',
    json: true,
    resolveWithFullResponse: true,
    time: true,
    body,
  }

  return request(options)
    .then(function(response) {
      return response.body
    })
    .catch(function(err) {
      throw err
    })
}

async function verifyIdentity(token) {
  const options = {
    uri: `https://${constants.VERIFYURL}`,
    method: 'GET',
    json: true,
    headers: {
      Authorization: `Bearer ${token}`
    },
    resolveWithFullResponse: true,
    time: true,
  }

  return request(options)
    .then(function(response) {
      return response.body
    })
    .catch(function(err) {
      throw err
    })
}

module.exports = {
  getPkey,
  verifyIdentity,
}

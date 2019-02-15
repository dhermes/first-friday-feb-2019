// import * as _ from 'lodash'
// import * as request from 'request-promise-native'
// import { GIMMEURL, VERIFYURL } from './constants';
const _ = require('lodash')
const request = require('request-promise-native')
const constants = require('./constants')
// const fromCallback = require('promise-cb').fromCallback;

async function getPkey(body) {
  const options = {
    uri: `https://${constants.GIMMEURL}`,
    method: 'GET',
    json: true,
    resolveWithFullResponse: true,
    time: true,
    body,
  }

  // tslint:disable-next-line: no-unsafe-any
  return request(options)
    .then(function(response) {
      // tslint:disable-next-line: no-unsafe-any
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

  // tslint:disable-next-line: no-unsafe-any
  return request(options)
    .then(function(response) {
      console.log('hiiiiiiiiiiiiii')
      console.log(response.body)
      // tslint:disable-next-line: no-unsafe-any
      return response.body
    })
    .catch(function(err) {
      console.log('hiii')
      console.log(err)
      throw err
    })
}

module.exports = {
  getPkey,
  verifyIdentity,
}
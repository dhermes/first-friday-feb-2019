// import * as jwt from 'jsonwebtoken'
// import * as uuid from 'uuid'
// import * as _ from 'lodash'
// import { EXPIRATION_TIME } from './constants'
const jwt = require('jsonwebtoken')
const uuid = require('uuid')
const _ = require('lodash')
const constants = require('./constants')

function generateJwtToken(pKey) {
  const jwtOptions = {
    algorithm: 'RS256',
    audience: constants.issuer,
    subject: uuid.v4(),
    issuer: constants.issuer,
    expiresIn: constants.EXPIRATION_TIME,
    keyid: uuid.v4(),
  };
  return jwt.sign({
    some: 'payload',
  }, pKey, jwtOptions)
}

module.exports = {
  generateJwtToken
}
// import { generateJwtToken } from "./jwt";
// import { getPkey, verifyIdentity } from "./lambdaExternal";
const generateJwtToken = require('./jwt').generateJwtToken
const getPkey = require('./lambdaExternal').getPkey
const verifyIdentity = require('./lambdaExternal').verifyIdentity

async function runscript () {
  const privateKey = await getPkey()
  const token = generateJwtToken(privateKey.privateKey.key)
  console.log('-----------------------------')
  console.log(token)
  console.log('-----------------------------')
  await verifyIdentity(token)
  return
}

runscript()

const crypto = require("crypto");

const arg = process.argv[2]

if (arg === 'string') {
  let length = 10

  if (
    process.argv[3] && 
    (process.argv[3] === '--length' || process.argv[3] === '-l' )
  ) {
    length = parseInt(process.argv[4], 10)
  }

  const id = crypto.randomBytes(length).toString("hex");
  console.log(id);
}
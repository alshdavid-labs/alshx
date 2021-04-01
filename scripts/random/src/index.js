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

if (arg === 'uuid' || arg === 'guid') {
  const guid = ([1e7]+-1e3+-4e3+-8e3+-1e11).replace(/[018]/g, c =>
    (c ^ crypto.randomInt(1000) & 15 >> c / 4).toString(16)
  );
  console.log(guid)
}
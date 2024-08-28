const fs = require("node:fs");
const content = "Some content again!";
try {
  fs.appendFileSync("test.txt", content);
} catch (err) {
  console.error(err);
}

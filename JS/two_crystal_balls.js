const arr = [
  false,
  false,
  false,
  false,
  false,
  false,
  false,
  false,
  true,
  true,
  true,
  true,
];

function tcb(arr) {
  const jump = Math.floor(Math.sqrt(arr.length));
  let i = jump;
  for (; i < arr.length; i += jump) {
    if (arr[i]) {
      break;
    }
  }
  i -= jump;
  for (let j = 0; j < jump && i < arr.length; j++, i++) {
    if (arr[i]) {
      return i;
    }
  }
}

console.log(tcb(arr));

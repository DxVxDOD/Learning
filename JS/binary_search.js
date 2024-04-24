function bs(arr, needle) {
  let h = arr.length;
  let l = 0;
  while (l < h) {
    let m = Math.floor(l + (h - l) / 2);
    const v = arr[m];
    if (v === needle) {
      return true;
    } else if (v > needle) {
      h = m;
    } else {
      l = m + 1;
    }
  }
  return false;
}

console.log(bs([1, 5, 7, 11, 15, 22, 30], 7));
console.log(bs([1, 5, 7, 11, 15, 22, 30], 30));
console.log(bs([1, 5, 7, 11, 15, 22, 30], 33));

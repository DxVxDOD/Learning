function bs(arr, needle) {
  let h = arr.length - 1
  let l = 0

  while (h >= l){
    let middle = Math.floor((h + l) / 2)
    if(arr[middle] === needle){
      return true
    } else if(arr[middle] > needle){
      h = middle - 1
    } else {
      l = middle + 1
    }
  }
  return false
}

console.log(bs([1,5,7,11,15,22,30], 7))
console.log(bs([1,5,7,11,15,22,30], 30))
console.log(bs([1,5,7,11,15,22,30], 33))

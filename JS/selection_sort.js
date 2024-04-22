function selection_sort(arr) {
  for (let i = 0; i < arr.length; i++) {
    let curr_low = i;
    for (let j = i + 1; j < arr.length; j++) {
      if (arr[curr_low] > arr[j]) {
        curr_low = j;
      }
    }
    if (curr_low !== i) {
      let temp = arr[i];
      arr[i] = arr[curr_low];
      arr[curr_low] = temp;
    }
  }
  return arr;
}

console.log(selection_sort([65, 55, 45, 69, 35, 25, 15, 10]));

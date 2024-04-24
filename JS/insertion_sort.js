function insertion_sort(arr) {
  for (let i = 1; i < arr.length; i++) {
    let p = i;
    const temp = arr[i];
    while (p > 0 && arr[p - 1] > temp) {
      arr[p] = arr[p - 1];
      p = p - 1;
    }
    arr[p] = temp;
  }
  return arr;
}

console.log(insertion_sort([65, 55, 2, 111, 45, 69, 35, 25, 15, 10]));

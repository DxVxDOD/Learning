function findMaxVal(subArr: number[], i: number): number {
  if (i > 0) {
    const j = findMaxVal(subArr, i - 1);
    if (subArr[i] < subArr[j]) return j;
  }
  return i;
}

export function selectionSort(arr: number[], i = -1) {
  if (i < 0) {
    i = arr.length - 1;
  }
  if (i > 0) {
    const j = findMaxVal(arr, i);
    if (i !== j) {
      arr[i] = arr[i] + arr[j];
      arr[j] = arr[i] - arr[j];
      arr[i] = arr[i] - arr[j];
    }
    selectionSort(arr, i - 1);
  }
}

function qs(arr: number[], low: number, high: number) {
  if (low < high) {
    const pivotIndex = partition(arr, low, high);
    qs(arr, low, pivotIndex);
    qs(arr, pivotIndex + 1, high);
  }
  return arr;
}

function partition(nums: number[], low: number, high: number) {
  const pivot = nums[low];
  let i = low;
  let j = high;

  while (i < j) {
    while (nums[i] <= pivot) {
      i++;
    }

    while (nums[j] > pivot) {
      j--;
    }
    if (i < j) {
      const temp = nums[i];
      nums[i] = nums[j];
      nums[j] = temp;
    }
  }
  const temp = nums[j];
  nums[j] = nums[low];
  nums[low] = temp;
  return j;
}

const arrNum = [5, 1, 1, 2, 0, 0];

console.log(qs(arrNum, 0, arrNum.length - 1));

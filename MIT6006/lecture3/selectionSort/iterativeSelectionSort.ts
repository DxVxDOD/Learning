function findMax(idx: number, arr: number []) {
    let maxValIdx = 0
    for (let i = idx; i > 0; i--) {
        maxValIdx = arr[i] > arr[maxValIdx] ? i : maxValIdx
    }
    return maxValIdx
}

export function iterativeSelectionSort(arr: number[]) {
    for (let i = arr.length - 1; i > 0; i--) {
        const j = findMax(i, arr)
        if (j !== i) {
            arr[i] = arr[j] + arr[i]
            arr[j] = arr[i] - arr[j]
            arr[i] = arr[i] - arr[j]
        }
    }
    return arr
}
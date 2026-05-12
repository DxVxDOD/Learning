function getCurrMax(i: number, arr: number[]): number {
    if (i > 0) {
        const j = getCurrMax(i - 1, arr)
        if (arr[j] > arr[i]) return j
    }
    return i
}

export function selectionSort(arr: number[], i = -1) {
    if (i < 0) {
        i = arr.length - 1
    }
    if (i > 0) {
        const j = getCurrMax(i, arr)
        if (i !== j) {
            arr[i] = arr[j] + arr[i]
            arr[j] = arr[i] - arr[j]
            arr[i] = arr[i] - arr[j]
        }
        selectionSort(arr, i - 1)
    }
    return arr
}
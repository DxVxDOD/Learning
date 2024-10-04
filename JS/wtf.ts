let arr = [40, 10, 20, 30];
let ht = new Map<number, number>();
const og_arr = structuredClone(arr);
arr.sort((a, b) => a - b);
console.log("arr", arr);
console.log("Og_arr", og_arr);
for (let i = 0; i < arr.length; i++) {
  ht.set(arr[i], i + 1);
}
console.log("hash", [...ht.entries()]);
let res: number[] = [];
for (let i = 0; i < og_arr.length; i++) {
  res[i] = ht.get(og_arr[i])!;
}
console.log("res", res);

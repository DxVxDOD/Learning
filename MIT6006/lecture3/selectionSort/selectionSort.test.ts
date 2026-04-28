import { suite, test } from "node:test";
import { selectionSort } from "./selectionSort.ts";
import assert from "node:assert";

suite("Green:", () => {
  test("Array is sorted in place", () => {
    const sorted = [1, 2, 3, 3, 3, 4, 4, 4, 5, 6, 6, 7, 8, 8];
    const toBeSorted = [1, 4, 3, 6, 8, 3, 4, 5, 6, 7, 8, 4, 3, 2];
    selectionSort(toBeSorted);
    assert.deepStrictEqual(toBeSorted, sorted);
  });

  test("Array is backward", () => {
    const sorted = [1, 2, 3, 4, 5, 6, 7, 8, 9];
    const toBeSorted = [9, 8, 7, 6, 5, 4, 3, 2, 1];
    selectionSort(toBeSorted);
    assert.deepStrictEqual(toBeSorted, sorted);
  });

  test("Array is already sorted", () => {
    const sorted = [1, 2, 3, 4, 5, 6, 7, 8, 9];
    const toBeSorted = [1, 2, 3, 4, 5, 6, 7, 8, 9];
    selectionSort(toBeSorted);
    assert.deepStrictEqual(toBeSorted, sorted);
  });
});

import { test, expect } from "vitest";
import twoSum from "./two-sum";

test("two sum", () => {
  const testCases: Array<{ nums: number[]; target: number; want: number[] }> = [
    {
      nums: [2, 7, 11, 15],
      target: 9,
      want: [0, 1],
    },
    {
      nums: [3, 2, 4],
      target: 6,
      want: [1, 2],
    },
    {
      nums: [3, 3],
      target: 6,
      want: [0, 1],
    },
  ];

  const sortCb = (a: number, b: number): number => a - b;

  for (const testCase of testCases) {
    const got = twoSum(testCase.nums, testCase.target);
    expect(got.sort(sortCb)).toEqual(testCase.want.sort(sortCb));
  }
});

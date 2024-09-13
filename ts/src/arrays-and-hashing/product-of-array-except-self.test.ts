import { test, expect } from "vitest";
import productExceptSelf from "./product-of-array-except-self";

test("product of array except self", () => {
  const testCases: Array<{ input: number[]; want: number[] }> = [
    {
      input: [1, 2, 3, 4],
      want: [24, 12, 8, 6],
    },
    {
      input: [-1, 1, 0, -3, 3],
      want: [0, 0, 9, 0, 0],
    },
  ];

  for (const testCase of testCases) {
    const got = productExceptSelf(testCase.input);
    expect(got).toEqual(testCase.want);
  }
});

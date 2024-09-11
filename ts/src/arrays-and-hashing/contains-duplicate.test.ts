import { test, expect } from "vitest";
import containsDuplicate from "./contains-duplicate";

test("contains duplicate", () => {
  const testCases: Array<{ input: number[]; want: boolean }> = [
    { input: [1, 2, 3, 1], want: true },
    { input: [1, 2, 3, 4], want: false },
  ];

  for (const testCase of testCases) {
    const got = containsDuplicate(testCase.input);
    expect(got).toBe(testCase.want);
  }
});

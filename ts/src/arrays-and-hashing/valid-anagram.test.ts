import { test, expect } from "vitest";
import isAnagram from "./valid-anagram";

test("valid anagram", () => {
  const testCases: Array<{ s: string; t: string; want: boolean }> = [
    {
      s: "anagram",
      t: "nagaram",
      want: true,
    },
    {
      s: "rat",
      t: "car",
      want: false,
    },
  ];

  for (const testCase of testCases) {
    const got = isAnagram(testCase.s, testCase.t);
    expect(got).toBe(testCase.want);
  }
});

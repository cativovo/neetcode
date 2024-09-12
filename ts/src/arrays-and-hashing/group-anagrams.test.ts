import { test, expect } from "vitest";
import groupAnagrams from "./group-anagrams";

test("group anagrams", () => {
  const testCases: Array<{ input: string[]; want: string[][] }> = [
    {
      input: ["eat", "tea", "tan", "ate", "nat", "bat"],
      want: [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]],
    },
    {
      input: [""],
      want: [[""]],
    },
    {
      input: ["a"],
      want: [["a"]],
    },
  ];

  for (const testCase of testCases) {
    const got = groupAnagrams(testCase.input);
    expect(sort(got)).toEqual(sort(testCase.want));
  }
});

function sort(input: string[][]) {
  return input.slice().map((v) => v.sort()).sort((a, b) => a.join('').localeCompare(b.join('')));
}

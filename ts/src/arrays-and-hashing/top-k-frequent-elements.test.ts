import { test, expect } from "vitest";
import topKFrequent from "./top-k-frequent-elements";

test("top k frequent elements", () => {
  const testCases: Array<{
    nums: int[];
    k: int;
    want: int[];
  }> = [
    {
      nums: [1, 1, 1, 2, 2, 3],
      k: 2,
      want: [1, 2],
    },
    {
      nums: [1],
      k: 1,
      want: [1],
    },
    {
      nums: [-1, -1],
      k: 1,
      want: [-1],
    },
  ];

  const sortCb = (a, b) => a - b;

  for (const testCase of testCases) {
    const got = topKFrequent(testCase.nums, testCase.k);
    expect(got.sort(sortCb)).toEqual(testCase.want.sort(sortCb));
  }
});

function twoSum(nums: number[], target: number): number[] {
  const map = new Map<number, number>();

  for (let i = 0; i < nums.length; i++) {
    const num = nums[i];
    const j = map.get(num);
    if (j !== undefined) {
      return [j, i];
    }

    map.set(target - num, i);
  }

  return [];
}

export default twoSum;

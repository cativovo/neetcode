function topKFrequent(nums: number[], k: number): number[] {
  const map = new Map<number, number>();

  for (const num of nums) {
    const v = map.get(num);
    if (v !== undefined) {
      map.set(num, v + 1);
    } else {
      map.set(num, 1);
    }
  }

  const a = new Array<number[]>(map.size + 2).fill([]);

  for (const [num, count] of map) {
    a[count] = [...a[count], num];
  }

  const result: Array<number> = [];

  for (let i = a.length - 1; i >= 0; i--) {
    if (result.length === k) {
      return result;
    }

    result.push(...a[i]);
  }

  return [];
}

export default topKFrequent;

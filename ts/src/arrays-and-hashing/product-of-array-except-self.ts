function productExceptSelf(nums: number[]): number[] {
  const numsLen = nums.length;
  const result: Array<number> = new Array(numsLen).fill(0);

  let prefix = 1;
  for (let i = 0; i < numsLen; i++) {
    result[i] = prefix;
    prefix *= nums[i];
  }

  let postFix = 1;
  for (let i = numsLen - 1; i >= 0; i--) {
    result[i] *= postFix;
    if (result[i] === -0) {
      result[i] = Math.abs(result[i]);
    }
    postFix *= nums[i];
  }

  return result;
}

export default productExceptSelf;

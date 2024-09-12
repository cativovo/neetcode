function groupAnagrams(strs: string[]): string[][] {
  const map = new Map<string, string[]>();

  for (const s of strs) {
    const key = getKey(s);
    const value = map.get(key);
    if (value !== undefined) {
      value.push(s);
      map.set(key, value);
    } else {
      map.set(key, [s]);
    }
  }

  const result: Array<string[]> = [];
  for (const [, value] of map) {
    result.push(value);
  }

  return result;
}

function getKey(input: string): string {
  const start = 97; // UTF-16 of 'a'
  const key = new Array(26).fill(0);

  for (const s of input) {
    key[s.charCodeAt() - start]++;
  }

  return key.join("-");
}

export default groupAnagrams;

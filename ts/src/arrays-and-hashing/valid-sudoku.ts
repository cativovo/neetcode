function isValidSudoku(board: string[][]): boolean {
  const rowMap = new Map<number, Map<string, boolean>>();
  const colMap = new Map<number, Map<string, boolean>>();
  const squareMap = new Map<number, Map<string, boolean>>();
  const EMPTY = ".";

  for (let rowIndex = 0; rowIndex < board.length; rowIndex++) {
    for (let colIndex = 0; colIndex < board[rowIndex].length; colIndex++) {
      const value = board[rowIndex][colIndex];
      if (value === EMPTY) {
        continue;
      }

      let row = rowMap.get(rowIndex);
      if (row === undefined) {
        row = new Map<string, boolean>();
        rowMap.set(rowIndex, row);
      }

      let col = colMap.get(colIndex);
      if (col === undefined) {
        col = new Map<string, boolean>();
        colMap.set(colIndex, col);
      }

      const squareKey = Math.floor(rowIndex / 3) + Math.floor(colIndex / 3) * 3;
      let square = squareMap.get(squareKey);
      if (square === undefined) {
        square = new Map<string, boolean>();
        squareMap.set(squareKey, square);
      }

      if (row.get(value) || col.get(value) || square.get(value)) {
        return false;
      }

      row.set(value, true);
      col.set(value, true);
      square.set(value, true);
    }
  }

  return true;
}

export default isValidSudoku;

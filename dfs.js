const graph = {
  A: ["B", "C"],
  B: ["A", "D"],
  C: ["A", "E"],
  D: ["B", "F"],
  E: ["C", "G"],
  F: ["D", "H", "I"],
  G: ["E", "J", "K"],
  H: ["F", "L"],
  I: ["F", "M"],
  J: ["G", "N"],
  K: ["G", "O"],
  L: ["H"],
  M: ["I", "P"],
  N: ["J"],
  O: ["K"],
  P: ["M"],
};
const bfs = (graph, start) => {
  const checked = []; // 탐색 완료 데이터
  const willCheck = []; // 탐색 예정 데이터

  willCheck.push(start);

  while (willCheck.length !== 0) {
    const node = willCheck.pop(); // 스택(Last In First Out)
    if (!checked.includes(node)) {
      checked.push(node);
      //reverse() 제거 시 그림의 4,3,2,1 순서로 탐색
      willCheck.push(...graph[node].reverse());
    }
  }

  return checked;
};
console.log(bfs(graph, "A"));

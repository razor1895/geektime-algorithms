function climbStairs(n: number): number {
  if (n < 2) {
    return n;
  }

  let a = 1;
  let b = 1;
  let sum = 0;

  for (let i = 2; i <= n; i++) {
    sum = a + b;
    a = b;
    b = sum;
  }

  return sum;
}

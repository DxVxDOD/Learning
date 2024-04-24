function fib(n) {
  if (n < 2) {
    return n;
  } else {
    return fib(n - 1) + fib(n - 2);
  }
}

console.log(fib(10));
console.log(fib(3));

function for_fib(n) {
  res = [0, 1, 1];
  for (let i = 2; i <= n; i++) {
    res[i] = res[i - 1] + res[i - 2];
  }
  return res[n];
}

console.log(for_fib(100));
console.log(fib(10));

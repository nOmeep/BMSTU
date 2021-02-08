'use strict'

const x = 20;

const memoize = (fn) => {
    let cache = {};
    return (...args) => {
        let n = args[0];
        if (n in cache)
            return cache[n];
        else {
            let result = fn(n);
            cache[n] = result;
            return result;
        }
    }
}

const factorial = (n) => {
    if (n<=1)
        return 1;
    else
    return n*factorial(n-1);
};

const memFactorial = memoize(factorial);
console.log(memFactorial(10));
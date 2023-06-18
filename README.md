# binsearch

## What is it?
It is yet another binary search algorithm implementation, but for storing strings. As this project originally was supposed to be used in the JSON-decoder, it tries to chase a performance.

## What are odds?
At the moment, it doesn't solve collisions. Also its hash-function (Jenkins' One-at-a-Time hash) is pretty fast, but on mid-length keys (15-20 bytes) with a difference of a few last symbols, gives ~100% collisions rate. 

## How can I contribute?
Just make some changes, and open a pull-request. But please, cover your new code by tests

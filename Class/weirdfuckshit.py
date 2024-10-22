
import hashlib
from functools import cache
from time import time
  
def timer(func):
    # This function shows the execution time of 
    # the function object passed
    def wrap_func(*args, **kwargs):
        t1 = time()
        result = func(*args, **kwargs)
        t2 = time()
        print(f'Function {func.__name__!r} executed in {(t2-t1):.4f}s')
        
        with open('logs.txt', 'a') as f:
            f.write(f"Function {func.__name__!r} executed in {(t2-t1):.4f}s\n")
        return result
    
    
    return wrap_func



@timer
@cache
def mine_block(difficulty = 2):
    attempt = 0
    hash_attempt = hashlib.sha256(str(attempt).encode()).hexdigest()
    while not hash_attempt.startswith('0' * difficulty):
        attempt += 1
        hash_attempt = hashlib.sha256(str(attempt).encode()).hexdigest()
        if attempt % 100000 == 0:
            print(f"Attempt {attempt}: {hash_attempt}")
    
    with open('logs.txt', 'a') as f:
        f.write(f"Attempt {attempt}: {hash_attempt}\n")
               
    return f"Block mined at attempt {attempt}: {hash_attempt}"


for i in range(5, 10):
    
    hash = mine_block(i)
    print(f"{i-1} leading zeros:  \t{hash}")

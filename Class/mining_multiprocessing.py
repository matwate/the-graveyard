from multiprocessing import Pool
import hashlib
from time import time
from functools import cache


def mine_block(start, difficulty=8):
    attempt = start
    hash_attempt = hashlib.sha256(str(attempt).encode()).hexdigest()
    while not hash_attempt.startswith('0' * difficulty):
        attempt += 1
        hash_attempt = hashlib.sha256(str(attempt).encode()).hexdigest()
        if attempt % 100000 == 0:
            print(f"Attempt {attempt}: {hash_attempt}")
    
    with open('logs.txt', 'a') as f:
        f.write(f"Attempt {attempt}: {hash_attempt}\n")
               
    return f"Block mined at attempt {attempt}: {hash_attempt}"

if __name__ == "__main__":
    with Pool() as p:
        intervals = [i * 10**13 for i in range(4,10)]  # Start attempts at 10 trillion intervals
        results = p.map(mine_block, intervals)
        for result in results:
            print(result)
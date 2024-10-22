

use rand::Rng;
use crate::transaction::Transaction;
use crate::transaction::TransactionFuncs;

#[derive(Debug)]
pub struct Block{
    pub index: i32,
    pub timestamp: i64,
    pub transactions: Vec<Transaction>,
    pub previous_hash: i32,
    pub nonce: i32,
    pub hash: i32,
}

pub trait BlockFuncs{
    fn compute_hash(&mut self) -> i32;
    fn mine(&mut self, difficulty: i32);
    fn to_json(&self) -> String;
}

impl BlockFuncs for Block{
    fn compute_hash(&mut self) -> i32{
        let mut rng = rand::thread_rng();
        let max_value = i32::MAX;
        let mut hash = rng.gen_range(0..max_value);
        self.hash = hash;
        hash
    }
    fn mine(&mut self, difficulty: i32){
        let mut rng = rand::thread_rng();
        let max_value = i32::MAX;
        let mut hash = rng.gen_range(0..max_value);
        while (hash >> (32 - difficulty)) != 0{
            self.nonce += 1;
            hash = self.compute_hash();
        }
    }
    fn to_json(&self) -> String{
        let mut transactions = String::new();
        for trx in &self.transactions{
            transactions.push_str(&trx.to_json());
        }
        format!("{{\"index\": {}, \"timestamp\": {}, \"transactions\": [{}], \"previous_hash\": {}, \"nonce\": {}, \"hash\": {}}}", self.index, self.timestamp, transactions, self.previous_hash, self.nonce, self.hash)
    }

}


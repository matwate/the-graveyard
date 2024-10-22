use crate::block::Block;
use crate::block::BlockFuncs;
use crate::transaction::Transaction;
use crate::transaction::TransactionFuncs;
use crate::transaction::User;
use std::fmt;
use rand::Rng;

pub struct Blockchain {
    pub chain: Vec<Block>,
    pub difficulty: i32,
    pub pending_transactions: Vec<Transaction>,
    pub mining_reward: i32,
    pub miners: Vec<User>,
}

pub trait BlockchainFuncs {
    fn create_genesis_block(&mut self);
    fn add_new_block(&mut self, _block: Block);
    fn select_miner(&self) -> &User;
    fn get_last_block(&self) -> &Block;
    fn to_json(&self) -> String;
}

impl BlockchainFuncs for Blockchain {
    fn create_genesis_block(&mut self) {
        let mut gen_trx = Transaction {
            sender: "0".to_string(),
            receiver: "0".to_string(),
            amount: 0,
            signature: "0".to_string(),
        };
        gen_trx.sign("0");

        let  gen_block = Block {
            index: 0,
            timestamp: 0,
            transactions: vec![gen_trx],
            previous_hash: 0,
            nonce: 0,
            hash: 0,
        };

        self.chain.push(gen_block);
    }
    fn add_new_block(&mut self, mut _block: Block) {
        let  miner = self.select_miner();
        let  miner_pub_key = &miner.public_key;
        let mut miner_reward_trx = Transaction {
            sender: "0".to_string(),
            receiver: miner_pub_key.to_string(),
            amount: self.mining_reward,
            signature: "0".to_string(),
        };
        miner_reward_trx.sign(&miner.private_key);
        _block.transactions.push(miner_reward_trx);
        _block.mine(self.difficulty);
        self.chain.push(_block);
    }
    fn select_miner(&self) -> &User {
        let mut rng = rand::thread_rng();
        let num_miners = self.miners.len();
        
        if num_miners > 0 {
            let index = rng.gen_range(0..num_miners);
            &self.miners[index]
        } else {
            panic!("No miners available");
        }
    }
    fn get_last_block(&self) -> &Block {
        let last_index = self.chain.len() - 1;
        &self.chain[last_index]
    }
    fn to_json(&self) -> String {
        let mut blocks = String::new();
        for block in &self.chain {
            blocks.push_str(&block.to_json());
        }
        format!("{{\"chain\": [{}], \"difficulty\": {}, \"pending_transactions\": [{:?}], \"mining_reward\": {}, \"miners\": [{:?}]}}", blocks, self.difficulty, self.pending_transactions, self.mining_reward, self.miners)
    }

    
}


impl fmt::Display for Blockchain {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "Chain: {:?}, Difficulty: {}, Pending Transactions: {:?}, Mining Reward: {}, Miners: {:?}", self.chain, self.difficulty, self.pending_transactions, self.mining_reward, self.miners)
    }
}

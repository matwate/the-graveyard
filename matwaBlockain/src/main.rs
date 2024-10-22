extern crate rand;

mod transaction;
mod block;
mod blockchain;
mod api;
use block::BlockFuncs;
use transaction::Transaction;
use transaction::TransactionFuncs;
use block::Block;

use api::start;



fn main() {
    start().unwrap();
    
}



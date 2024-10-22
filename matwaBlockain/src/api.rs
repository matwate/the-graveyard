use actix_web::{get, post, web, App, HttpResponse, HttpServer, Responder};
use crate::blockchain::Blockchain;
use crate::blockchain::BlockchainFuncs;
use crate::transaction::Transaction;
use crate::transaction::User;
use rand::Rng;

#[get("/")]
async fn helloworld() -> impl Responder{
    HttpResponse::Ok().body("Hello World")
}

#[get("/newuser")]
async fn new_user() -> impl Responder{

    let random_key: String = rand::thread_rng().gen_range(0..i32::MAX).to_string();  
    // Generate the public key from the private key
    let public_key: String = (random_key.parse::<i32>().unwrap().pow(2)).to_string();

    let mut user: User = User{
        name: "0".to_string(),
        balance: 100,
        private_key: random_key.clone(),
        public_key: "0".to_string(),
    };
    users.push(user); // Append the new user to the users vector
    HttpResponse::Ok().body("New user created")
}

#[actix_web::main]
pub async fn start() -> std::io::Result<()>{
    let mut blockchain: Blockchain = Blockchain{
        chain: Vec::new(),
        difficulty: 2,
        pending_transactions: Vec::new(),
        mining_reward: 100,
        miners: Vec::new(),
    };

    blockchain.create_genesis_block();

    let mut users: Vec<User> = Vec::new();

    let mut pending_trxs: Vec<Transaction> = Vec::new();

    println!("Starting server at http://localhost:8080");

    HttpServer::new(||{
        App::new()
        .service(helloworld)

    }).bind(("localhost", 8080))?
    .run() 
    .await
   
}


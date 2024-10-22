use std::fmt;

#[derive(Debug)]
pub struct User {
    pub name: String,
    pub balance: i32,
    pub private_key: String,
    pub public_key: String,
}

#[derive(Debug)]
pub struct Transaction {
    pub sender: String,
    pub receiver: String,
    pub amount: i32,
    pub signature: String,
}

pub trait TransactionFuncs {
    fn to_json(&self) -> String;
    fn sign(&mut self, _private_key: &str);
    fn verify(&self, _public_key: &str) -> bool;
    fn check_balance(&self, user: &User) -> bool;
}

impl TransactionFuncs for Transaction {
    fn to_json(&self) -> String {
        format!(
            "{{\"sender\": \"{}\", \"receiver\": \"{}\", \"amount\": {}, \"signature\": \"{}\"}}",
            self.sender, self.receiver, self.amount, self.signature
        )
    }

    fn sign(&mut self, _private_key: &str) {
        self.signature = format!("{} sent to {}", self.sender, self.receiver);
    }

    fn verify(&self, _public_key: &str) -> bool {
        self.signature == format!("{} sent to {}", self.sender, self.receiver)
    }

    fn check_balance(&self, user: &User) -> bool {
        user.balance >= self.amount
    }
}

impl fmt::Display for Transaction {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(
            f,
            "Sender: {}, Receiver: {}, Amount: {}, Signature: {}",
            self.sender, self.receiver, self.amount, self.signature
        )
    }
}

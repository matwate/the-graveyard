from cryptography.hazmat.primitives import hashes
from cryptography.hazmat.primitives.asymmetric import rsa, padding

# Function to generate a new RSA key pair
def generate_keys():
    """
    Generates a new RSA key pair.

    Returns:
        Tuple[rsa.RSAPrivateKey, rsa.RSAPublicKey]: A tuple containing the private and public keys.
    """
    private_key: rsa.RSAPrivateKey = rsa.generate_private_key(
        public_exponent=65537,
        key_size=512,
    )
    public_key: rsa.RSAPublicKey = private_key.public_key()
    return private_key, public_key

# Function to sign a transaction using a private key
def sign_transaction(private_key: rsa.RSAPrivateKey, transaction_data: str) -> bytes:
    """
    Signs the given transaction data using the provided private key.

    Args:
        private_key (rsa.RSAPrivateKey): The private key used for signing.
        transaction_data (str): The transaction data to be signed.

    Returns:
        bytes: The digital signature of the transaction data.
    """
    encoded_transaction_data: bytes = transaction_data.encode()
    signature: bytes = private_key.sign(
        encoded_transaction_data,
        padding.PSS(
            mgf=padding.MGF1(hashes.SHA256()),
            salt_length=padding.PSS.MAX_LENGTH
        ),
        hashes.SHA256()
    )
    return signature

# Function to verify a signature using a public key
def verify_signature(public_key: rsa.RSAPublicKey, transaction_data: str, signature: bytes) -> bool:
    """
    Verifies the signature of the given transaction data using the provided public key.

    Args:
        public_key (rsa.RSAPublicKey): The public key used for signature verification.
        transaction_data (str): The transaction data associated with the signature.
        signature (bytes): The digital signature to be verified.

    Returns:
        bool: True if the signature is valid, False otherwise.
    """
    encoded_transaction_data: bytes = transaction_data.encode()
    try:
        public_key.verify(
            signature,
            encoded_transaction_data,
            padding.PSS(
                mgf=padding.MGF1(hashes.SHA256()),
                salt_length=padding.PSS.MAX_LENGTH
            ),
            hashes.SHA256()
        )
        return True
    except Exception as e:
        return False
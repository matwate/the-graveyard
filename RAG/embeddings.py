from langchain_community.embeddings.ollama import OllamaEmbeddings
from langchain_chroma import Chroma
from langchain.schema.document import Document

def get_embedding_function():
    embeddings = OllamaEmbeddings(
        model = 'llama3.1',
    )
    return embeddings

CHROMA_PATH = "./chroma/"

def add_chroma(chunks: list[Document]):
    db = Chroma(
        persist_directory=CHROMA_PATH,
        embedding_function=get_embedding_function(),
    )
    db.add_documents(chunks)
    
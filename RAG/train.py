from docs import load_documents, split_documents
from embeddings import add_chroma
documents = load_documents()
chunks = split_documents(documents)
add_chroma(chunks)
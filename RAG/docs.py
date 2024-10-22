from langchain_community.document_loaders import PyPDFLoader
from langchain_text_splitters import RecursiveCharacterTextSplitter
from langchain.schema.document import Document


DOCUMENT_PATH = "./data/test.pdf"


def load_documents():
    doc_loader = PyPDFLoader(DOCUMENT_PATH)
    return doc_loader.load()


def split_documents(documents: list[Document]):
    text_splitter = RecursiveCharacterTextSplitter(
        chunk_size=1500,
        chunk_overlap=1000,
        length_function= len,
        is_separator_regex=False,
    )
    return text_splitter.split_documents(documents)


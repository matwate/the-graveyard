from embeddings import get_embedding_function
from langchain_chroma import Chroma
from langchain.prompts import ChatPromptTemplate
from langchain_community.llms.ollama import Ollama
PROMPT_TEMPLATE = """
Answer the question with the following text:
{context}
---
Answer the question based on the text above.
{question}

"""

CHROMA_PATH = "./chroma/"   

def query_rag(text: str):
    embedding_function = get_embedding_function()
    db = Chroma(
        persist_directory=CHROMA_PATH,
        embedding_function=embedding_function,
    )
    
    results = db.similarity_search_with_score(text, 15)
    context = "\n\n---\n\n".join([doc.page_content for doc, _score in results])
    prompt_template = ChatPromptTemplate.from_template(PROMPT_TEMPLATE)
    prompt = prompt_template.format(context=context, question=text)
    
    model = Ollama(model="llama3.1")
    response = model.invoke(prompt)
    print(response)
    return response
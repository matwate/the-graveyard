from query import query_rag


def main():
    query = "Could you give all the details required to properly execute experiment #6?"
    response = query_rag(query)
    
if __name__ == "__main__":
    main()
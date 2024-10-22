import flet as ft
import requests

API_BASE_URL = "http://localhost:8000"

def main(page):
    page.title = "Blockchain App"
    page.vertical_alignment = ft.MainAxisAlignment.START
    page.horizontal_alignment = ft.CrossAxisAlignment.START

    # Create a new user
    def create_new_user(e):
        response = requests.get(f"{API_BASE_URL}/newuser")
        if response.status_code == 200:
            pubkey = response.json()["public_key"]
            output.value = f"New user created with public key: {pubkey}"
            output.update()

    # View the blockchain
    def view_blockchain(e):
        response = requests.get(f"{API_BASE_URL}/")
        if response.status_code == 200:
            blockchain_str = response.text
            output.value = f"Blockchain:\n{blockchain_str}"
            output.update()

    # Add balance container
    user_index_input = ft.Ref[ft.TextField]()
    balance_input = ft.Ref[ft.TextField]()

    def show_add_balance_container(e):
        add_balance_container = ft.Container(
            content=ft.Column(
                [
                    ft.TextField(label="User Index", width=200, ref=user_index_input),
                    ft.TextField(label="Balance", width=200, ref=balance_input),
                    ft.ElevatedButton("Add Balance", on_click=add_balance),
                ]
            ),
            
            padding=ft.padding.all(20),
            border_radius=10,
        )
        page.add(add_balance_container)

    def add_balance(e):
        user_index = int(user_index_input.current.value)
        balance = int(balance_input.current.value)
        response = requests.get(f"{API_BASE_URL}/debug/setbalance/{user_index}/{balance}")
        if response.status_code == 200:
            output.value = f"Balance set for user {user_index} to {balance}"
            output.update()
        page.remove_at(-1)  # Close the add balance container

    
    sender_index_input = ft.Ref[ft.TextField]()
    receiver_index_input = ft.Ref[ft.TextField]()
    amout_input = ft.Ref[ft.TextField]()
    def show_make_transaction_container(e):
        make_transaction_container = ft.Container(
            content=ft.Column(
               [ ft.TextField(label="Sender User Index", width=200, ref=sender_index_input),
                ft.TextField(label="Receiver User Index", width=200, ref=receiver_index_input),
                ft.TextField(label="Amount", width=200, ref=amout_input),
                ft.ElevatedButton("Make Transaction", on_click=make_transaction)],    
            ),
            padding=ft.padding.all(20),
            border_radius=10,
        )
        page.add(make_transaction_container)
    def make_transaction(e):
        sender_index = int(sender_index_input.current.value)
        receiver_index = int(receiver_index_input.current.value)
        amount = int(amout_input.current.value)
        response = requests.get(f"{API_BASE_URL}/newtransaction/{sender_index}/{receiver_index}/{amount}")
        if response.status_code == 200:
            output.value = f"Transaction made from user {sender_index} to user {receiver_index} of amount {amount}"
            output.update()
        page.remove_at(-1)    
    # View transactions container
    def show_transactions_container(e):
        response = requests.get(f"{API_BASE_URL}/debug/viewPendingTransactions")
        if response.status_code == 200:
            transactions = response.json()
            
            transactions_container = ft.Container(
                
                content=(
                    ft.ListView(
                        controls=(
                            [
                                ft.Text(
                                    value=transaction_str,
                                    
                                    
                                )
                                for transaction_str in transactions
                            ]
                        )
                    )    
                ),
                padding=ft.padding.all(20),
                border_radius=10,
            )
            page.add(transactions_container)

    
    
    new_user_btn = ft.ElevatedButton("Create New User", on_click=create_new_user)
    view_chain_btn = ft.ElevatedButton("View Blockchain", on_click=view_blockchain)
    add_balance_btn = ft.ElevatedButton("Add Balance", on_click=show_add_balance_container)
    view_transactions_btn = ft.ElevatedButton("View Pending Transactions", on_click=show_transactions_container)
    new_transaction_btn = ft.ElevatedButton("New Transaction", on_click=show_make_transaction_container)
    output = ft.Text()

    buttons_column = ft.Column(
        [new_user_btn, view_chain_btn, add_balance_btn, view_transactions_btn, new_transaction_btn]
    )

    page.add(buttons_column, output)

ft.app(target=main)
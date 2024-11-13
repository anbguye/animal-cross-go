import plaid
from datetime import datetime, timedelta
from typing import List, Dict

class BankConnector:
    def __init__(self, plaid_client_id: str, plaid_secret: str):
        self.client = plaid.Client(
            client_id=plaid_client_id,
            secret=plaid_secret,
            environment='development'  # Change to 'production' for live data
        )
        
    def get_transactions(self, access_token: str, days: int = 1) -> List[Dict]:
        """Fetch transactions for the specified number of past days."""
        end_date = datetime.now().date()
        start_date = end_date - timedelta(days=days)
        
        try:
            response = self.client.transactions.get(
                access_token,
                start_date.isoformat(),
                end_date.isoformat()
            )
            return response['transactions']
        except plaid.errors.PlaidError as e:
            print(f"Error fetching transactions: {e}")
            return [] 
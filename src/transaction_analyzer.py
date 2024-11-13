from typing import List, Dict
from collections import defaultdict
import numpy as np
from datetime import datetime, timedelta

class TransactionAnalyzer:
    def __init__(self, historical_transactions: List[Dict]):
        self.historical_transactions = historical_transactions
        
    def categorize_spending(self, transactions: List[Dict]) -> Dict[str, float]:
        """Group transactions by category and sum amounts."""
        categories = defaultdict(float)
        for transaction in transactions:
            category = transaction.get('category', ['Unknown'])[0]
            categories[category] += transaction['amount']
        return dict(categories)
    
    def detect_anomalies(self, transactions: List[Dict]) -> List[Dict]:
        """Detect unusual spending patterns."""
        anomalies = []
        category_averages = self._calculate_category_averages()
        
        current_spending = self.categorize_spending(transactions)
        
        for category, amount in current_spending.items():
            if category in category_averages:
                if amount > category_averages[category] * 1.5:  # 50% more than usual
                    anomalies.append({
                        'category': category,
                        'current_amount': amount,
                        'usual_amount': category_averages[category],
                        'difference': amount - category_averages[category]
                    })
        
        return anomalies
    
    def generate_savings_tips(self, anomalies: List[Dict]) -> List[str]:
        ## Adding AI integration here
        """Generate personalized savings tips based on spending patterns."""
        tips = []
        for anomaly in anomalies:
            category = anomaly['category']
            difference = anomaly['difference']
            
            if category == 'Dining':
                tips.append(
                    f"You spent ${difference:.2f} more than usual on dining. "
                    "Consider meal prepping to reduce restaurant expenses."
                )
            elif category == 'Shopping':
                tips.append(
                    f"Your shopping expenses were ${difference:.2f} higher than normal. "
                    "Try making a shopping list and sticking to it."
                )            
        return tips
    
    def _calculate_category_averages(self) -> Dict[str, float]:
        """Calculate average spending by category from historical data."""
        return self.categorize_spending(self.historical_transactions) 
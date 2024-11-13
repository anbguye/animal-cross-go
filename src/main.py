import yaml
from bank_connector import BankConnector
from transaction_analyzer import TransactionAnalyzer
from notification_service import NotificationService

def load_config():
    with open('config/config.yaml', 'r') as f:
        return yaml.safe_load(f)

def main():
    # Load configuration
    config = load_config()
    
    # Initialize services
    bank = BankConnector(
        config['plaid']['client_id'],
        config['plaid']['secret']
    )
    
    # Get transactions
    daily_transactions = bank.get_transactions(
        config['plaid']['access_token'],
        days=1
    )
    historical_transactions = bank.get_transactions(
        config['plaid']['access_token'],
        days=30  # Get last 30 days for historical comparison
    )
    
    # Analyze transactions
    analyzer = TransactionAnalyzer(historical_transactions)
    spending = analyzer.categorize_spending(daily_transactions)
    anomalies = analyzer.detect_anomalies(daily_transactions)
    tips = analyzer.generate_savings_tips(anomalies)
    
    # Send notifications
    notifier = NotificationService(
        config['email'],
        config['sms']
    )
    
    for user in config['users']:
        if user.get('email'):
            notifier.send_email_summary(
                user['email'],
                spending,
                anomalies,
                tips
            )
        
        if user.get('phone'):
            notifier.send_sms_summary(
                user['phone'],
                spending,
                anomalies,
                tips
            )

if __name__ == '__main__':
    main() 
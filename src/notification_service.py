import smtplib
from email.mime.text import MIMEText
from typing import List, Dict
from twilio.rest import Client

class NotificationService:
    def __init__(self, email_config: Dict, sms_config: Dict):
        self.email_config = email_config
        self.sms_client = Client(
            sms_config['account_sid'],
            sms_config['auth_token']
        )
        self.sms_from = sms_config['from_number']
    
    def send_email_summary(self, 
                          to_email: str, 
                          spending: Dict[str, float],
                          anomalies: List[Dict],
                          tips: List[str]):
        """Send spending summary email."""
        body = self._format_summary(spending, anomalies, tips)
        
        msg = MIMEText(body)
        msg['Subject'] = 'Your Daily Spending Summary'
        msg['From'] = self.email_config['from_email']
        msg['To'] = to_email
        
        with smtplib.SMTP_SSL(self.email_config['smtp_server']) as server:
            server.login(
                self.email_config['username'],
                self.email_config['password']
            )
            server.send_message(msg)
    
    def send_sms_summary(self,
                        to_number: str,
                        spending: Dict[str, float],
                        anomalies: List[Dict],
                        tips: List[str]):
        """Send spending summary via SMS."""
        body = self._format_summary(spending, anomalies, tips, sms=True)
        
        self.sms_client.messages.create(
            body=body,
            from_=self.sms_from,
            to=to_number
        )
    
    def _format_summary(self,
                       spending: Dict[str, float],
                       anomalies: List[Dict],
                       tips: List[str],
                       sms: bool = False) -> str:
        """Format the summary message for email or SMS."""
        lines = ['Your Daily Spending Summary:']
        
        # Add spending by category
        lines.append('\nSpending by Category:')
        for category, amount in spending.items():
            lines.append(f'- {category}: ${amount:.2f}')
        
        # Add anomalies
        if anomalies:
            lines.append('\nUnusual Spending Patterns:')
            for anomaly in anomalies:
                lines.append(
                    f"- {anomaly['category']}: ${anomaly['current_amount']:.2f} "
                    f"(usually ${anomaly['usual_amount']:.2f})"
                )
        
        # Add savings tips
        if tips:
            lines.append('\nSavings Tips:')
            for tip in tips:
                lines.append(f'- {tip}')
        
        return '\n'.join(lines) 
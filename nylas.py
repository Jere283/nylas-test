from nylas import Client
from dotenv import dotenv_values

config = dotenv_values(".env")


apiKey = config.get("API_KEY")
grantId = config.get("GRANT_ID")
uri = "https://api.us.nylas.com"
threadId = "18df121c8f5b46df"

nylas = Client(apiKey, uri)

thread = nylas.threads.find(
    grantId,
    threadId,
)

foldersIds = thread.data.message_ids

for folderId in foldersIds:
    message = nylas.messages.find(
        grantId,
        folderId,
    )
    print(message.data.subject)
    print(message.data.body)

newM = nylas.messages.send(
    grantId,
    request_body={
        "to": [{"name": "Dominik", "email": "dominic.j@nylas.com"}],
        "reply_to": [{"name": "Dominik", "email": "dominic.j@nylas.com"}],
        "subject": "My Favorite Snack Jeremy Figueroa",
        "body": "My favorite snack is the chocolate",
    },
)


print(newM)

from nylas import Client

apiKey = "nyk_v0_YEZVfhzQRYWIjpKLLJ74JkEDqLCl7RIPqFVOILxItlhjN2nPbYgyzOAaCKNFwKRy"
grantId = "c0f2f708-567b-4a10-9b28-efad7d156283"
uri = "https://api.us.nylas.com"
threadId = "18df121c8f5b46df"

custom_headers = {"Accept": "message/rfc822"}

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

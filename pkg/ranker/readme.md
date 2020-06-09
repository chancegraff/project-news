1. User votes a post up, userID is added to article and articleID is added to user
2. User's info enters a queue with the ID of the article
3. 2-5 minutes later, the queue is processed via batch
4. Every user is collected by article ID
5. A new record is created with the userID and the articleID
6. Endpoints are exposed to retrieve by article ID or by user ID

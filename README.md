# SimpleStore

SimpleStore is a demo JSON API webapp which stores and retrieves arbitrary data.

There are two endpoints:

## Data Storage

### URL

POST /messages/

### Return Value

  If successful, the reponse will be:

    {"digest":<sha sum of post body>}

## Data Retrieval

### URL

    GET /messages/<sha sum>

If successful, the response will be the contents of the message.


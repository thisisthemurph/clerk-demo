# Clerk auth demo

Uses React on the front-end and authenticates the Go backend using the Clerk provided JWT.

## Setup

**./client/.env.local**

Create and add the following:

```
VITE_CLERK_PUBLISHABLE_KEY=pk_test_...
```

**./server/.env**

Create and add the following:

``` 
CLERK_KEY=sk_test_...
PORT=42069
```

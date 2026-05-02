### DSA
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

 

Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps
Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step
 

Constraints:

1 <= n <= 45



---
### Design API and DB Schema for URL Shortener

### APIs
- POST /urls  → create short URL
- GET /{code} → redirect (302)

### DB
urls(id, code UNIQUE, long_url, expiry_at)

### Logic
create:
- save → get id
- id → base62 → code

redirect:
- code → lookup
- if expired → 410
- else → 302

### Optimization
- Redis: code → long_url

---

### Design Cart & Checkout System

### APIs
- POST /cart/items
- GET  /cart
- DELETE /cart/items/{id}

- POST /checkout
- POST /payments


### DB
- carts(id, user_id)
- cart_items(id, cart_id, product_id, qty, price)

- orders(id, user_id, total_amount, status)
- order_items(id, order_id, product_id, qty, price)


### Flow
cart:
- add/update items

checkout:
- validate cart + stock
- create order (PENDING)
- payment call
- success → CONFIRMED
- failure → FAILED


### HLD

- Idempotency:
  prevent duplicate orders (use idempotency key)

- Consistency:
  stock check + order creation should be atomic (transaction / lock)

- Concurrency:
  avoid overselling (DB lock / inventory service)

- Async:
  use queue for payment confirmation, email, notifications

- Scaling:
  cache cart (Redis), DB read replicas

- Fault Handling:
  payment failure → retry / mark FAILED

- Price Safety:
  store price in order_items (not from product table later)


### LLD

- Strategy:
  payment methods (Card / UPI / Wallet)

- Factory:
  create payment handler based on type

- State:
  order status (PENDING → CONFIRMED → FAILED)

- Singleton:
  DB / Redis connection

- Observer (optional):
  order placed → notify email / inventory / analytics

# Go Interview Assignment


## The Assignment: "FairSplit API"

The candidate will build a REST API that mimics the core functionality of Splitwise. Users can create groups, add expenses, check their balances, and—crucially—simplify the group's debts to minimize the number of transactions required to settle up.

### Core Requirements

1. **Authentication & User Management**
   - Endpoints to register and log in (`POST /register`, `POST /login`).
   - Secure password storage.
   - Protected routes using JWT or Session middleware.

2. **Groups & Expenses (The Core Elements)**
   - `POST /groups`: Create a new group and add users to it.
   - `POST /groups/{id}/expenses`: Add an expense to the group.
     - *Example Payload:* User A paid $120. It is split equally between User A, User B, and User C. (Result: B owes A $40, C owes A $40).
     - *Scope limit for freshers:* We can restrict splitting to "Equal Splits" to keep the focus on the architecture rather than complex fractional math.

3. **Dashboard (Balances)**
   - `GET /dashboard`: Returns a summary for the logged-in user.
   - Shows total amount the user owes others, total amount others owe the user, and a detailed breakdown (e.g., "You owe Alice $40", "Bob owes you $10").

4. **The Algorithmic Challenge: Simplify Debts**
   - `GET /groups/{id}/simplify`: Returns the optimized settlement plan for the group.
   - **The Problem:** If A owes B $10, and B owes C $10, the system should simplify this to: A owes C $10.
   - The candidate must implement the "Minimize Cash Flow" algorithm (often solved using a Greedy approach by calculating net balances for each user and matching the largest debtors with the largest creditors).

## Technical Scope

This assignment judges API design, relational database modeling, and algorithmic problem solving.

- **Database:**  The candidate must model `Users`, `Groups`, `GroupMembers`, `Expenses`, and `DebtEdges`.
- **"Simplify Debts" algorithm** 

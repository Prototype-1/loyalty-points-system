# Loyalty Points System (Go + PostgreSQL)

A backend system that allows users to **earn and redeem loyalty points** through transactions. Features include point calculation, balance tracking, redemption logic, expiration handling, and user authentication.

This project was originally built as part of a machine test (48-hour limit). The version here includes **post-submission improvements** based on self-review and real-world production standards.

---

##  Features
-  **JWT-based authentication**
-  **Earn points** via purchase transactions
-  **Redeem points**, with proper deduction and balance update
-  **Points history** with pagination and filtering by date/status
-  **Automated expiration** of unused points via cron job
-  PostgreSQL + GORM for data storage
-  Detailed redemption feedback (success/failure, remaining balance)

---

##  What was improved after initial submission
> After submitting the initial version, I revisited the code to fix missing requirements and logic gaps:

- Added **transactional safety** using DB transactions for earning and redeeming
- Corrected point balance logic to exclude **both** redeemed and expired points
- Added `reason` field to transaction records for **traceability** (earned/redeemed for what)
- Improved redemption response to return **updated balance**
- Ensured clean separation of logic for maintainability

---

##  Tech Stack
- **Language**: Go
- **Database**: PostgreSQL
- **ORM**: GORM
- **Scheduler**: [`robfig/cron`](https://pkg.go.dev/github.com/robfig/cron/v3)

---

##  API Overview
| Endpoint               | Description                             |
|------------------------|-----------------------------------------|
| `POST /auth/login`     | User login + JWT                        |
| `POST /transactions/add`   | Add a new earning transaction           |
| `POST /points/redeem`  | Redeem points                           |
| `GET /points/balance`  | View current balance                    |
| `GET /points/history`  | View transaction history (paginated)    |

---

##  Disclaimer
This was originally a timed task. The current version reflects my own standards and expectations after reviewing my initial mistakes. It now serves as both a case study and a solid reference for point-based systems.

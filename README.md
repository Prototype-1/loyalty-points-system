# Loyalty Points System API 

This is a Go-based **Loyalty Points System** that allows users to earn and redeem points through transactions. It includes APIs for managing loyalty points, transactions, 
view user balances and history with optional filtering.

## Features
- Authentication using JWT
- User-based **loyalty points management** (earn & redeem)
- Transaction tracking
- Database integration with PostgreSQL
- Scheduled tasks using `cron` (for automations)
-  View balance and history

## Tech Stack
- **Backend**: Go
- **Database**: PostgreSQL
- **Scheduler**: `github.com/robfig/cron/v3`
- **ORM**: GORM

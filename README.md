# GoLedger - Practical implementation of Bitcoin key concepts through a simplified blockchain built in Go

## Description

GoLedger is a practical and educational implementation of the core concepts presented in the Bitcoin paper, built as a simplified blockchain system in Go. The project focuses on reproducing the essential mechanisms behind decentralized systems while keeping the implementation minimal, transparent, and easy to reason about.

Rather than replicating the full complexity of Bitcoin, this project is designed to provide a clear and hands-on understanding of how its foundational components interact.

---

## Core Concepts Covered

- Cryptographic hash functions (SHA-256)  
- Block structure and hash-based chaining  
- Data integrity and immutability  
- Proof of Work (simplified mining)  
- Transaction modeling  
- Digital signatures and asymmetric cryptography  
- Basic validation and consensus rules  
- Peer-to-peer communication (simplified)

---

## Objective

The main goal of GoLedger is to demonstrate how trust can emerge in a distributed and potentially adversarial environment without relying on a central authority.

This project serves as a learning tool for understanding the architecture, design decisions, and trade-offs behind blockchain-based systems.

---

## Technology Stack

- Go (Golang)  
- Standard library (`crypto/sha256`, `encoding/hex`, etc.)

---

## Getting Started

Clone the repository and run the project:

```bash
git clone https://github.com/your-username/goledger.git
cd goledger
go run .

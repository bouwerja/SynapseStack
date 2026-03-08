# 🛠️ SynpaseStack: AI-Driven Venture Intelligence

**SynpaseStack** is a high-performance business intelligence engine that predicts startup success and identifies market gaps using real-world industry benchmarks. Engineered to run on a resource-constrained 2GB RAM / 2vCPU environment, it demonstrates the intersection of Data Engineering, Machine Learning, and Systems Optimization.

---

## 🚀 The Problem

Aspiring entrepreneurs often lack data-driven insights into their specific niche. Most "Market Research" tools are either too expensive or too generic. **SynpaseStack** provides a "Crystal Ball" for founders by scraping real-time industry data and running local ML inference to calculate success probabilities.

---

## 🏗️ Tech Stack & Constraints

To prove technical discipline, this project was architected to stay within a $15/month AWS Lightsail footprint:

- Backend: Golang (Fiber) — Chosen for its minimal memory footprint and native concurrency.

- Database: PostgreSQL + pgvector — A unified relational and vector store for company embeddings.

- AI/ML: Python (XGBoost) & Llama 3 (Quantized) — Local inference via Ollama to avoid high API costs.

- Frontend: Nuxt 3 (SSG) — Statically generated to ensure zero-lag delivery to 1,000+ daily users.

- Infrastructure: Docker, Nginx, and Linux Cgroups for strict memory capping.

---

## 🧠 System Architecture

- Ingestion Engine: A concurrent Go-based scraper (using Colly) pulls normalized data from SBA.gov, Crunchbase, and industry forums.

- Vector Similarity: Uses pgvector to perform "Nearest Neighbor" searches. It finds the top 50 "Lookalike" businesses to a user's profile to identify common failure points.

- Inference Layer: A sidecar service runs an XGBoost model to predict success probability based on capital, location, and industry saturation.

- Optimization Layer: Utilizes Cgroups and Zram to ensure the LLM (Llama 3 3B) doesn't starve the PostgreSQL process of memory.

---

## 🛠️ Performance Optimizations

In a 2GB environment, "Standard" configurations crash. This project implements:

- Capped Connection Pooling: Go-side limits (SetMaxOpenConns(25)) to prevent PostgreSQL process bloat.

- IVFFlat Indexing: Optimized vector search to prioritize RAM savings over HNSW memory overhead.

- GGUF 4-bit Quantization: Shrinking the LLM footprint by 70% to allow for local summary generation.

- SSD Swap Tuning: A 4GB SSD-backed swap file to handle burst traffic from the 1,000+ daily user target.

---

## 📂 Repository Structure

```Bash

├── cmd/api # Go Backend (The Orchestrator)
├── cmd/scraper # Go Ingestion Engine
├── ml/models # Python XGBoost & Quantized GGUF files
├── web/ # Nuxt 3 Frontend (Statically Generated)
├── docker-compose.yml # Resource-limited environment config
└── scripts/optimize # Linux Cgroup & Memory tuning scripts
```

---

## 📈 Key Milestones

    [ ] Phase 1: Concurrent Go Scraper & PostgreSQL schema.

    [ ] Phase 2: Entity Resolution & Vector Embedding generation.

    [ ] Phase 3: Local ML Inference (XGBoost + Llama 3).

    [ ] Phase 4: Nuxt 3 Dashboard & Global Deployment.

---

## 🤝 Contact & Portfolio

Developer: [Jason Bouwer]

Focus: Backend Engineering / AI-Ops

LinkedIn: [ ]

Website: [ ]

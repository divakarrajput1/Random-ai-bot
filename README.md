# Random AI-Bot 🤖

**Random AI-Bot** is a Go-based AI service that responds to text with a random sentiment: **positive**, **neutral**, or **negative**.  

This project demonstrates:  
- Deploying an AI model as an **HTTP service**  
- Containerizing it and hosting it on **Docker Hub**  
- Running it on **Kubernetes**  
- Using **KEDA HTTP ScaledObject** to **auto-scale pods based on incoming HTTP traffic**  

Perfect for learning **cloud-native AI deployment and autoscaling concepts** 🚀  

---

## ✨ Features
- **REST API**
  - `POST /infer` → Returns a random sentiment + confidence score  
  - `GET /healthz` → Health check  
- Built with **Go**  
- **Dockerized** and hosted on **Docker Hub**  
- **Kubernetes ready** with manifests in `k8s/`  
- **Auto-scaling HTTP workloads** with **KEDA**  

---

## 🛠️ Tech Stack
- **Language**: Go (Golang)  
- **Containerization**: Docker Hub  
- **Orchestration**: Kubernetes  
- **Autoscaling**: KEDA (HTTP ScaledObject)  

---

## 🚀 Getting Started

### 1. Deploy on Kubernetes
Apply all Kubernetes manifests (Deployment, Service, KEDA ScaledObject):
```bash
kubectl apply -f k8s/

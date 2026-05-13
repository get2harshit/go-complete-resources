# HLD — AI Powered Coding Evaluation Platform

## 1. User Flow

- Users solve coding/backend/system design problems through the frontend platform.
- Submissions are sent to backend APIs for validation and evaluation orchestration.
- Evaluation jobs are pushed into async queues instead of processing synchronously.
- Worker services process submissions and interact with AI evaluation providers.
- Generated feedback, scores, and reports are stored and shown on user dashboards.
- Users receive near real-time updates during evaluation lifecycle.

---

## 2. Infrastructure

- Frontend is deployed separately from backend services for independent scaling.
- Backend APIs are built using modular service architecture.
- Redis-backed queues are used for async evaluation workflows.
- PostgreSQL/MongoDB stores submissions, reports, and evaluation metadata.
- Redis is used for caching dashboards and evaluation results.
- Object storage is used for storing generated reports and large assets.
- Stateless APIs allow future horizontal scaling.

---

## 3. Architecture Decisions

- Currently deployed on free-tier infrastructure, so the architecture is optimized for simplicity and cost efficiency.
- A modular monolith architecture is sufficient at current traffic scale.
- Evaluation processing is isolated from API request lifecycle to support future worker-based scaling.
- Async processing was preferred because AI evaluation latency is unpredictable.
- Queue-based architecture improves resiliency and prevents API blocking.
- Components are designed to later support service extraction without major refactoring.

---

## 4. Scalability Thinking

- Worker services can scale independently from API services during traffic spikes.
- Stateless backend services simplify horizontal scaling.
- Queue-based processing smooths sudden spikes near assignment deadlines.
- Redis caching reduces repeated database and AI evaluation load.
- Evaluation pipelines are designed to support future multi-worker execution.
- Architecture supports future migration toward distributed evaluation services.

---

## 5. Bottlenecks & Real Problems

### Cold Starts

- Free-tier deployments introduced cold-start latency after inactivity.
- Sleeping infrastructure caused delayed first-request response times.

### AI Evaluation Latency

- LLM evaluation sometimes took 20–40 seconds depending on prompt complexity.
- Synchronous evaluation caused poor user experience during long-running requests.
- Async job processing improved frontend responsiveness significantly.

### Concurrent Evaluations

- Multiple simultaneous evaluations slowed worker throughput.
- Queue-based orchestration helped stabilize processing during peak traffic.

### External API Limits

- AI providers introduced token and rate limitations during heavy usage.
- Prompt optimization reduced unnecessary token consumption.

### Dashboard Performance

- Repeated dashboard and report fetches increased database load.
- Redis caching improved report retrieval latency.

---

## 6. Reliability & Failure Handling

- AI provider failures occasionally occurred due to request timeout or temporary outages.
- Retry mechanisms were implemented for transient provider failures.
- Submission state tracking prevents data loss during partial failures.
- Failed evaluations are preserved with retry capability instead of discarding requests.
- Structured logging improved debugging of prompt failures and evaluation errors.
- Job status transitions like PENDING, PROCESSING, FAILED, and COMPLETED improve workflow tracking.
- Queue isolation prevents one failed evaluation from affecting unrelated jobs.

---

## 7. Traffic Handling

- Traffic spikes mainly occurred near assignment deadlines and mock interview sessions.
- Concurrent AI evaluations created temporary processing bottlenecks.
- Rate limiting was implemented to prevent abuse and uncontrolled API usage.
- Async queues stabilized throughput during peak concurrency periods.
- Cached evaluations reduced repeated compute load during heavy usage.

---

## 8. Performance Optimization

- Evaluation workflows were moved to background workers to reduce request latency.
- Prompt engineering reduced token usage and evaluation response size.
- Redis caching minimized repeated AI calls for previously evaluated submissions.
- Lightweight payloads improved API responsiveness.
- Async frontend polling/WebSocket updates improved user experience during long evaluations.

---

## 9. Cost Optimization

- Since AI calls are expensive, prompt engineering was used to minimize token usage.
- Cached evaluation reports reduced repeated inference costs.
- Modular monolith architecture reduced operational overhead at current scale.
- Async workers improved compute utilization efficiency.
- Infrastructure decisions prioritized low operational cost while keeping future scalability in mind.
- AI evaluations are triggered only for evaluation-heavy workflows to avoid unnecessary inference costs.

---

## 10. Future Improvements

- Independent evaluation worker scaling based on queue depth.
- Multi-provider AI routing for fallback and cost optimization.
- Event-driven architecture for notifications and analytics.
- Streaming evaluation feedback for faster perceived responsiveness.
- Dedicated analytics pipeline for evaluation insights and platform metrics.

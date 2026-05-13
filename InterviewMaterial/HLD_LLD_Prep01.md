
# HLD Talking Points — Textile VTO + Catalogue + Order Management

## Architecture

- Modular monolith architecture was chosen to reduce operational complexity while keeping services logically separated.
- The system is designed to be microservice-ready for future scaling requirements.
- Frontend communicates with backend APIs through authenticated REST endpoints.
- JWT-based authentication secures APIs and user sessions.

---

## Scalability

- Compute-heavy workflows like VTO rendering and catalogue generation are processed asynchronously using worker queues.
- Queue-based processing prevents long-running tasks from blocking API request lifecycle.
- Worker services can scale independently from API services during bulk catalogue generation.
- Bulk catalogue generation supports parallel processing for higher throughput.
- The architecture supports horizontal scaling because processing workers are stateless.
- Traffic spikes mainly occur during simultaneous batch uploads and generation requests.

---

## Storage & Delivery

- Generated textile assets are stored in object storage instead of local filesystem for scalability and durability.
- CDN delivery is used for generated catalogue images to reduce latency and bandwidth usage.
- Signed URLs are used for secure upload and access of media assets.

---

## Performance & Optimization

- Redis is used for caching frequently accessed catalogue previews and metadata.
- Image preprocessing pipelines optimize uploaded assets before AI generation workflows.
- Background removal and normalization improve consistency in generated outputs.
- Async processing improves responsiveness for latency-sensitive workflows.

---

## Failure Handling

- Retry mechanisms handle temporary AI provider failures during image rendering.
- Generation jobs maintain state tracking like PENDING, PROCESSING, FAILED, and COMPLETED.
- Processing workflows are designed to fail gracefully without affecting unrelated jobs.

---

## Business Workflows

- Order workflows follow state transitions such as CREATED, CONFIRMED, SHIPPED, and DELIVERED.
- Async notifications can later be integrated using event-driven workflows.


- The architecture is extensible enough to support multiple AI/image providers in future.
- Prompt and rendering configurations are isolated from business logic for easier experimentation.


---

# LLD Talking Points — Textile VTO + Catalogue + Order Management

## Code Structure

- The backend follows layered architecture with controller, service, repository, and provider separation.
- Controllers are responsible only for request validation and authentication handling.
- Business logic is isolated inside service layers for maintainability and testability.
- Repository abstraction prevents database-specific logic from leaking into business workflows.

---

## Abstractions & Interfaces

- AI/image generation integrations are abstracted behind provider interfaces.
- Provider abstraction allows switching between different image generation services without changing core logic.
- Interfaces are used to improve extensibility and support dependency injection.
- Services communicate using structured payload contracts to improve consistency.

---

## Workflow Design

- Image generation requests are handled through dedicated orchestration services.
- Order management logic is separated from catalogue generation workflows to maintain loose coupling.
- Watermarking and branding are implemented as configurable processing stages.
- Asset processing pipeline is designed as reusable utility modules instead of tightly coupled logic.


- DTO validation ensures malformed requests are rejected early in request lifecycle.
- Validation layers ensure uploaded assets meet supported resolution and format constraints.
- Retry handling is implemented at worker level for transient processing failures.
- Job state persistence helps recover interrupted rendering workflows.
- Idempotency checks prevent duplicate generation requests during retries.
- Processing workflows are designed to fail gracefully without affecting unrelated jobs.

---

## Database & Persistence

- Database indexing is applied on order and asset metadata for faster retrieval.
- Repository layer centralizes persistence logic for consistency and easier maintenance.

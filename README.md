# Forge

Forge is a **personal learning project** focused on understanding **systems programming concepts** by building a small, distributed task execution system from scratch.

It is **not** a framework and **not** a production product. Forge exists to progressively move from application-level concerns toward **low-level systems thinking**: correctness, concurrency, resource control, failure handling, and observability.

---

## What is Forge?

Forge executes **declarative, data-defined tasks** (*jobs*) composed of ordered *steps*.

Jobs are treated as **data, not code**, enabling:

* Inspection and validation before execution
* Reproducibility
* Clear execution semantics

---

## Design Principles

* **Learning-first**: clarity over convenience
* **Incremental complexity**: features are added only when the concept is understood
* **Minimal magic**: explicit behavior, no heavy frameworks
* **Clear boundaries**: each layer has a well-defined responsibility
* **Failure is explicit**: errors and crashes are part of the model

---

## Architecture (Layered)

Forge is structured as a small set of explicit layers. Each layer has a single responsibility and minimal overlap with others.

### Orchestrator *(coordination and scheduling · Go)*

Responsible for:

* Job orchestration and state management
* Scheduling and coordination of work
* Concurrency and lifecycle control
* Failure handling and retries

This layer represents the **control plane** of the system.

---

### CLI *(user interface · Go)*

Responsible for:

* User interaction and command parsing
* Input validation
* Triggering orchestration actions
* Presenting results and errors

The CLI is intentionally thin and delegates all real work to other layers.

---

### Log *(observability primitives · Go)*

Responsible for:

* Structured log generation
* Layer-aware log attribution
* Deterministic log file layout and naming

Logging is treated as a **first-class systems concern**, not a debugging afterthought.

---

### Engine *(step execution · Rust)*

Responsible for:

* Executing steps correctly and deterministically
* Resource and process management
* OS-level interactions (files, processes, signals, permissions)
* Enforcing safety and correctness boundaries

Rust is used deliberately to force explicit handling of:

* Ownership and lifetimes
* Error propagation
* Concurrency primitives

This layer represents the **execution plane** of the system.

---

## Non-Goals

Forge intentionally avoids:

* Production readiness guarantees
* Backward compatibility
* Feature completeness
* Polished UX

Breaking changes are expected if they improve understanding.

---

## Project Status

Forge is under active, exploratory development.

Expect frequent refactors, changing abstractions, and experimental code — each driven by a concrete learning objective.

---

## License

This project is for personal and educational use.

License details may be added later if needed.


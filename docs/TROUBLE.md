This document outlines notable technical challenges encountered during the development and deployment of the DevOps Challenge project, along with solutions and workarounds.

---

❌ Go Module Directory Mismatch
Problem:
Initial project layout placed go.mod and go.sum inside the `app/` directory, while travis expected them in the root.

Solution:
Moved go.mod and go.sum to the root of the repository. This simplified Docker builds and aligned with Go tooling expectations. The app source (main.go) remained in `app/`, and import paths were adjusted if needed.

---

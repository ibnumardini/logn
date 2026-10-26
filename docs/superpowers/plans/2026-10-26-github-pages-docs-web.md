# GitHub Pages Deployment for `docs/web` Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Publish the static site in `docs/web` to GitHub Pages on pushes to `master` and on manual workflow runs.

**Architecture:** Add one GitHub Actions workflow with a read-only checkout/build job that uploads `docs/web` as a Pages artifact, followed by the official Pages deployment action. Use explicit Pages permissions, the `github-pages` environment, and concurrency cancellation; no build tool or `gh-pages` branch is needed.

**Tech Stack:** GitHub Actions YAML, `actions/checkout@v4`, `actions/configure-pages@v5`, `actions/upload-pages-artifact@v3`, `actions/deploy-pages@v4`.

---

## Chunk 1: GitHub Pages workflow

### Task 1: Add the artifact-based Pages workflow

**Files:**
- Create: `.github/workflows/pages.yml`
- Read: `docs/web/index.html`, `docs/superpowers/specs/2026-10-26-github-pages-docs-web-design.md`

- [ ] **Step 1: Create the workflow metadata and triggers**

  Define workflow name `Deploy docs/web to GitHub Pages`, trigger on pushes to `master`, and add `workflow_dispatch`.

- [ ] **Step 2: Add permissions and concurrency**

  Set workflow-level permissions to `contents: read`, `pages: write`, and `id-token: write`. Configure a `pages` concurrency group with `cancel-in-progress: true`.

- [ ] **Step 3: Add the build job**

  Create an `ubuntu-latest` job that checks out the repository with `actions/checkout@v4`, configures Pages with `actions/configure-pages@v5`, and uploads `./docs/web` with `actions/upload-pages-artifact@v3`.

- [ ] **Step 4: Add the deployment job**

  Create a `deploy` job that depends on `build`, targets the `github-pages` environment, and deploys with `actions/deploy-pages@v4`. Give the deploy step `id: deployment` and configure the environment URL as `${{ steps.deployment.outputs.page_url }}`.

- [ ] **Step 5: Validate the workflow shape**

  Run:

  ```bash
  python3 - <<'PY'
  from pathlib import Path
  import yaml

  path = Path(".github/workflows/pages.yml")
  workflow = yaml.safe_load(path.read_text())
  triggers = workflow.get("on", workflow.get(True))
  assert triggers["push"]["branches"] == ["master"]
  assert "workflow_dispatch" in triggers
  assert workflow["permissions"] == {
      "contents": "read",
      "pages": "write",
      "id-token": "write",
  }
  assert workflow["concurrency"]["group"] == "pages"
  assert workflow["concurrency"]["cancel-in-progress"] is True
  build = workflow["jobs"]["build"]
  deploy = workflow["jobs"]["deploy"]
  assert build["runs-on"] == "ubuntu-latest"
  assert deploy["needs"] == "build"
  assert deploy["environment"]["name"] == "github-pages"
  steps = build["steps"]
  assert any(step.get("uses") == "actions/checkout@v4" for step in steps)
  assert any(step.get("uses") == "actions/configure-pages@v5" for step in steps)
  upload = next(step for step in steps if step.get("uses") == "actions/upload-pages-artifact@v3")
  assert upload["with"]["path"] == "./docs/web"
  deployment = next(step for step in deploy["steps"] if step.get("uses") == "actions/deploy-pages@v4")
  assert deployment["id"] == "deployment"
  assert deploy["environment"]["url"] == "${{ steps.deployment.outputs.page_url }}"
  print("workflow checks passed")
  PY
  ```

  If PyYAML is unavailable, install it only for validation with `python3 -m pip install pyyaml`, then rerun the check.

- [ ] **Step 6: Run the existing Go tests**

  Run `go test ./...`. Expected: the existing package tests pass; the workflow does not modify Go behavior.

- [ ] **Step 7: Review the diff and commit**

  Run `git diff --check` and `git status --short`, then commit:

  ```bash
  git add .github/workflows/pages.yml
  git commit -m "ci: deploy docs web to github pages"
  ```

## Chunk 2: Final acceptance

### Task 2: Confirm deployment configuration

**Files:**
- Validate: `.github/workflows/pages.yml`
- Validate: `docs/web/index.html`

- [ ] **Step 1: Confirm the site entry point exists**

  Run `test -s docs/web/index.html`. Expected: exit code 0.

- [ ] **Step 2: Confirm no unrelated files changed**

  Run `git status --short`. Expected: clean working tree after the workflow commit.

- [ ] **Step 3: Confirm the final commit has no co-author trailer**

  Run `git log -1 --format=%B | grep -q '^Co-authored-by:'` and invert the result. Expected: no matching trailer.

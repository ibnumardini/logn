# GitHub Pages Deployment for `docs/web`

## Goal

Deploy the static documentation site in `docs/web` to GitHub Pages automatically
when changes land on `master`, with an option to run the deployment manually.

## Design

Add `.github/workflows/pages.yml` using GitHub's artifact-based Pages deployment:

1. Trigger on pushes to `master` and `workflow_dispatch`.
2. Grant the minimum required permissions:
   - `contents: read`
   - `pages: write`
   - `id-token: write`
3. Check out the repository before reading site files.
4. Use the official Pages setup and artifact actions:
   - `actions/checkout`
   - `actions/configure-pages`
   - `actions/upload-pages-artifact`
   - `actions/deploy-pages`
   Pin each action to its major version.
5. Upload only `./docs/web` as the Pages artifact.
6. Deploy the artifact through the `github-pages` environment.
7. Use concurrency so a newer deployment supersedes an in-progress deployment.

No build tool, package installation, Jekyll processing, or `gh-pages` branch is
needed because `docs/web/index.html` is already a complete static site.

## Validation

- Parse the workflow as YAML.
- Assert the workflow has the required triggers, permissions, action steps, and
  `./docs/web` upload path.
- Run the existing Go test suite to ensure the workflow change does not affect
  the library.

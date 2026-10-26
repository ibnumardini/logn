# Logn Static Documentation Landing Page Design

## Goal

Create a single static HTML developer-docs landing page for Logn, the existing Go
logging library. The page should make installation, configuration, usage, cron
zipping, contributing, and licensing easy to scan without introducing a
frontend build step.

## Chosen approach

Use one self-contained `index.html` with inline CSS and Alpine.js loaded from its
CDN. This keeps the deliverable deployable by copying one file, matches the
repository's current lack of a frontend toolchain, and preserves progressive
enhancement through regular anchor navigation.

## Page structure

The page will contain:

1. Sticky top bar with Logn branding, documentation navigation, and GitHub/install
   actions.
2. Hero describing Logn as a lightweight Go logging library.
3. Feature strip for daily/type logs, Telegram reporting, monthly zipping, and
   cleanup.
4. Installation section based on the existing README command.
5. Configuration section with the supplied JSON example and concise field
   explanations.
6. Usage section for `InfoLog`, `WarningLog`, and `ErrorLog`.
7. Cron zip section based on the existing `CronZip` usage.
8. Contributing and MIT license footer links.

Content will be grounded in the current README and source behavior; the page
will not claim features that are not documented by the project.

## Alpine.js behavior

Use a small Alpine state object to provide:

- Responsive mobile navigation open/close behavior.
- Active documentation section tracking while scrolling.
- Copy-to-clipboard controls for install, configuration, and code snippets.
- A brief success state after copying.
- An explicit fallback message when clipboard access is unavailable.

Anchored links remain usable when JavaScript is unavailable. There is no
backend, bundler, form submission, or persistent client-side state.

## Visual and accessibility direction

Use a near-black terminal-inspired palette with warm off-white text, amber
highlights, muted green status accents, subtle grid/noise treatment, and
monospace code surfaces. The layout must remain readable on mobile and use a
wider two-column documentation layout on desktop.

Include semantic headings, visible keyboard focus states, adequate contrast,
reduced-motion-friendly transitions, and accessible labels for navigation and
copy controls.

## Validation

Validate the static file structure, then open the page locally in a browser to
check:

- Desktop and mobile layout.
- Anchor navigation and active-section behavior.
- Mobile menu open/close behavior.
- Copy success and clipboard fallback behavior.
- Focus visibility and readable contrast.

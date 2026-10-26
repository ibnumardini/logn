# Logn Static Documentation Landing Page Design

## Goal

Create a single static HTML developer-docs landing page for Logn, the existing Go
logging library. The page should make installation, configuration, usage, cron
zipping, contributing, and licensing easy to scan without introducing a
frontend build step.

## Chosen approach

Use one single-file `index.html` with inline CSS and Alpine.js loaded from its
CDN. “Single-file” describes the repository deliverable; the page has an
explicit external runtime dependency on `https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js` (pin the concrete
3.x version in the implementation) and must remain readable if that request is
blocked. This keeps the deliverable
deployable by copying one file, matches the repository's current lack of a
frontend toolchain, and preserves progressive enhancement through regular
anchor navigation.

## Scope and deliverables

The only implementation deliverable is `/index.html`. Do not add a build
system, backend, analytics, external fonts, persistent storage, or a second
frontend asset file. The README does not need to be rewritten; the page itself
may link to the repository at
`https://github.com/ibnumardini/logn`, the README at
`https://github.com/ibnumardini/logn/blob/master/README.md`, and the sample at
`https://github.com/ibnumardini/logn/blob/master/logn_config_sample.json`.
Footer links must target
`https://github.com/ibnumardini/logn/blob/master/CONTRIBUTING.md` and
`https://github.com/ibnumardini/logn/blob/master/LICENSE.md`.

## Page structure

The page will contain:

1. Sticky top bar with Logn branding, documentation navigation, and GitHub/install
   actions. The GitHub action targets the repository URL; the install action
   targets `#installation`.
2. Hero describing Logn as a lightweight Go logging library.
3. Feature strip for daily/type logs, conditional warning/error Telegram
   reporting, monthly zipping triggered by the documented command/configuration,
   and optional old-directory cleanup.
4. Installation section with `go get -u github.com/ibnumardini/logn`.
5. Configuration section with the exact JSON example currently in `README.md`
   (including its obviously redacted token/chat-id examples) and explanations
   for `logn_is_active`, `log`, `tg`, and `zip` fields. Do not introduce real
   credentials or alter the example into a working secret.
6. Usage section with the documented `InfoLog`, `WarningLog`, and `ErrorLog`
   examples.
7. Cron zip section with `logn.CronZip(1)` and the `./app logn_zip_run`
   invocation.
8. Contributing and MIT license footer links targeting the existing
   `CONTRIBUTING.md` and `LICENSE.md` files.

Content will be grounded in the current README and source behavior; the page
will not claim features that are not documented by the project.

Required section IDs and navigation destinations are `#overview`,
`#installation`, `#configuration`, `#usage`, `#cron`, and `#contributing`.
`#overview` is the first section and the default active section. Active tracking
uses the section whose top edge is nearest to, but not below, the 35% viewport
line; if none qualifies, use the first section below that line. Ties are
resolved by document order. Both desktop and mobile links for the same section
receive the same `aria-current="location"` value. Scrolling does not rewrite
the URL; only an activated anchor navigation changes the hash.

## Alpine.js behavior

Use a small Alpine state object to provide:

- Responsive mobile navigation open/close behavior.
- Active documentation section tracking while scrolling.
- Copy-to-clipboard controls for install, configuration, and code snippets.
- A brief success state after copying.
- An explicit fallback message when `navigator.clipboard` is unavailable,
  the page is not in a secure context, or the promise rejects access. The
  message must tell the user to select/copy the visible snippet manually;
  repeated attempts must clear any previous success/error state before reporting
  the new result.

The desktop/normal-flow navigation is present in plain HTML. Mobile links are
visible in the default HTML/CSS state at all breakpoints. Do not use `x-cloak`
or CSS that hides the menu before Alpine initializes. Alpine enhances the same
menu by adding the `hidden` state after initialization on narrow viewports,
with a button using
`aria-controls="site-menu"` and `aria-expanded`; opening removes `hidden`,
Escape closes the menu, and focus returns to the toggle. If Alpine fails to
load or only partially initializes, no script hides the links and the visible
list remains usable. On initial load, active-section state is derived from the
current hash when present, otherwise the first documentation section is active.
The active link has `aria-current="location"` and is updated when the observed
section crosses a 35% viewport threshold; navigation updates the hash without
preventing normal anchor behavior. Without JavaScript, the first section may
retain a static active style. Copy controls are real buttons but are inert
without JavaScript and never claim success. There is no backend, bundler, form
submission, or persistent client-side state.

## Visual and accessibility direction

Use a near-black terminal-inspired palette with warm off-white text, amber
highlights, muted green status accents, subtle grid/noise treatment, and
monospace code surfaces. The layout must remain readable on mobile and use a
wider two-column documentation layout on desktop.

Include semantic headings, visible keyboard focus states, and accessible labels
for navigation and copy controls. Target WCAG AA contrast (4.5:1 for normal
text and 3:1 for large text), measured with a contrast checker against the
final computed colors. The mobile menu uses a button with
`aria-expanded`/`aria-controls` and a `nav` landmark containing the menu;
keyboard users can open it, tab through links, press Escape to close it, and
return focus to the toggle. Copy status must be announced through an
`aria-live="polite"` region. Success text is “Copied to clipboard.”; fallback
text is “Clipboard unavailable. Select and copy the snippet manually.” The
status clears before each attempt and after 3 seconds. Handle both synchronous
exceptions and rejected promises. Disable nonessential transitions under
`prefers-reduced-motion: reduce`.

## Validation

Validate the static file structure, then open the page locally in a browser to
check the following acceptance criteria:

- Desktop and mobile layout.
- Anchor navigation and active-section behavior.
- Mobile menu open/close behavior.
- Copy success and fallback behavior when clipboard access is denied.
- Direct loads with a section hash.
- No horizontal overflow at a narrow mobile viewport.
- Keyboard-only navigation, focus visibility, menu ARIA state, and live copy
  announcements.
- Readable content and navigation when JavaScript is disabled or the pinned Alpine CDN
  is blocked.
- WCAG AA contrast for the chosen text colors.

The implementation must be checked in Chromium 140 or newer at desktop (1440px
wide) and mobile (375px wide), both from a local HTTP server started with
`python3 -m http.server 8000` and with JavaScript disabled. Use a static HTML
parser or equivalent check to
confirm every required section/id, `aria-current`/menu attributes, and exact
external link exists. Block the Alpine CDN in browser devtools to verify the
plain-navigation invariant, and force a clipboard rejection in devtools to
verify the fallback text. Measure final colors with a WCAG contrast checker.
A validation run passes only when all listed checks succeed and no horizontal
overflow is present.

## Content verification checklist

- Feature copy states Telegram reporting is for warning/error logs when enabled.
- Zip copy states configuration enables the behavior and
  `./app logn_zip_run` triggers the documented cron path.
- Usage signatures match `InfoLog`, `WarningLog`, `ErrorLog`, and `CronZip` in
  the current Go source.
- Configuration keys and sample values match the current README without
  exposing credentials.

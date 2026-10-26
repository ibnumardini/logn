# Logn Static Docs Landing Page Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a single-file, terminal-inspired developer documentation landing page for Logn using static HTML, inline CSS, and Alpine.js CDN enhancements.

**Architecture:** Add one root `index.html` containing semantic markup, all page styles, and one Alpine component. The HTML and CSS provide the complete readable experience without JavaScript; Alpine progressively enhances the mobile menu, active section tracking, and clipboard actions.

**Tech Stack:** HTML5, CSS3, Alpine.js 3 via pinned jsDelivr CDN, browser Clipboard API.

---

## File map

- Create: `index.html` — the complete static page, including documentation content, responsive styles, accessibility semantics, and Alpine state/actions.
- Read-only source: `README.md`, `logn.go`, `config.go`, `logn_config_sample.json` — authoritative content and API/configuration references.
- Validate: `index.html` with a local HTTP server, Chromium, JavaScript-disabled mode, blocked CDN mode, and forced clipboard failure.

## Chunk 1: Static page structure and content

### Task 1: Create semantic documentation markup

**Files:**
- Create: `index.html`
- Read: `README.md`, `logn.go`, `config.go`, `logn_config_sample.json`

- [ ] **Step 1: Add the document shell and metadata**

  Create `<!doctype html>`, language metadata, viewport metadata, a descriptive title, and a short meta description. Add the exact pinned Alpine.js CDN script `https://cdn.jsdelivr.net/npm/alpinejs@3.14.9/dist/cdn.min.js` with `defer`; do not use `x-cloak` because the menu must remain visible if the CDN is blocked.

- [ ] **Step 2: Add the accessible header and navigation**

  Add Logn branding, a GitHub action linking to `https://github.com/ibnumardini/logn`, an install action linking to `#installation`, and a `nav` containing links to `#overview`, `#installation`, `#configuration`, `#usage`, `#cron`, and `#contributing`. The controlled menu element must be `id="site-menu"`. Include a mobile toggle with `aria-controls="site-menu"` and `aria-expanded="false"`. Keep the menu links visible in the default CSS state, and render desktop and mobile links from the same section IDs so both can receive matching `aria-current="location"` values.

- [ ] **Step 3: Add the overview and feature content**

  Add the `#overview` hero and feature strip. Describe Logn as a Go logging library and accurately state daily/type log generation, optional warning/error Telegram reporting, configured monthly zipping triggered through `logn_zip_run`, and optional old-directory cleanup.

- [ ] **Step 4: Add installation and configuration content**

  Add `#installation` with the exact `go get -u github.com/ibnumardini/logn` command and a copy control. Add `#configuration` with the README JSON example, preserving its redacted credential-shaped values, and explain the `logn_is_active`, `log`, `tg`, and `zip` groups without adding undocumented behavior.

- [ ] **Step 5: Add usage, cron, and footer content**

  Add `#usage` with the documented `InfoLog`, `WarningLog`, and `ErrorLog` examples. Add `#cron` with `logn.CronZip(1)` and `./app logn_zip_run`. Add `#contributing` with links to the repository's `CONTRIBUTING.md` and `LICENSE.md` GitHub URLs.

- [ ] **Step 6: Verify static anchors and content**

  Run:

  ```bash
  python3 - <<'PY'
  from html.parser import HTMLParser
  from pathlib import Path
  html = Path("index.html").read_text()
  required_ids = {"overview", "installation", "configuration", "usage", "cron", "contributing", "site-menu"}
  required_urls = {
      "https://github.com/ibnumardini/logn",
      "https://github.com/ibnumardini/logn/blob/master/README.md",
      "https://github.com/ibnumardini/logn/blob/master/logn_config_sample.json",
      "https://github.com/ibnumardini/logn/blob/master/CONTRIBUTING.md",
      "https://github.com/ibnumardini/logn/blob/master/LICENSE.md",
  }
  assert all(f'id="{value}"' in html for value in required_ids)
  assert all(url in html for url in required_urls)
  assert all(value in html for value in ("#overview", "#installation", "#configuration", "#usage", "#cron", "#contributing"))
  assert 'aria-controls="site-menu"' in html and 'aria-expanded="false"' in html
  assert "InfoLog" in html and "WarningLog" in html and "ErrorLog" in html and "logn_zip_run" in html
  assert "https://cdn.jsdelivr.net/npm/alpinejs@3.14.9/dist/cdn.min.js" in html
  print("static checks passed")
  PY
  ```

  Before writing the markup, compare the JSON/configuration keys against `README.md` and `logn_config_sample.json`, and compare API/cron signatures against `logn.go`. Expected: the script prints `static checks passed`.

- [ ] **Step 7: Commit the static structure**

  ```bash
  git add index.html
  git commit -m "feat: add logn static docs page"
  ```

## Chunk 2: Visual system and responsive layout

### Task 2: Add the terminal-inspired visual system

**Files:**
- Modify: `index.html` — inline `<style>` block

- [ ] **Step 1: Define design tokens and base styles**

  Add near-black background, warm off-white text, amber accents, muted green status colors, border/shadow tokens, readable font stacks, consistent spacing, and visible `:focus-visible` styles. Keep normal text at WCAG AA contrast targets.

- [ ] **Step 2: Style the page regions**

  Style the sticky header, hero, feature cards, section headings, code surfaces, copy buttons, config explanations, footer, and active navigation state. Use pseudo-elements or gradients for subtle grid/noise treatment without external assets or fonts.

- [ ] **Step 3: Add responsive breakpoints**

  Make the page single-column and readable at 375px wide, then use an explicit two-column documentation layout at desktop widths. Ensure long code lines wrap or scroll inside their own code surface and never cause page-level horizontal overflow.

- [ ] **Step 4: Add reduced-motion behavior**

  Under `@media (prefers-reduced-motion: reduce)`, remove nonessential transitions and scroll effects while retaining usable focus and active-state changes.

- [ ] **Step 5: Validate layout manually**

  Start:

  ```bash
  python3 -m http.server 8000
  ```

  Open `http://127.0.0.1:8000/` in Chromium 140+ at 1440px and 375px widths. Expected: two documentation columns at 1440px, one column at 375px, code overflow contained inside code surfaces, no page-level horizontal overflow, readable contrast, visible focus/active styles, all sections reachable, and no nonessential motion with reduced motion enabled.

- [ ] **Step 6: Commit visual styles**

  ```bash
  git add index.html
  git commit -m "style: add terminal docs layout"
  ```

## Chunk 3: Alpine progressive enhancements

### Task 3: Implement menu, active navigation, and clipboard behavior

**Files:**
- Modify: `index.html` — Alpine component markup and inline script/configuration

- [ ] **Step 1: Add Alpine state and menu enhancement**

  Define one Alpine state object with `menuOpen`, `activeSection`, `copyState`, and section references. Use the same `768px` breakpoint as CSS; at widths below it, Alpine hides `#site-menu` only after initialization, while at or above it the menu stays visible. Bind `aria-expanded`, `aria-current`, and the menu `hidden` state. Implement Escape-to-close and focus return to the toggle. If Alpine fails or initialization is interrupted, no script may hide the default-visible menu; keep the menu markup free of `x-cloak`.

- [ ] **Step 2: Add deterministic active-section tracking**

  Initialize `activeSection` from the URL hash only when it names one of the six sections; otherwise default to `overview`. Observe the six required sections and choose the section whose top edge is nearest to but not below the 35% viewport line; if none qualifies, choose the first section below it, with document order breaking ties. Apply `aria-current="location"` to both desktop and mobile links for the active ID. Do not rewrite the URL while scrolling; only normal anchor activation changes it. Verify default and invalid hashes, both link sets, tie-breaking, upward scrolling, and URL preservation.

- [ ] **Step 3: Add copy controls with explicit failure handling**

  Add copy buttons for the install command, configuration JSON, and usage/cron code snippets. Clear prior state before each attempt. Handle unavailable API, insecure context, synchronous exceptions, and rejected promises with the exact fallback message `Clipboard unavailable. Select and copy the snippet manually.` On success, announce `Copied to clipboard.` through an `aria-live="polite"` region and clear the status after three seconds. Never show success when copying fails. Verify each failure path and repeated attempts.

- [ ] **Step 4: Verify progressive enhancement**

  In Chromium:

  - Disable JavaScript: expected, all content and links remain usable; copy buttons do nothing.
  - Block the pinned Alpine CDN: expected, menu links remain visible and usable.
  - Enable JavaScript: expected, narrow menu collapses only after Alpine initializes, Escape closes it, and focus returns to the toggle.
  - Interrupt Alpine after loading or force a menu initialization error: expected, default-visible menu links remain usable.
  - Force unavailable/insecure-context, synchronous-throw, and rejected-promise clipboard paths: expected, exact fallback message appears, no success state appears, and each new attempt clears the previous status.
  - Load `/#configuration`: expected, configuration is the initial active section.
  - Load `/` and an invalid hash: expected, `overview` is active; scroll in both directions: expected, deterministic threshold/tie behavior and matching desktop/mobile `aria-current` values without hash rewrites.

- [ ] **Step 5: Commit Alpine behavior**

  ```bash
  git add index.html
  git commit -m "feat: add alpine docs interactions"
  ```

## Chunk 4: Final acceptance

### Task 4: Run the complete page checklist

**Files:**
- Validate: `index.html`

- [ ] **Step 1: Run static structure checks**

  Run the Python parser/assertion from Chunk 1 and additionally assert the exact URLs `https://github.com/ibnumardini/logn`, `https://github.com/ibnumardini/logn/blob/master/README.md`, `https://github.com/ibnumardini/logn/blob/master/logn_config_sample.json`, `https://github.com/ibnumardini/logn/blob/master/CONTRIBUTING.md`, `https://github.com/ibnumardini/logn/blob/master/LICENSE.md`, and `https://cdn.jsdelivr.net/npm/alpinejs@3.14.9/dist/cdn.min.js`. Expected: all assertions pass.

- [ ] **Step 2: Run browser acceptance checks**

  Using Chromium 140+ and `python3 -m http.server 8000`, verify desktop/mobile layout, direct hash loads, active-link updates while scrolling, menu keyboard behavior, copy success/failure states, JavaScript-disabled rendering, CDN-blocked rendering, and `document.documentElement.scrollWidth <= document.documentElement.clientWidth` at 375px. Place section tops around the 35% viewport line to verify the nearest-not-below rule, below-line fallback, document-order tie resolution, upward scrolling, matching desktop/mobile `aria-current`, and unchanged URL hash during scrolling. Force unavailable/insecure clipboard, synchronous throw, rejected promise, and repeated success/failure attempts. Expected: each listed behavior matches the spec, including exact copy messages and usable navigation with JavaScript/CDN unavailable.

- [ ] **Step 3: Check accessibility and contrast**

  Keyboard-tab through the page and menu, verify focus return after Escape, confirm matching `aria-current` values and live status updates (`Copied to clipboard.` / `Clipboard unavailable. Select and copy the snippet manually.`), confirm status clears after three seconds, and measure final computed colors with a WCAG contrast checker. Expected: 4.5:1 for normal text and 3:1 for large text.

- [ ] **Step 4: Review the final diff**

  ```bash
  git --no-pager diff 12b4f29..HEAD --stat
  git status --short
  ```

  Expected: the diff contains the intended root `index.html` plus the committed design/plan documents, and the worktree is clean. Do not modify the design or plan documents during implementation unless requirements change.

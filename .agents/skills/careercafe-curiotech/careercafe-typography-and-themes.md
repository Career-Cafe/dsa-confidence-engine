---
name: careercafe-typography-and-themes
scope: codebase-curiotech-careercafe
description: >-
  CurioTech and CareerCafe typography hierarchy and 3-state theme engineering: IBM Plex Sans/Mono, System/Light/Dark selector, no-flash theme resolution, and accessibility QA.
---

# CareerCafe Typography & Theme Engineering Standard

> [!IMPORTANT]
> **PROJECT-SPECIFIC SCOPE**: This skill is strictly specific to **CurioTech** and **CareerCafe** projects (`CurioTech-CareerCafe`). It governs font pairing, typographic scales, and the three-state theme architecture.

---

## 1. Typography Core Foundation

- **Product / UI Font**: `IBM Plex Sans` (Humanist, technical, highly readable for long sessions).
- **Code / Data / Timers**: `IBM Plex Mono` (Syntax-accurate, fixed-width numerals).
- **Strict Prohibition**: Zero third product typefaces in V1. No handwritten, geometric-display, or novelty fonts.
- **Font Weight Cap**: No font weight above `600` (Semi-Bold). Boldness is achieved via hierarchy and contrast, not 700/800/900 bloat.
- **Ligatures OFF**: IBM Plex Mono must have font ligatures disabled by default (`font-variant-ligatures: none;`) so that relational operators (`!=`, `<=`, `===`, `=>`) remain literal and unambiguous.
- **Tabular Numerals**: Every timer, counter, percentage, or dynamically changing numerical display must enforce tabular numbers (`font-variant-numeric: tabular-nums;`).

---

## 2. Complete Typographic Scales

### 2.1 Marketing Scale (Public Pages & Landing Experience)

| Role | Desktop (Size / Line-Height) | Mobile (Size / Line-Height) | Weight & Tracking | Usage |
|---|---:|---:|---|---|
| **Hero H1** | 60px / 63px | 32px / 38px | 600 · `-0.03em` | Primary value proposition headline. |
| **H2** | 40px / 48px | 26px / 32px | 600 · `-0.024em` | Major section headers. |
| **H3** | 26px / 34px | 21px / 28px | 600 · `-0.012em` | Sub-section & feature card headers. |
| **Lead Paragraph** | 20px / 32px | 18px / 28px | 400 · normal | Hero introductory and problem framing. |
| **Body** | 17px / 28px | 16px / 26px | 400 · normal | Descriptive copy, paragraphs, list items. |
| **Button** | 16px / 20px | 16px / 20px | 600 · normal | Action buttons and navigation links. |
| **Metadata** | 13px / 18px | 13px / 18px | 500 · normal | Category chips, dates, timestamps. |

### 2.2 Product Scale (App Screens, Question Bank, IDE, Interview Mode)

| Role | Size / Line-Height | Weight | Strict Measure | Usage |
|---|---:|---|---|---|
| **Page Title H1** | 30px / 38px | 600 | Full width | Top-level dashboard & view titles. |
| **Section H2** | 24px / 32px | 600 | Full width | Major in-app view groupings. |
| **Subsection H3** | 19px / 28px | 600 | Full width | Form headers, modal headers. |
| **UI Body** | 15px / 24px | 400 | Bounded | Forms, panels, settings, controls. |
| **Long-Form Body** | 17px / 30px | 450 | **Max 680px** | Question Bank problems, case studies, solutions. |
| **Reading Toggle** | 19px / 34px | 450 | **Max 680px** | Optional large reading mode for long sessions. |
| **Card Title** | 16px / 22px | 600 | Card container | Compact question & drill titles. |
| **Card Body** | 14px / 21px | 400 | Card container | Compact support copy and hints. |
| **Metadata / Badges** | 13px / 18px | 500 | Compact | Difficulty tags, completion stats. |
| **Code Editor** | 14px / 21.7px | 400 Mono | Editor container | SQL/Python input area; ligatures off. |
| **Timers / Counts** | 20px–32px | 500 Mono | Tabular | Interview session countdowns, clock. |

---

## 3. Three-State Theme Architecture (System / Light / Dark)

The CareerCafe theme model treats Light and Dark as first-class expressions of the same brand, not disparate visual identities.

```mermaid
flowchart TD
    Init[Browser Initial Request / Page Request] --> ResolveTheme[Resolve Theme in <head> before Hydration]
    ResolveTheme --> CheckStored{Manual Preference in LocalStorage?}
    CheckStored -- Yes ('light' or 'dark') --> ApplyManual[Apply html class='light' or 'dark']
    CheckStored -- No ('system' or unset) --> DetectOS[Read window.matchMedia '(prefers-color-scheme: dark)']
    DetectOS --> ApplyOS[Apply OS Resolved Theme]
    ApplyManual --> Render[Zero-Flash First Paint]
    ApplyOS --> Render
    Render --> Listener[Listen for OS Theme Changes if System Selected]
```

### Engineering Requirements:
1. **Three-State Selector**: Navigation contains a subtle theme control supporting `System`, `Light`, and `Dark`. Never use a binary-only sun/moon toggle.
2. **Zero-Flash Hydration**: The active theme MUST be resolved via inline blocking script before the first render/paint. Light-theme flashes on dark preference are classified as critical bugs.
3. **Landing Page Default**: Designed light-first conceptually, but the public website ALWAYS respects the visitor's resolved active system theme.
4. **Accessible Labels**: Theme controls must expose accessible labels matching their state (e.g. `Theme: System (Dark)`).

---

## 4. Responsive & Accessibility QA

### Viewport Gateways:
- **360px–390px (Compact Mobile)**: Hero CTAs stack vertically; touch chips maintain 44px+ height; no horizontal scrollbars; readable typography.
- **768px–1024px (Tablet / Narrow Desktop)**: Fluid card wrapping; screenshot text remains legible; side panels collapse into bottom drawers.
- **1366px × 768px (Standard Laptop)**: Primary hero CTA immediately visible above the fold; screenshot text legible at 1x scale without zooming.

### Accessibility Invariants:
- Minimum interactive tap target: `44 × 44px` (prefer `48px` for primary actions).
- Visible keyboard focus ring: 2px solid ring (`#17252B` light / `#FFFFFF` dark) with 2px offset.
- Contrast ratio: Minimum `4.5:1` for standard text; `3.0:1` for large text and interactive boundaries.

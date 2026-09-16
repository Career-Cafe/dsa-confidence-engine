---
name: careercafe-landing-page-spec
scope: codebase-curiotech-careercafe
description: >-
  CurioTech and CareerCafe landing page specifications: locked 12-section hierarchy, visual spacing rhythm, hero proof rules, and real product screenshot standards.
---

# CareerCafe Landing Page Architecture Specification

> [!IMPORTANT]
> **PROJECT-SPECIFIC SCOPE**: This skill is strictly specific to **CurioTech** and **CareerCafe** projects (`CurioTech-CareerCafe`). It enforces the frozen v3.2 landing page structure, section-by-section requirements, and spacing rhythm.

---

## 1. Locked 12-Section Architecture

The landing page structure is frozen. Contributors and AI agents must NOT add additional sections; all efforts focus on implementation fidelity, screenshot clarity, and proof quality:

| # | Section Name | Final Visual & Functional Treatment |
|---:|---|---|
| **1** | **Navigation** | Sticky compact header: Programmes / Practice / Company Prep / How It Works / Sign In / `Start Practising Free` + secondary System/Light/Dark selector. |
| **2** | **Hero + Compact Proof** | Eyebrow + `Knowledge ≠ Performance` H1 + product definition + 2 CTAs + max 3 real student cards. |
| **3** | **Choose Your Programme** | Foundation + Placement Pass; Sprint as compact route. 2 cards; 3 bullets max; Placement Pass stronger. |
| **4** | **How CareerCafe Prepares You** | De-carded 3-step journey: Practice $\rightarrow$ Interview $\rightarrow$ Human Validation. Whitespace + simple connector. |
| **5** | **Real Product Experience** | Large real Practice + Interview screenshots. Page visual peak; ~16:10 aspect ratio. |
| **6** | **Analyst Directions** | 4 light border-only cards using locked taxonomy; no heavy shadows; 3 chips max. |
| **7** | **AI for Analysts** | Compact inset panel in Charcoal or Soft Sage; 4 concepts + 1 CTA; no syllabus. |
| **8** | **Company Preparation** | Two-column desktop; company chips + selected panel; coverage/last reviewed visible. |
| **9** | **Human Validation & Outcomes**| Editorial trust split; max 3 approved, verified student stories. |
| **10** | **Campus Analyst Challenge** | Compact B2B section; must not dominate the primary student conversion funnel. |
| **11** | **Final Student CTA** | Full-width dark Charcoal band (`#17252B`) + single orange action (`#BC4A1E`) + reassurance. |
| **12** | **Footer** | Three compact link groups maximum; real links only. |

---

## 2. Hero & Product Proof Rules

### Copy & Structure Constraints
- **Eyebrow**: `Brewing Future Analysts` (Single Charcoal or subtle secondary text, no colored badge).
- **H1**: `Knowledge ≠ Performance` — strictly rendered in one Charcoal color (`#17252B`). Never colorize individual words.
- **Problem Statement**: *"The gap between knowing and performing is where most candidates struggle."*
- **Product Definition**: *"Realistic practice + interview-style simulations + readiness validation."*
- **Primary CTA**: `Start Practising Free` (Orange filled `#BC4A1E`, 48px height, 6px radius).
- **Secondary CTA**: `Explore Programmes` (Subtle control border, ghost style).
- **Capability Strip**: `SQL & Python · Cases & Guesstimates · Interview Questions · Company Prep`.
- **Proof Cards**: Maximum 3 student cards (1 primary + 2 supporting). Never use tiny fragmented quotes.
- **Company Logos**: Do not imply legal partnerships; always include a coverage disclaimer (*"Curated interview questions based on public candidate reports"*).

---

## 3. Real Product Experience Screenshot Invariants

> [!CAUTION]
> **ZERO FAKE DASHBOARDS OR PLACEHOLDER MOCKUPS**:
> This section represents the visual peak of the landing page. It is strictly blocked until representative, authentic Practice Mode and Interview Mode screens exist.

- **Container Constraint**: Up to `1180px` centered width.
- **Aspect Ratio**: Approximately `16:10`.
- **Border & Radius**: `12px` rounded corners with `1px` subtle border (`#E7E4DD` light / `#40545D` dark).
- **Shadow**: Marketing proof shadow token (`0 8px 24px rgba(23,37,43,0.08)` in light theme; `rgba(0,0,0,.32)` in dark).
- **Theme Match**: Screenshots must dynamically match the active theme, or be explicitly captioned as a cross-theme example.
- **Desktop Legibility**: Question text and primary controls must be legible at 1x scale on standard 1366×768 screens.
- **Mobile Art Direction**: On mobile viewports, crop and focus on the core editor/timer interaction. Never shrink desktop UI to unreadable micro-dimensions.

---

## 4. Spacing Tokens & Visual Rhythm

| Spacing Token | Desktop | Mobile | Purpose |
|---|---:|---:|---|
| **Major Section Padding** | 96px | 64px | Vertical spacing between primary landing sections. |
| **Eyebrow $\rightarrow$ Heading** | 16px | 12px | Spacing above section titles. |
| **Heading $\rightarrow$ Body** | 16px | 12px | Spacing between title and lead text. |
| **Body $\rightarrow$ Major Content** | 40px | 32px | Spacing between intro text and interactive components. |
| **Grid / Card Gap** | 24px | 16px | Gap between cards in grids. |
| **Card Padding** | 32px | 24px | Interior padding within cards. |
| **CTA Gap** | 12px | 12px | Spacing between primary and secondary buttons. |
| **Page Gutter** | 48px | 24px | Horizontal viewport margins. |
| **Max Marketing Width** | 1200px | 100% | Maximum page layout constraint. |

---

## 5. Visual Rhythm & Layout Rules

1. **Break Card Repetition**: Do not build the landing page as an unending stack of identical 3-card grids. Alternate with editorial text splits, large screenshots, lightweight border cards, and inset panels.
2. **Left-Aligned Bias**: Prefer clean left-aligned headings and descriptions over centering every paragraph. Centered text is reserved for the Hero and Final CTA.
3. **Restrained Insets**: Use the inset Charcoal panel (`#17252B`) sparingly for high-impact technical features (e.g. AI for Analysts).
4. **Final Conversion Anchor**: Only the Final CTA section receives full-width edge-to-edge dark treatment on the light canvas.

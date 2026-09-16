# CareerCafe
## FINAL COMBINED VISUAL DESIGN SYSTEM & LANDING PAGE HANDOFF
**v3.2 - FROZEN SOURCE OF TRUTH - Light + Dark**

## Core identity

Orange + Sage + teal-cast Charcoal. Minimal cards. Strong real product screenshots. Restrained shadows. Large whitespace. Premium, readable, technical - not playful edtech.

### Decision / Frozen choice

| Decision | Frozen choice |
|---|---|
| Product font | IBM Plex Sans |
| Code / data / timers | IBM Plex Mono |
| Theme model | System / Light / Dark |
| Light canvas | `#F8F7F4` warm off-white |
| Dark canvas | `#10181C` deep teal-charcoal |
| Primary action | Orange |
| Practice / progress state | Sage |
| Structure / trust | Teal-cast Charcoal |
| Overall feel | Calm, editorial, technical, evidence-led |

**Visual-design source of truth.**

---

# 1. Brand and Design Principles - Frozen

| Principle | Final rule |
|---|---|
| Premium | Premium comes from restraint, typography, spacing, content clarity and real product UI - not gradients or excessive decoration. |
| Readable | CareerCafe is content-heavy. Long reading, coding and interview sessions must remain comfortable for 30-60+ minutes. |
| Technical | The product should feel like a serious interview-performance tool, not a coaching brochure or generic course marketplace. |
| Colour semantics | Orange = action. Sage = state. Charcoal = structure/trust. |
| Minimal cards | Use cards only when grouping or interaction requires them. Do not wrap every section in rounded containers. |
| Real proof | Product screenshots and verified outcomes carry trust. Decorative marketing graphics do not replace evidence. |
| Whitespace | Use generous section rhythm and controlled reading width. Do not fill empty space with badges/icons/copy. |
| One system | Marketing, Question Bank, playgrounds and Interview Mode share one system and differ only in density and visual temperature. |
| Themes | Light and Dark are first-class expressions of the same brand, not separate visual identities. |

### Freeze rule

Re-open these decisions only if user testing reveals a concrete readability, accessibility or conversion problem.

---

# 2. Typography System - Frozen

Product/UI typography: **IBM Plex Sans**. Code/data/timers: **IBM Plex Mono**. No third product typeface in V1.

## 2.1 Marketing scale

| Role | Desktop | Mobile | Weight / tracking |
|---|---:|---:|---|
| Hero H1 | 60 / 63 | 32 / 38 | 600 - `-0.03em` |
| H2 | 40 / 48 | 26 / 32 | 600 - `-0.024em` |
| H3 | 26 / 34 | 21 / 28 | 600 - `-0.012em` |
| Lead paragraph | 20 / 32 | 18 / 28 | 400 |
| Body | 17 / 28 | 16 / 26 | 400 |
| Button | 16 / 20 | 16 / 20 | 600 |
| Metadata | 13 / 18 | 13 / 18 | 500 |

## 2.2 Product scale

| Role | Size / line-height | Weight | Use |
|---|---:|---:|---|
| Page title H1 | 30 / 38 | 600 | App screens |
| Section H2 | 24 / 32 | 600 | App sections |
| Subsection H3 | 19 / 28 | 600 | Subsections |
| UI body | 15 / 24 | 400 | Forms, panels, controls |
| Long-form body | 17 / 30 | 450 | Question Bank, cases, explanations |
| Reading toggle | 19 / 34 | 450 | Optional larger reading mode |
| Card title | 16 / 22 | 600 | Compact UI cards |
| Card body | 14 / 21 | 400 | Short support copy |
| Metadata | 13 / 18 | 500 | Timestamps/counts |
| Editor | 14 / 21.7 | 400 Mono | SQL/Python editor; ligatures off |
| Timer / changing numerals | 20-32 | 500 Mono | Tabular numerals |

- Long-form reading measure: max 680px.
- No handwritten, geometric-display or decorative product font.
- No font weight above 600 in V1.
- IBM Plex Mono ligatures OFF by default so `!=`, `<=` and other syntax remain literal.
- Use tabular numerals anywhere numbers change.

---

# 3. Theme Behaviour - Frozen

| System decision | Frozen decision |
|---|---|
| Theme model | System / Light / Dark |
| Default | New or unset preference follows the operating-system theme |
| Manual override | Light or Dark persists until changed |
| Landing page | Designed light-first, but always respects the active theme |
| Brand identity | Orange = action - Sage = state - Charcoal = structure in both themes |

## 3.1 Theme selector

- Use a subtle theme icon/control in navigation or account UI; it opens System / Light / Dark.
- Do not use a binary sun/moon toggle as the only control.
- Accessible label reflects the current state, e.g. `Theme: System`.
- Selected state is exposed to assistive technology and is keyboard operable.
- Resolve the active theme before hydration/paint to avoid a light-theme flash.
- When preference = System, respond to operating-system theme changes without refresh.
- The theme control must not compete visually with Sign In or Start Practising Free.

### Landing default clarified

Light-first describes design priority, not a forced runtime default. The public site always respects the visitor's resolved active theme.

---

# 4. Light / Dark Surface and Text Tokens - Frozen

| Token | Light | Dark |
|---|---|---|
| Canvas | `#F8F7F4` | `#10181C` |
| Primary surface | `#FFFFFF` | `#17252B` |
| Elevated surface | `#FFFFFF` | `#1D2D33` |
| Sunken surface | `#EFEDE7` | `#0D1518` |
| Practice / sage surface | `#EEF3E8` | `#1B2A1E` |
| Static code surface | `#F2F4F0` | `#1A282D` |
| Input background | `#FFFFFF` | `#132025` |
| Heading / strong ink | `#17252B` | `#F2F4F3` |
| UI body | `#2C3B42` | `#D9E0DD` |
| Long-form reading | `#33444B` | `#C9D2CF` |
| Secondary text | `#4A5D64` | `#AEB6B9` |
| Metadata | `#5B6E75` | `#95A3A7` |
| Subtle border | `#E7E4DD` | `#40545D` |
| Control border | `#767E86` | `#60747B` |
| Sage graphical | `#7F9B6D` | `#9FB58D` |
| Sage text | `#4F6640` | `#B7C9A8` |

### Charcoal identity

`#17252B` is deliberately teal-cast. Do not substitute a neutral grey-black in either theme; the teal-charcoal foundation is part of CareerCafe's identity.

---

# 5. Interaction, Border and Semantic Tokens - Frozen

| Token | Light | Dark | Rule |
|---|---|---|---|
| Primary action fill | `#BC4A1E` | `#BC4A1E` | White text; orange stays consistent |
| Primary hover | `#A83E17` | `#A83E17` | Hover darkens, never lightens |
| Focus ring | `#17252B` | `#FFFFFF` | 2px ring + 2px offset |
| Hovered row | `#F3F1EC` | `#1B2C31` | Subtle; no accent required |
| Selected row | `#FBEEE8` | `#233940` | Also use text/icon/state cue |
| Disabled bg | `#F3F1EC` | `#1A252A` | No hover |
| Disabled text | `#8A9AA0` | `#7F8D91` | Still identifiable |
| Disabled border | `#D9D5CC` | `#33474F` | Passive boundary |
| Overlay | `rgba(23,37,43,.28)` | `rgba(0,0,0,.68)` | Modal/sheet backdrop |
| Marketing shadow | `rgba(23,37,43,.08)` | `rgba(0,0,0,.32)` | Proof/screenshots only |
| Product UI shadow | `rgba(23,37,43,.08)` | `rgba(0,0,0,.35)` | Max `0 1px 3px` in product |

## 5.1 Border distinction

- Subtle border (`#E7E4DD` light / `#40545D` dark) = decorative divider or passive card boundary only.
- Control border (`#767E86` light / `#60747B` dark) = inputs, outline buttons and other interactive boundaries.
- Do not rely on the subtle border alone to communicate an interactive control.

## 5.2 Semantic states

| State | Light | Dark | Dark surface | Rule |
|---|---|---|---|---|
| Success | `#4F6640` | `#B7C9A8` | `#1B2A1E` | Pair with icon + label |
| Warning | `#8A5A00` | `#E8C47A` | `#2A2314` | Caution/timing/incomplete |
| Error | `#9E1C2B` | `#F29A9A` | `#321A1E` | Distinct from action orange |
| Info | `#2F5D7C` | `#8FC3E0` | `#172833` | Neutral information |

### Mode semantics

Orange = action and Sage = state in both themes. Practice and Interview are never identified by colour fill alone; every stage indicator carries a text label and/or distinct icon.

---

# 6. Component Geometry, Shadows and Motion

| Token | Frozen value | Rule |
|---|---|---|
| Card/input radius | 8px | Default product containers |
| Button radius | 6px | No oversized pill buttons |
| Product screenshot radius | 12px | Marketing/product proof only |
| Marketing/product proof shadow | `0 8px 24px rgba(23,37,43,0.08)` | Light theme; dark uses `rgba(0,0,0,.32)` |
| Product UI shadow | `0 1px 3px rgba(23,37,43,0.08)` max | Dark uses `rgba(0,0,0,.35)` |
| Motion | `150ms ease-out` | Standard interaction changes |

- No gradients, glassmorphism, glowing badges, confetti, blobs or mascot illustrations.
- No more than two accent colours visible in one viewport.
- Icons use one stroke-based library with a consistent approximately 1.5px stroke.
- Sage base and decorative orange are shapes/surfaces, not normal body/link text colours.

---

# 7. One Brand, Four Surface Behaviours

| Surface | Light / default treatment | Dark treatment / rule |
|---|---|---|
| Marketing landing page | Warm, spacious, 1200px max, strong screenshots, orange CTA | Deep canvas with Charcoal-family surfaces; same hierarchy; no neon treatment |
| Question Bank / content | White 680px reading column on warm canvas; 17/30 reading type | Dark reading surface `#17252B` + softened `#C9D2CF`; no pure-white long paragraphs |
| SQL / Python playground | Dense controls; dark executable editor; monochrome chrome | Editor integrates with `#17252B/#1D2D33` panels; maintain panel separation |
| Interview Mode | Near-monochrome; distraction-free; one question/input/timer | No sage during active interview; no encouragement/success colour mid-session |

## 7.1 Playground editor palette

| Token | Hex |
|---|---|
| Editor background | `#132025` |
| Editor chrome | `#17252B` |
| Foreground | `#D4DCDA` |
| Keyword | `#B79BE8` |
| String | `#A9C295` |
| Number | `#E0B26A` |
| Function | `#7FB4E3` |
| Type | `#72C4B8` |
| Operator | `#E8825A` |
| Comment | `#7C8F94` |
| Error | `#E88070` |

---

# 8. Landing Page Structure - Frozen

| # | Section | Final treatment |
|---:|---|---|
| 1 | Navigation | Sticky compact header; Programmes / Practice / Company Prep / How It Works / Sign In / Start Practising Free + secondary theme selector |
| 2 | Hero + compact proof | Knowledge ≠ Performance + product definition + 2 CTAs + 3 real student cards max |
| 3 | Choose Your Programme | Foundation + Placement Pass; Sprint as compact route |
| 4 | How CareerCafe Prepares You | De-carded 3-step journey: Practice → Interview → Human Validation |
| 5 | Real Product Experience | Large real Practice + Interview screenshots; page visual peak |
| 6 | Analyst Directions | 4 lighter cards using locked taxonomy |
| 7 | AI for Analysts | Compact inset panel; not a course section |
| 8 | Company Preparation | Company chips + selected panel; Explore Preparation primary |
| 9 | Human Validation + Verified Outcomes | One trust region; max 3 real approved stories |
| 10 | Campus Analyst Challenge | Compact B2B section |
| 11 | Final Student CTA | Full-width dark Charcoal band + one orange CTA |
| 12 | Footer | Compact; real links only |

### Do not add more sections

The structure is complete. Improve implementation quality, proof and real product screenshots - not page length.

---

# 9. Hero and Product Proof - Frozen

## 9.1 Hero

| Element | Rule |
|---|---|
| Eyebrow | `Brewing Future Analysts` |
| H1 | `Knowledge ≠ Performance` - one Charcoal colour; no coloured word |
| Problem line | The gap between knowing and performing is where most candidates struggle. |
| Product definition | Realistic practice + interview-style simulations + readiness validation. |
| Primary CTA | `Start Practising Free` - orange filled |
| Secondary CTA | `Explore Programmes` - outline/ghost |
| Capability row | SQL & Python · Cases & Guesstimates · Interview Questions · Company Prep |
| Proof cards | 3 students max; one primary + two supporting; no tiny quotes |
| Company logos | Do not imply partnerships; use coverage disclaimer |

## 9.2 Real Product Experience screenshot rules

| Rule | Final requirement |
|---|---|
| Container | Up to 1180px |
| Aspect ratio | Approx. 16:10 |
| Radius | 12px |
| Border | 1px subtle border |
| Shadow | Marketing proof shadow token |
| Theme match | Where possible, screenshots match the active landing-page theme |
| Cross-theme example | Allowed only when clearly framed/captioned as Light mode or Dark mode example |
| Desktop legibility | Question text + primary interaction readable at 1x on 1366px |
| Mobile | Art-directed/cropped screenshots in both themes; never shrink full desktop UI |
| Asset QA | Prepare both light and dark screenshot variants once representative screens exist |

### Dependency

Do not use fake dashboards or placeholder mockups. This section is blocked until representative Practice and Interview screens exist.

---

# 10. Section Treatment, Spacing and Visual Rhythm

| Section | Final visual rule |
|---|---|
| Programmes | Two cards; 3 bullets max; Placement Pass stronger, Foundation quieter. |
| How It Works | No cards. Three numbered steps + whitespace + simple connector. |
| Directions | Light border-only cards; no heavy shadows; 3 chips max. |
| AI for Analysts | Inset Charcoal or Soft Sage panel; 4 concepts + 1 CTA; no syllabus. |
| Company Prep | Two-column desktop; coverage/last reviewed visible; Explore Preparation primary. |
| Human + Outcomes | Editorial trust split; only real proof. |
| Campus Challenge | Compact; B2B section must not dominate student funnel. |
| Final CTA | Only full-width dark conversion band; one orange action + reassurance. |
| Footer | Three compact groups max; no dead links. |

## Spacing token

| Token | Desktop | Mobile |
|---|---:|---:|
| Major section padding | 96px | 64px |
| Eyebrow → heading | 16px | 12px |
| Heading → body | 16px | 12px |
| Body → major content | 40px | 32px |
| Grid/card gap | 24px | 16px |
| Card padding | 32px | 24px |
| CTA gap | 12px | 12px |
| Page gutter | 48px | 24px |
| Max marketing width | 1200px | 100% |

- Break card repetition with editorial text, product screenshots, light grids, selected-company panels and trust splits.
- Use more left-aligned sections; do not centre every heading and paragraph.
- Reserve full-width dark treatment for the final CTA; dark mode itself still follows the same hierarchy.

---

# 11. Responsive, Accessibility and Theme QA

| Viewport | Required checks |
|---|---|
| 360-390px | Stack hero CTAs; 44px+ chips; no unreadable screenshots; no horizontal breaks |
| 768-1024px | Screenshot legibility, card wrapping, company-panel transition |
| 1366×768 | Hero CTA visible; screenshots readable; no oversized first screen |

## 11.1 Theme QA matrix

| Area | Required checks |
|---|---|
| Landing + nav | System/Light/Dark selector; no theme flash; CTA contrast; proof cards; footer; sticky nav |
| Question Bank | Long-form comfort; code blocks; links; metadata; selected filters; focus states |
| SQL/Python playground | Editor, tabs, results grid, Run/Submit, console errors, input/control boundaries |
| Interview Mode | Timer, question stem, editor/input, expiry state, focus visibility, no mid-session sage feedback |
| Company Prep | Chips, selected company panel, CTA hierarchy, disclaimer |
| Screenshots | Active-theme match or explicit cross-theme framing |

- Minimum tap target 44×44px; prefer 48px primary buttons.
- Visible keyboard focus on every interactive control in both themes.
- Colour is never the only indicator of mode or status.
- No light-theme flash during first paint, auth transition or route navigation.
- Subtle borders are never the sole boundary of an interactive control.

---

# 12. Implementation Priority - Frozen

## Do today

Remove borrowed/placeholder testimonials or proof and remove dead public links. Do not leave competitor copy or false trust elements on a live preview.

## 12.1

- ☐ Apply IBM Plex Sans + IBM Plex Mono product tokens.
- ☐ Apply Light + Dark surface/text/state tokens from this handoff.
- ☐ Implement the three-state Theme selector: System / Light / Dark.
- ☐ Resolve theme before first paint; persist manual choice; System tracks OS.
- ☐ Apply Orange = action / Sage = state / Charcoal = structure.
- ☐ Apply typography and spacing tokens exactly.
- ☐ Reduce hero proof to 3 student cards.
- ☐ Simplify Programme cards.
- ☐ De-card How CareerCafe Works.
- ☐ Lock direction labels and fix duplicate tags/titles.
- ☐ Add compact AI for Analysts panel.
- ☐ Fix Company Prep CTA hierarchy.
- ☐ Keep Campus Challenge compact.
- ☐ Add final dark student CTA.
- ☐ Simplify footer.
- ☐ Run responsive + keyboard + contrast + full theme QA.

## 12.2 Blocked on representative product UI

- ☐ Capture/finalise Practice Mode light + dark screenshots.
- ☐ Capture/finalise Interview Mode light + dark screenshots.
- ☐ Build Real Product Experience section from those real screens.

---

# 13. Explicit Visual Non-Goals

- No extra display font or serif in V1.
- No Poppins/Montserrat/Nunito-style generic edtech direction.
- No gradient-mesh/aurora hero.
- No glassmorphism.
- No heavy shadow system.
- No mascot/blob/3D emoji illustration language.
- No large orange or sage full-screen sections.
- No pill-everything UI; pills are for tags/chips only.
- No giant analytics dashboard on landing page or student Home.
- No fake product screenshots.
- No repeated all-caps eyebrow above every section.
- No coloured word inside every headline.
- No forced-light public site. Light-first design must still respect the active theme.
- No binary-only theme toggle; retain System / Light / Dark.

### Premium test

If a visual choice makes CareerCafe look more like a generic edtech/coaching site than a serious interview-performance product, do not use it.

---

# 14. Final Visual Acceptance Checklist

- ☐ First screen explains what CareerCafe is without requiring inference.
- ☐ IBM Plex Sans / Mono are applied consistently to the product.
- ☐ Handoff itself uses Calibri 13/12/11 as requested.
- ☐ Light mode uses warm off-white canvas + white content surfaces.
- ☐ Dark mode uses deep teal-charcoal canvas + softened light text.
- ☐ Theme behaviour is System / Light / Dark with persisted manual override and no flash.
- ☐ Orange is used for primary action; Sage for state/progress; Charcoal for structure.
- ☐ Practice and Interview remain distinguishable without relying on colour alone.
- ☐ No normal text uses inaccessible base Sage or decorative orange.
- ☐ Long-form content is limited to ~680px with 17/30 reading typography.
- ☐ Landing page does not feel like identical card grids.
- ☐ Real Practice + Interview UI become the main proof once available.
- ☐ Shadows remain restrained and limited to meaningful elevation/proof.
- ☐ Whitespace feels deliberate.
- ☐ Both themes pass 360/390px, tablet and 1366×768 QA.
- ☐ Screenshots match active theme or are intentionally labelled cross-theme examples.
- ☐ No placeholder testimonials, false proof or dead CTA remains.

## Final freeze

CareerCafe now has one combined visual source of truth: Orange + Sage + teal-cast Charcoal, IBM Plex Sans + Mono, warm light mode, focused dark mode, minimal cards, strong real product screenshots, restrained shadows and large whitespace. Re-open only against concrete user-testing evidence.

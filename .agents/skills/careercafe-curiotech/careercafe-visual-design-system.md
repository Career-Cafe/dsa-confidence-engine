---
name: careercafe-visual-design-system
scope: codebase-curiotech-careercafe
description: >-
  CurioTech and CareerCafe visual design system standard: color tokens (Orange action, Sage state, teal-cast Charcoal structure), surface tokens, component geometry, shadows, motion, and visual invariants.
---

# CareerCafe Visual Design System Standard

> [!IMPORTANT]
> **PROJECT-SPECIFIC SCOPE**: This skill is strictly specific to **CurioTech** and **CareerCafe** projects (`CurioTech-CareerCafe`). It enforces the frozen v3.2 visual design system across marketing pages, Question Banks, playgrounds, and interview experiences.

---

## 1. Core Brand Identity & Palette

The CareerCafe visual brand is rooted in **Orange + Sage + teal-cast Charcoal**. It is designed to feel calm, editorial, technical, and evidence-led—deliberately avoiding playful, generic edtech aesthetics.

### Color Tokens & Semantic Assignments

| Semantic Role | Token Name | Light Value | Dark Value | Immutable Rule |
|---|---|---|---|---|
| **Primary Action** | `color-action` | `#BC4A1E` | `#BC4A1E` | White text on fill. Orange stays consistent across both themes. |
| **Primary Hover** | `color-action-hover`| `#A83E17` | `#A83E17` | Hover always darkens, never lightens. |
| **Practice / State** | `color-state-sage` | `#7F9B6D` (graphical)<br>`#4F6640` (text) | `#9FB58D` (graphical)<br>`#B7C9A8` (text) | Used exclusively for practice/progress states and badges. |
| **Structure / Trust**| `color-charcoal` | `#17252B` | `#F2F4F3` (ink)<br>`#17252B` (surface) | Deliberately teal-cast. NEVER substitute a neutral grey-black. |
| **Focus Ring** | `color-focus` | `#17252B` | `#FFFFFF` | 2px solid ring with 2px offset. |

### Color Semantics Directives
- **Orange = Action**: Reserved exclusively for primary interactive calls to action, submit buttons, and forward progression.
- **Sage = State / Progress**: Communicates learning status, practice progress, and positive validation. Sage is a shape/surface, NEVER normal body or link text.
- **Charcoal = Structure / Trust**: Deliberately teal-cast `#17252B`. Establishes page headers, structural dividers, dark conversion bands, and trust anchors.

---

## 2. Surface & Text Tokens

| Token | Light Value | Dark Value | Usage |
|---|---|---|---|
| **Canvas** | `#F8F7F4` | `#10181C` | Page base background (warm off-white / deep teal-charcoal). |
| **Primary Surface** | `#FFFFFF` | `#17252B` | Main cards, panels, content reading containers. |
| **Elevated Surface**| `#FFFFFF` | `#1D2D33` | Modals, flyouts, popovers, elevated cards. |
| **Sunken Surface** | `#EFEDE7` | `#0D1518` | Inset wells, code blocks, secondary sections. |
| **Practice / Sage** | `#EEF3E8` | `#1B2A1E` | Practice mode accents, completed card surfaces. |
| **Static Code** | `#F2F4F0` | `#1A282D` | Read-only syntax and code snippet containers. |
| **Input Background** | `#FFFFFF` | `#132025` | Form inputs, textareas, code editors. |
| **Strong Ink** | `#17252B` | `#F2F4F3` | Headings, bold emphasis, titles. |
| **UI Body** | `#2C3B42` | `#D9E0DD` | Form labels, control text, dashboard copy. |
| **Long-Form Reading**| `#33444B` | `#C9D2CF` | Question Bank content, case studies, explanations. |
| **Secondary Text** | `#4A5D64` | `#AEB6B9` | Subtitles, helper text, inactive options. |
| **Metadata** | `#5B6E75` | `#95A3A7` | Timestamps, counters, tags, caption copy. |

---

## 3. Borders, Dividers & State Semantics

### Border Distinction Rule
- **Subtle Border** (`#E7E4DD` light / `#40545D` dark): Used strictly for decorative dividers and passive card boundaries.
- **Control Border** (`#767E86` light / `#60747B` dark): Used for inputs, outline buttons, selects, and interactive control boundaries.
- *Never rely on the subtle border alone to communicate an interactive control.*

### Semantic Status States
| State | Light Text/Icon | Dark Text/Icon | Dark Surface | Rule |
|---|---|---|---|---|
| **Success** | `#4F6640` | `#B7C9A8` | `#1B2A1E` | Must always pair with an icon + descriptive label. |
| **Warning** | `#8A5A00` | `#E8C47A` | `#2A2314` | Caution, timing expiry, incomplete checkpoints. |
| **Error** | `#9E1C2B` | `#F29A9A` | `#321A1E` | Distinct from action orange; indicates validation failure. |
| **Info** | `#2F5D7C` | `#8FC3E0` | `#172833` | Neutral system information and hints. |

---

## 4. Component Geometry, Shadows & Motion

```text
┌─────────────────────────────────────────────────────────────┐
│                      COMPONENT GEOMETRY                     │
├───────────────────────────────┬─────────────────────────────┤
│ Card & Input Radius           │ 8px                         │
│ Button Radius                 │ 6px (No oversized pills)    │
│ Product Screenshot Radius     │ 12px                        │
│ Tag / Pill Chip Radius        │ 9999px (Pills for chips only)
│ Marketing Proof Shadow        │ 0 8px 24px rgba(23,37,43,0.08)
│ Product UI Shadow             │ 0 1px 3px rgba(23,37,43,0.08) max
│ Transition Motion             │ 150ms ease-out              │
└───────────────────────────────┴─────────────────────────────┘
```

---

## 5. Explicit Visual Non-Goals & Anti-Patterns

> [!CAUTION]
> **STRICT VISUAL PROHIBITIONS**:
> 1. **No Edtech Clichés**: No Poppins, Montserrat, or Nunito fonts. No cartoon mascots, confetti, blobs, or playful badges.
> 2. **No Neon / Aurora / Glassmorphism**: No gradient-mesh hero backgrounds, glassmorphism overlays, or pulsing neon effects.
> 3. **No Heavy Shadows**: Maximum product UI shadow is `0 1px 3px`. Never apply diffuse multi-layer drop shadows to UI cards.
> 4. **No Pill-Everything UI**: Pill radii (`rounded-full`) are reserved exclusively for tags and filter chips. Buttons MUST use 6px radius (`rounded-md`).
> 5. **No Hover Displacement**: Elements must NEVER shift vertically or float on hover (`translate-y` is strictly banned).
> 6. **No Fake Window Chrome**: Do NOT add fake red/yellow/green macOS terminal dots to modals or screenshots.
> 7. **No More than Two Accent Colors in One Viewport**: Maintain strict visual discipline.

---

## 6. Pre-Commit Design System Checklist

Before committing any frontend UI change in CurioTech or CareerCafe:
- [ ] Primary action buttons use `#BC4A1E` and darken to `#A83E17` on hover.
- [ ] No gray-black substitutes used for teal-cast `#17252B`.
- [ ] Cards use 8px radius; buttons use 6px radius.
- [ ] Subtle border is distinguished from interactive control borders.
- [ ] Shadows are restrained to `0 1px 3px` in product UI.
- [ ] All interactive controls feature visible 2px focus rings with 2px offset.

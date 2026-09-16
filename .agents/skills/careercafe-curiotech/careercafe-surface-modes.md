---
name: careercafe-surface-modes
scope: codebase-curiotech-careercafe
description: >-
  CurioTech and CareerCafe four-surface behavior protocols: Marketing Landing, Question Bank reading column, executable SQL/Python playground, and distraction-free Interview Mode.
---

# CareerCafe Four-Surface Behavior Protocols

> [!IMPORTANT]
> **PROJECT-SPECIFIC SCOPE**: This skill is strictly specific to **CurioTech** and **CareerCafe** projects (`CurioTech-CareerCafe`). It governs the four distinct operational surfaces within the platform to maintain brand unity while tailoring density and interaction rules.

---

## 1. Architectural Matrix: One Brand, Four Surfaces

CareerCafe features four primary surface contexts sharing one core design system, differing only in density, focus, and visual temperature:

| Surface | Density | Visual Temperature | Primary Goal |
|---|---|---|---|
| **1. Marketing Landing Page** | Generous / Spacious | Warm, editorial, proof-driven | Value proposition communication and conversion. |
| **2. Question Bank & Content** | Medium / Editorial | Calm, comfortable reading | Deep technical study and comprehension. |
| **3. SQL & Python Playground** | High / Compact | Technical, dark-executable | Code execution, testing, and debugging. |
| **4. Interview Mode** | High / Strict | Near-monochrome, high-pressure | Realistic timed interview simulation. |

---

## 2. Surface 1: Marketing Landing Page

- **Container Constraint**: Maximum width of `1200px` with generous section spacing (`96px` desktop / `64px` mobile).
- **Hero Treatment**: Crisp single-color Charcoal headline (`#17252B`), warm off-white canvas (`#F8F7F4`), vibrant `#BC4A1E` CTA.
- **Dark Mode Transition**: Uses `#10181C` canvas with `#17252B` and `#1D2D33` surface cards; maintains exact typographic hierarchy without neon treatments.
- **Conversion Band**: The only full-width dark band is the final student CTA section.

---

## 3. Surface 2: Question Bank & Long-Form Reading

- **Strict Reading Measure**: The content column must NEVER exceed `680px` in width. Unbounded text lines cause cognitive strain during 45+ minute study sessions.
- **Typography**: `17px` font size with generous `30px` line-height (`leading-[30px]`) in `IBM Plex Sans` (weight 450).
- **Dark Mode Reading Invariant**:
  - In dark mode, long-form reading surfaces use `#17252B`.
  - Body text uses softened `#C9D2CF`.
  - **PROHIBITION**: Never use pure white (`#FFFFFF`) for long-form paragraphs in dark mode, as it produces intense contrast glare.

---

## 4. Surface 3: SQL & Python Playground

The playground provides an executable code environment alongside dataset schemas and execution results:

- **Chrome & Layout**: Split-panel workspace with `#17252B` pane dividers and dark `#132025` editor background.
- **Editor Controls**: Compact 14px UI buttons with 6px border radius.
- **Syntax Token Palette (Monaco / CodeMirror)**:

| Syntax Element | Token Hex | Visual Identity |
|---|---|---|
| **Editor Background** | `#132025` | Deep dark teal canvas |
| **Editor Chrome** | `#17252B` | Teal-cast panel header |
| **Default Foreground**| `#D4DCDA` | Crisp muted text |
| **Keywords** | `#B79BE8` | Soft lilac |
| **Strings** | `#A9C295` | Olive-tinted sage |
| **Numbers** | `#E0B26A` | Warm ochre |
| **Functions** | `#7FB4E3` | Muted cornflower blue |
| **Types** | `#72C4B8` | Technical aquamarine |
| **Operators** | `#E8825A` | Restrained orange-coral |
| **Comments** | `#7C8F94` | Slate teal (low contrast) |
| **Error Underline** | `#E88070` | Soft salmon red |

---

## 5. Surface 4: Interview Mode (Simulation Invariant)

Interview Mode simulates genuine high-stakes technical interviews under time constraints:

> [!CAUTION]
> **NO SAGE OR ENCOURAGEMENT DURING ACTIVE SESSIONS**:
> - **Near-Monochrome Environment**: Active interview sessions must eliminate extraneous color cues, gamification elements, and progress confetti.
> - **Zero Success Feedback Mid-Session**: Never show green/sage status highlights, "Good Job!" banners, or hint confirmations while the countdown timer is running.
> - **Focus Invariant**: Display only three elements:
>   1. The active question stem / dataset prompt.
>   2. The code/answer input area.
>   3. The tabular countdown timer (`20px–32px IBM Plex Mono`).
> - **Completion Transition**: State and evaluation feedback (Sage `#7F9B6D`) is revealed ONLY after the interview concludes and results are finalized.

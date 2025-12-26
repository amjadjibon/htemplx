# Complete Shadcn-Style Components List

This document lists all available shadcn-style components in the htemplx project.

## Component Count: 30+ Components

All components are located in `/app/views/components/` and are fully integrated with HTMX, Hyperscript, and Tailwind CSS.

---

## Form Components

### 1. **Button** (`button.templ`)
- Variants: default, destructive, outline, secondary, ghost, link
- Sizes: sm, default, lg, icon
- Components: `Button`, `ButtonWithIcon`, `IconButton`

### 2. **Input** (`input.templ`)
- Components: `Input`, `InputWithLabel`, `InputWithError`
- Additional: `Textarea`, `TextareaWithLabel`, `Select`, `SelectWithLabel`

### 3. **Checkbox** (`checkbox.templ`)
- Components: `Checkbox`, `CheckboxWithLabel`, `CheckboxGroup`
- Interactive with Hyperscript

### 4. **Radio** (`radio.templ`)
- Components: `RadioGroup`, `RadioItem`
- Mutually exclusive selection

### 5. **Switch** (`switch.templ`)
- Components: `Switch`, `SwitchWithLabel`, `SwitchWithDescription`
- Toggle on/off states

### 6. **Slider** (`slider.templ`)
- Components: `Slider`, `SliderWithLabel`
- Range input with custom styling

### 7. **Label** (`label.templ`)
- Components: `Label`, `LabelWithDescription`
- Form labels with optional required indicator

---

## Data Display

### 8. **Card** (`card.templ`)
- Components: `Card`, `CardHeader`, `CardTitle`, `CardDescription`, `CardContent`, `CardFooter`
- Pre-built: `CardSimple`, `CardWithFooter`

### 9. **Badge** (`badge.templ`)
- Variants: default, secondary, destructive, outline, success, warning
- Components: `Badge`, `BadgeWithDot` (animated)

### 10. **Avatar** (`avatar.templ`)
- Sizes: sm, md, lg, xl
- Components: `Avatar`, `AvatarWithFallback`, `AvatarFallback`, `AvatarGroup`
- Automatic initials generation

### 11. **Progress** (`progress.templ`)
- Components: `Progress`, `ProgressWithLabel`, `ProgressIndeterminate`, `ProgressCircle`
- Linear and circular progress indicators

### 12. **Table** (`table.templ`)
- Existing component for tabular data

### 13. **Accordion** (`accordion.templ`)
- Existing component for expandable sections

---

## Feedback

### 14. **Alert2** (`alert2.templ`)
- Variants: default, destructive, success, warning, info
- Components: `Alert2`, `Alert2Simple`, `Alert2WithAction`
- Icons for each variant

### 15. **Toast** (`toast.templ`)
- Variants: default, success, error, warning, info
- Components: `Toast`, `ToastSimple`, `ToastContainer`
- Auto-dismiss with manual close

### 16. **Skeleton** (`skeleton.templ`)
- Components: `Skeleton`, `SkeletonText`, `SkeletonCard`, `SkeletonAvatar`
- Pre-built: `SkeletonTable`, `SkeletonList`, `SkeletonProfile`, `SkeletonCardGrid`

---

## Overlays

### 17. **Dialog** (`dialog.templ`)
- Components: `Dialog`, `DialogSimple`, `DialogConfirm`, `AlertDialog`
- Modal dialogs with backdrop

### 18. **Sheet** (`sheet.templ`)
- Components: `Sheet`, `SheetSimple`
- Slide-out panels from any side (left, right, top, bottom)

### 19. **Popover** (`popover.templ`)
- Components: `Popover`, `PopoverContent`
- Floating content panels

### 20. **Tooltip** (`tooltip.templ`)
- Components: `Tooltip`, `TooltipSimple`
- Positions: top, bottom, left, right

### 21. **Hover Card** (using `popover.templ`)
- Similar to Popover but triggered on hover

---

## Navigation

### 22. **Tabs** (`tabs.templ`)
- Components: `Tabs` (HTMX), `SimpleTabs` (client-side)
- Lazy loading support

### 23. **Breadcrumb** (`breadcrumb.templ`)
- Components: `Breadcrumb`, `BreadcrumbWithHome`, `BreadcrumbCollapsed`
- Navigation breadcrumbs

### 24. **Dropdown Menu** (`dropdown.templ`)
- Components: `DropdownMenu`, `DropdownMenuItem`
- Context menus and action menus

### 25. **Pagination** (`pagination.templ`)
- Existing component for page navigation

### 26. **Navbar** (`navbar.templ`)
- Existing navigation bar component

---

## Layout

### 27. **Separator** (`separator.templ`)
- Components: `Separator`, `SeparatorWithText`
- Horizontal and vertical dividers

### 28. **Aspect Ratio** (`aspectratio.templ`)
- Components: `AspectRatio`, `AspectRatioImage`, `AspectRatioVideo`
- Ratios: 16/9, 4/3, 1/1, 21/9, 3/2, 9/16

### 29. **Scroll Area** (`scrollarea.templ`)
- Components: `ScrollArea`, `ScrollAreaHorizontal`
- Custom scrollbars

### 30. **Collapsible** (`collapsible.templ`)
- Components: `Collapsible`, `CollapsibleWithTitle`
- Expandable/collapsible content

---

## Interactive

### 31. **Toggle** (`toggle.templ`)
- Components: `Toggle`, `ToggleWithIcon`, `ToggleGroup`
- On/off button states

### 32. **Carousel** (`carousel.templ`)
- Existing image carousel component

---

## Existing Components (Pre-existing)

### 33. **Alert** (`alert.templ`)
- Original alert component (simpler version)

### 34. **Footer** (`footer.templ`)
- Page footer

### 35. **Sidebar** (`sidebar.templ`)
- Navigation sidebar

### 36. **Contact** (`contact.templ`)
- Contact form component

### 37. **Login** (`login.templ`)
- Login form

### 38. **Register** (`register.templ`)
- Registration form

### 39. **Forgot Password** (`forgot_password.templ`)
- Password reset form

### 40. **Pricing** (`pricing.templ`)
- Pricing table

### 41. **Newsletter** (`newsletter.templ`)
- Newsletter subscription

### 42. **Clipboard** (`clipboard.templ`)
- Copy to clipboard functionality

### 43. **Pastebin** (`pastebin.templ`)
- Code snippet display

### 44. **About** (`about.templ`)
- About section

### 45. **Under Construction** (`under_construction.templ`)
- Under construction page

---

## Usage Categories

### Form & Input (9 components)
Button, Input, Checkbox, Radio, Switch, Slider, Label, Select, Textarea

### Data Display (6 components)
Card, Badge, Avatar, Progress, Table, Accordion

### Feedback (3 components)
Alert2, Toast, Skeleton

### Overlays (4 components)
Dialog, Sheet, Popover, Tooltip

### Navigation (5 components)
Tabs, Breadcrumb, Dropdown Menu, Pagination, Navbar

### Layout (4 components)
Separator, Aspect Ratio, Scroll Area, Collapsible

### Interactive (2 components)
Toggle, Carousel

---

## Component Features Matrix

| Component | Dark Mode | HTMX | Hyperscript | Accessible | Responsive |
|-----------|-----------|------|-------------|------------|------------|
| Button | ✅ | ✅ | ✅ | ✅ | ✅ |
| Card | ✅ | ✅ | ✅ | ✅ | ✅ |
| Input | ✅ | ✅ | ✅ | ✅ | ✅ |
| Badge | ✅ | ✅ | ✅ | ✅ | ✅ |
| Dialog | ✅ | ✅ | ✅ | ✅ | ✅ |
| Tabs | ✅ | ✅ | ✅ | ✅ | ✅ |
| Skeleton | ✅ | ✅ | ✅ | ✅ | ✅ |
| Toast | ✅ | ✅ | ✅ | ✅ | ✅ |
| Alert2 | ✅ | ✅ | ✅ | ✅ | ✅ |
| Avatar | ✅ | ✅ | ✅ | ✅ | ✅ |
| Breadcrumb | ✅ | ✅ | ✅ | ✅ | ✅ |
| Checkbox | ✅ | ✅ | ✅ | ✅ | ✅ |
| Radio | ✅ | ✅ | ✅ | ✅ | ✅ |
| Switch | ✅ | ✅ | ✅ | ✅ | ✅ |
| Slider | ✅ | ✅ | ✅ | ✅ | ✅ |
| Progress | ✅ | ✅ | ✅ | ✅ | ✅ |
| Separator | ✅ | ✅ | ✅ | ✅ | ✅ |
| Label | ✅ | ✅ | ✅ | ✅ | ✅ |
| Tooltip | ✅ | ✅ | ✅ | ✅ | ✅ |
| Popover | ✅ | ✅ | ✅ | ✅ | ✅ |
| Dropdown | ✅ | ✅ | ✅ | ✅ | ✅ |
| Sheet | ✅ | ✅ | ✅ | ✅ | ✅ |
| Toggle | ✅ | ✅ | ✅ | ✅ | ✅ |
| Collapsible | ✅ | ✅ | ✅ | ✅ | ✅ |
| Aspect Ratio | ✅ | ✅ | ✅ | ✅ | ✅ |
| Scroll Area | ✅ | ✅ | ✅ | ✅ | ✅ |

---

## Quick Start

### Import in Templ
```templ
import "htemplx/app/views/components"
```

### Basic Usage
```templ
// Button
@components.Button("Click me", "default", "default")

// Card
@components.CardSimple("Title", "Description", contentComponent)

// Form inputs
@components.InputWithLabel("Email", "email", "you@example.com", "email")
@components.CheckboxWithLabel("agree", "yes", "I agree to terms", false)
@components.Switch("notifications", true)

// Feedback
@components.Toast("toast-1", "Success!", "Saved successfully", "success")
@components.SkeletonCard()

// Navigation
@components.Tabs("my-tabs", tabItems, "tab1")
@components.Breadcrumb(breadcrumbItems)
```

---

## Documentation

- **Comprehensive Guide**: [`docs/SHADCN_COMPONENTS.md`](./SHADCN_COMPONENTS.md)
- **Developer Guide**: [`CLAUDE.md`](../CLAUDE.md)
- **Component Examples**: See individual component files

---

## Design System

All components follow:
- **shadcn/ui** design principles
- **Tailwind CSS** utility classes
- **Dark mode** first approach
- **Accessibility** standards (WCAG AA)
- **HTMX** integration patterns
- **Hyperscript** for interactivity

---

## Contributing

When adding new components:
1. Follow shadcn/ui design patterns
2. Support all variants and sizes
3. Include dark mode styles
4. Add ARIA attributes
5. Test with HTMX
6. Document with examples

---

## Resources

- [shadcn/ui Documentation](https://ui.shadcn.com/)
- [Tailwind CSS](https://tailwindcss.com/)
- [HTMX](https://htmx.org/)
- [Hyperscript](https://hyperscript.org/)
- [Templ Guide](https://templ.guide/)

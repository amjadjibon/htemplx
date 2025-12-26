# Shadcn-Style Components for Templ

This document provides comprehensive documentation for the shadcn-style components implemented in Templ for the htemplx project.

## Table of Contents

1. [Overview](#overview)
2. [Button](#button)
3. [Card](#card)
4. [Input](#input)
5. [Badge](#badge)
6. [Dialog](#dialog)
7. [Tabs](#tabs)
8. [Skeleton](#skeleton)
9. [Toast](#toast)
10. [Usage Examples](#usage-examples)

---

## Overview

These components are inspired by [shadcn/ui](https://ui.shadcn.com/) and ported to Templ for use in Go applications. They feature:

- **Tailwind CSS** for styling
- **Dark mode support** with automatic theming
- **HTMX compatibility** for dynamic interactions
- **Hyperscript support** for client-side logic
- **Accessibility** features (ARIA labels, keyboard navigation)
- **Type-safe** component props via Templ

All components are located in `/app/views/components/` and follow the naming convention `*.templ`.

---

## Button

**Location:** `app/views/components/button.templ`

### Variants

- `default` - Primary button (dark background)
- `destructive` - Red button for dangerous actions
- `outline` - Outlined button with border
- `secondary` - Secondary gray button
- `ghost` - Transparent button with hover effect
- `link` - Text button with underline on hover

### Sizes

- `sm` - Small button (h-9)
- `default` - Default size (h-10)
- `lg` - Large button (h-11)
- `icon` - Icon-only button (h-10 w-10)

### Components

#### Button

```templ
@components.Button("Click me", "default", "default", templ.Attributes{
    "type": "submit",
})
```

#### ButtonWithIcon

```templ
@components.ButtonWithIcon("Save", iconComponent, "default", "default", templ.Attributes{
    "hx-post": "/save",
})
```

#### IconButton

```templ
@components.IconButton(iconComponent, "ghost", templ.Attributes{
    "aria-label": "Close",
})
```

### Examples

```templ
// Primary button
@components.Button("Submit", "default", "default")

// Destructive button
@components.Button("Delete", "destructive", "default")

// Outline button (small)
@components.Button("Cancel", "outline", "sm")

// Ghost button (large)
@components.Button("Learn More", "ghost", "lg")

// Link button
@components.Button("Read docs", "link", "default")

// With HTMX
@components.Button("Load More", "secondary", "default", templ.Attributes{
    "hx-get": "/api/items",
    "hx-target": "#items-list",
    "hx-swap": "beforeend",
})
```

---

## Card

**Location:** `app/views/components/card.templ`

### Components

- `Card` - Container for card sections
- `CardHeader` - Header section
- `CardTitle` - Title text
- `CardDescription` - Description text
- `CardContent` - Main content area
- `CardFooter` - Footer section
- `CardSimple` - Pre-built simple card
- `CardWithFooter` - Pre-built card with footer

### Examples

```templ
// Manual card construction
@components.Card(
    components.CardHeader(
        components.CardTitle("Project Overview"),
        components.CardDescription("Last updated 2 hours ago"),
    ),
    components.CardContent(
        // Your content here
    ),
    components.CardFooter(
        components.Button("View Details", "outline", "sm"),
    ),
)

// Simple card
@components.CardSimple(
    "Dashboard",
    "View your analytics",
    myContentComponent,
)

// Card with footer
@components.CardWithFooter(
    "Settings",
    "Manage your account settings",
    myContentComponent,
    myFooterComponent,
)
```

---

## Input

**Location:** `app/views/components/input.templ`

### Components

#### Input

Basic input field with Tailwind styling.

```templ
@components.Input("text", "Enter your name", "username")
```

#### InputWithLabel

Input with label above.

```templ
@components.InputWithLabel("Email", "email", "you@example.com", "email", templ.Attributes{
    "required": "true",
})
```

#### InputWithError

Input with error state and message.

```templ
@components.InputWithError(
    "Password",
    "password",
    "Enter password",
    "password",
    "Password must be at least 8 characters",
)
```

#### Textarea

Multiline text input.

```templ
@components.Textarea("Enter your message", "message", 4)
```

#### TextareaWithLabel

Textarea with label.

```templ
@components.TextareaWithLabel("Description", "Tell us about yourself", "bio", 5)
```

#### Select

Dropdown select input.

```templ
@components.Select("country", []string{"USA", "Canada", "UK", "Australia"})
```

#### SelectWithLabel

Select with label.

```templ
@components.SelectWithLabel("Country", "country", []string{"USA", "Canada", "UK"})
```

### Examples

```templ
// Login form
<form hx-post="/login" hx-swap="outerHTML">
    @components.InputWithLabel("Email", "email", "you@example.com", "email", templ.Attributes{
        "required": "true",
    })
    @components.InputWithLabel("Password", "password", "", "password", templ.Attributes{
        "required": "true",
    })
    @components.Button("Sign In", "default", "default", templ.Attributes{
        "type": "submit",
    })
</form>

// Contact form
<form hx-post="/contact">
    @components.InputWithLabel("Name", "text", "Your name", "name")
    @components.InputWithLabel("Email", "email", "you@example.com", "email")
    @components.TextareaWithLabel("Message", "Your message", "message", 5)
    @components.Button("Send", "default", "default", templ.Attributes{"type": "submit"})
</form>
```

---

## Badge

**Location:** `app/views/components/badge.templ`

### Variants

- `default` - Dark badge
- `secondary` - Gray badge
- `destructive` - Red badge
- `outline` - Outlined badge
- `success` - Green badge
- `warning` - Yellow badge

### Components

#### Badge

Basic badge.

```templ
@components.Badge("New", "success")
```

#### BadgeWithDot

Badge with animated status dot.

```templ
@components.BadgeWithDot("Active", "success")
```

### Examples

```templ
// Status badges
@components.Badge("Published", "success")
@components.Badge("Draft", "secondary")
@components.Badge("Deleted", "destructive")
@components.Badge("Pending", "warning")

// With animated dot
@components.BadgeWithDot("Live", "success")
@components.BadgeWithDot("Processing", "warning")
@components.BadgeWithDot("Error", "destructive")

// Outline badge
@components.Badge("v2.0", "outline")
```

---

## Dialog

**Location:** `app/views/components/dialog.templ`

### Components

#### Dialog

Full-featured modal dialog.

```templ
@components.Dialog(
    "confirm-dialog",
    "Are you sure?",
    "This action cannot be undone.",
    contentComponent,
    footerComponent,
)
```

#### DialogSimple

Simple dialog with just content.

```templ
@components.DialogSimple("my-dialog", "Welcome!", contentComponent)
```

#### DialogConfirm

Confirmation dialog with confirm/cancel buttons.

```templ
@components.DialogConfirm(
    "delete-dialog",
    "Delete Item",
    "Are you sure you want to delete this item?",
    "Delete",
    "Cancel",
    "/api/items/123",
)
```

#### AlertDialog

Alert dialog with single action button.

```templ
@components.AlertDialog("alert", "Success!", "Your changes have been saved.", "OK")
```

### Examples

```templ
// Trigger dialog with HTMX
<button
    hx-get="/dialogs/confirm-delete"
    hx-target="body"
    hx-swap="beforeend"
>
    Delete Item
</button>

// Handler returns:
@components.DialogConfirm(
    "delete-confirm",
    "Delete Item",
    "This will permanently delete the item.",
    "Delete",
    "Cancel",
    "/api/items/delete/123",
)

// Simple info dialog
@components.DialogSimple("info", "Information",
    templ.Raw("<p>This is some important information.</p>"),
)
```

### Dialog Features

- **Backdrop blur** - Visual focus on dialog
- **Click outside to close** - Uses Hyperscript
- **Escape key support** - Browser default
- **ARIA attributes** - Accessibility
- **Auto-remove** - Cleans up DOM after close

---

## Tabs

**Location:** `app/views/components/tabs.templ`

### Components

#### Tabs

Tabs with HTMX support for lazy loading.

```templ
@components.Tabs("my-tabs", []components.TabItem{
    {ID: "overview", Label: "Overview", Content: overviewComponent},
    {ID: "details", Label: "Details", Content: detailsComponent},
    {ID: "settings", Label: "Settings", Content: settingsComponent},
}, "overview")
```

#### SimpleTabs

Tabs without HTMX (pure client-side).

```templ
@components.SimpleTabs("simple-tabs", []components.TabItem{
    {ID: "tab1", Label: "Tab 1", Content: tab1Component},
    {ID: "tab2", Label: "Tab 2", Content: tab2Component},
}, "tab1")
```

### Examples

```templ
// Basic tabs
@components.SimpleTabs("product-tabs", []components.TabItem{
    {
        ID: "description",
        Label: "Description",
        Content: templ.Raw("<p>Product description here...</p>"),
    },
    {
        ID: "reviews",
        Label: "Reviews",
        Content: reviewsComponent,
    },
    {
        ID: "shipping",
        Label: "Shipping",
        Content: shippingComponent,
    },
}, "description")

// With HTMX lazy loading
// Server needs to handle: GET /tabs/user-tabs/{tab-id}
@components.Tabs("user-tabs", []components.TabItem{
    {ID: "profile", Label: "Profile", Content: profileComponent},
    {ID: "activity", Label: "Activity", Content: nil}, // Loaded via HTMX
    {ID: "settings", Label: "Settings", Content: nil},
}, "profile")
```

---

## Skeleton

**Location:** `app/views/components/skeleton.templ`

### Components

#### Skeleton

Basic skeleton with custom classes.

```templ
@components.Skeleton("h-4 w-full")
```

#### SkeletonText

Multiple text lines.

```templ
@components.SkeletonText(3)
```

#### SkeletonCard

Card skeleton.

```templ
@components.SkeletonCard()
```

#### SkeletonAvatar

Avatar skeleton (sm, md, lg, xl).

```templ
@components.SkeletonAvatar("md")
```

#### SkeletonButton

Button skeleton.

```templ
@components.SkeletonButton("default")
```

#### SkeletonInput

Input field skeleton.

```templ
@components.SkeletonInput()
```

#### SkeletonTable

Table skeleton.

```templ
@components.SkeletonTable(5, 4) // 5 rows, 4 columns
```

#### SkeletonList

List skeleton.

```templ
@components.SkeletonList(5) // 5 items
```

#### SkeletonCardGrid

Grid of cards.

```templ
@components.SkeletonCardGrid(6) // 6 cards
```

#### SkeletonProfile

Profile layout skeleton.

```templ
@components.SkeletonProfile()
```

### Examples

```templ
// Show skeleton while loading data
<div hx-get="/api/users" hx-trigger="load" hx-swap="outerHTML">
    @components.SkeletonList(5)
</div>

// Card grid loading state
<div id="cards" hx-get="/api/cards" hx-trigger="load">
    @components.SkeletonCardGrid(6)
</div>

// Profile loading
<div hx-get="/api/profile" hx-trigger="load" hx-swap="innerHTML">
    @components.SkeletonProfile()
</div>
```

---

## Toast

**Location:** `app/views/components/toast.templ`

### Variants

- `default` - White/dark toast
- `success` - Green toast
- `error` - Red toast
- `warning` - Yellow toast
- `info` - Blue toast

### Components

#### ToastContainer

Add once to your base layout.

```templ
@components.ToastContainer()
```

#### Toast

Full toast with title and description.

```templ
@components.Toast("toast-1", "Success!", "Your changes have been saved.", "success")
```

#### ToastSimple

Simple toast with just message.

```templ
@components.ToastSimple("toast-2", "Item deleted", "destructive")
```

### Examples

```templ
// In your base layout (layouts/base.templ)
<body>
    @components.ToastContainer()
    // ... rest of layout
</body>

// Handler returning toast via HTMX
func (h *WebHandler) SaveItem(w http.ResponseWriter, r *http.Request) {
    // ... save logic

    // Return toast with OOB swap
    w.Header().Set("HX-Trigger", "item-saved")
    render(w, r, components.Toast(
        "toast-save",
        "Success",
        "Item saved successfully",
        "success",
    ))
}

// In template, trigger via HTMX
<form
    hx-post="/items/save"
    hx-target="#toast-container"
    hx-swap="beforeend"
>
    <!-- form fields -->
</form>

// Different variants
@components.ToastSimple("t1", "Operation successful", "success")
@components.ToastSimple("t2", "An error occurred", "error")
@components.ToastSimple("t3", "Warning: Check your input", "warning")
@components.ToastSimple("t4", "New message received", "info")
```

### Toast Features

- **Auto-dismiss** - 5 seconds by default
- **Manual close** - X button
- **Fade out animation** - Smooth transition
- **Stacking** - Multiple toasts stack vertically
- **Icons** - Variant-specific icons

---

## Usage Examples

### Complete Form with Validation

```templ
<form hx-post="/api/users" hx-swap="outerHTML" class="space-y-4">
    @components.CardSimple(
        "Create User",
        "Add a new user to the system",
        templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
            return components.InputWithLabel("Name", "text", "John Doe", "name", templ.Attributes{
                "required": "true",
            }).Render(ctx, w)
        }),
    )

    @components.InputWithLabel("Email", "email", "john@example.com", "email", templ.Attributes{
        "required": "true",
    })

    @components.SelectWithLabel("Role", "role", []string{"Admin", "User", "Guest"})

    <div class="flex gap-2">
        @components.Button("Create User", "default", "default", templ.Attributes{"type": "submit"})
        @components.Button("Cancel", "outline", "default", templ.Attributes{"type": "button"})
    </div>
</form>
```

### Dashboard with Skeletons

```templ
<div class="space-y-6">
    <h1 class="text-3xl font-bold">Dashboard</h1>

    <!-- Stats cards -->
    <div
        hx-get="/api/stats"
        hx-trigger="load"
        hx-swap="innerHTML"
        class="grid gap-4 md:grid-cols-3"
    >
        @components.SkeletonCardGrid(3)
    </div>

    <!-- Recent activity -->
    <div hx-get="/api/activity" hx-trigger="load" hx-swap="innerHTML">
        @components.SkeletonList(5)
    </div>
</div>
```

### Modal Confirmation Flow

```templ
// Delete button
<button
    hx-get="/confirm-delete/123"
    hx-target="body"
    hx-swap="beforeend"
    class="text-red-600"
>
    Delete
</button>

// Handler returns (GET /confirm-delete/:id)
@components.DialogConfirm(
    "delete-confirm-123",
    "Delete Item",
    "This action cannot be undone. Are you sure?",
    "Delete",
    "Cancel",
    "/api/items/123",
)

// After deletion, return toast (POST /api/items/123)
@components.Toast(
    "delete-success",
    "Deleted",
    "Item has been deleted successfully",
    "success",
)
```

### Tabbed Interface with Lazy Loading

```templ
@components.Tabs("user-profile", []components.TabItem{
    {
        ID: "overview",
        Label: "Overview",
        Content: templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
            // Render overview immediately
            return nil
        }),
    },
    {
        ID: "activity",
        Label: "Activity",
        Content: components.SkeletonList(5), // Show skeleton initially
    },
    {
        ID: "settings",
        Label: "Settings",
        Content: components.SkeletonCard(),
    },
}, "overview")

// Server handles lazy loading
// GET /tabs/user-profile/activity
// Returns actual activity content
```

---

## Integration with HTMX

All components are designed to work seamlessly with HTMX. Common patterns:

### OOB Swaps

```templ
// Return toast notification out-of-band
hx-swap-oob="beforeend:#toast-container"
```

### Target Swapping

```templ
// Replace dialog content
hx-target="#dialog-container"
hx-swap="innerHTML"
```

### Trigger Events

```templ
// Trigger after successful action
hx-trigger="item-saved from:body"
```

---

## Accessibility Features

All components include:

- **ARIA labels** and roles
- **Keyboard navigation** support
- **Focus management** for dialogs
- **Screen reader** friendly text
- **Color contrast** meeting WCAG AA standards
- **Focus indicators** for keyboard users

---

## Dark Mode Support

All components automatically support dark mode using Tailwind's `dark:` prefix. Toggle dark mode by adding the `dark` class to the `<html>` element:

```javascript
// Toggle dark mode
document.documentElement.classList.toggle('dark')
```

Or configure in `tailwind.config.js`:

```javascript
darkMode: 'class', // or 'media'
```

---

## Customization

To customize component styles:

1. **Modify Tailwind classes** directly in component files
2. **Extend Tailwind config** in `tailwind.config.js`
3. **Add custom CSS** in `public/assets/css/input.css`
4. **Create variant functions** for new variants

Example custom variant:

```go
// In button.templ
case "brand":
    variantClasses = "bg-purple-600 text-white hover:bg-purple-700"
```

---

## Best Practices

1. **Use semantic variants** - Choose the variant that matches the action
2. **Combine with HTMX** - Leverage server-side rendering
3. **Show loading states** - Use skeletons for better UX
4. **Provide feedback** - Use toasts for user actions
5. **Keep modals focused** - Use for important actions only
6. **Test accessibility** - Use keyboard navigation and screen readers
7. **Respect dark mode** - Test in both light and dark themes

---

## Contributing

When adding new components:

1. Follow the shadcn/ui design system
2. Support all variants and sizes
3. Include dark mode support
4. Add accessibility features
5. Document with examples
6. Test with HTMX integration

---

## Resources

- [shadcn/ui](https://ui.shadcn.com/) - Original React components
- [Tailwind CSS](https://tailwindcss.com/) - Utility-first CSS framework
- [HTMX](https://htmx.org/) - HTML over the wire
- [Hyperscript](https://hyperscript.org/) - Frontend scripting
- [Templ](https://templ.guide/) - Type-safe Go templates

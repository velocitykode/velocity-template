# Velocity Template - React

React 19 + Inertia.js starter template for the [Velocity](https://github.com/velocitykode/velocity) Go web framework.

This repo is a **template** consumed by the Velocity installer. To start a new project:

```bash
velocity new myapp --stack=react
```

The installer clones this template, rewrites the module placeholders, installs dependencies, builds the project's `vel` binary, and runs the initial migrations.

## Stack

- **Backend**: Velocity Go framework
- **Frontend**: React 19 + TypeScript 5
- **Rendering**: Inertia.js 3 (`@inertiajs/react`)
- **UI**: shadcn/ui (Radix primitives) + Headless UI + lucide-react
- **Styling**: Tailwind CSS 4
- **Build**: Vite 7 (with `@velocitykode/velocity-vite-plugin`)

## Documentation

Full documentation at **[velocity.velocitykode.com/docs](https://velocity.velocitykode.com/docs)**

## Sibling Templates

- [`velocity-template-vue`](https://github.com/velocitykode/velocity-template-vue) - Vue 3 + Inertia.js (with SSR)
- [`velocity-template-api`](https://github.com/velocitykode/velocity-template-api) - API only (no frontend)

## License

MIT

# Velocity Template - React

Full-stack starter template for the [Velocity](https://github.com/velocitykode/velocity) Go web framework using React + Inertia.js.

## Stack

- **Backend**: Velocity Go framework
- **Frontend**: React 19 + TypeScript 5
- **Rendering**: Inertia.js 3 (`@inertiajs/react`)
- **UI**: shadcn/ui (Radix primitives) + Headless UI + lucide-react
- **Styling**: Tailwind CSS 4
- **Build**: Vite 7 (with `@velocitykode/velocity-vite-plugin`)

## Usage

This template is used automatically by the Velocity CLI:

```bash
velocity new myapp
cd myapp
./vel serve
```

## Development Commands

```bash
# Start development server with hot reload
./vel serve

# Run database migrations
./vel migrate

# Generate a new controller
./vel make:controller User

# Build for production
./vel build
```

## Documentation

Full documentation at **[velocity.velocitykode.com/docs](https://velocity.velocitykode.com/docs)**

## License

MIT

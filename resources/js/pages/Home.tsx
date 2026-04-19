import { AppLogo, Badge, Button, ThemeToggle } from '@/components/ui';
import { usePageTitle } from '@/hooks/use-page-title';
import { Link } from '@inertiajs/react';
import { ArrowRight, BookOpen, Github } from 'lucide-react';

export default function Home({ message }: { message?: string }) {
  usePageTitle('Welcome');

  return (
    <div className="min-h-screen bg-bg text-fg flex flex-col">
      {/* Top bar */}
      <header className="flex h-12 items-center justify-between border-b border-border px-4 sm:px-5 lg:px-6">
        <AppLogo variant="full" size="md" href="/" />
        <div className="flex items-center gap-2">
          <ThemeToggle />
          <Link
            href="/login"
            className="caption-mono-upper text-muted hover:text-fg transition-colors px-2"
          >
            Sign in
          </Link>
        </div>
      </header>

      {/* Main */}
      <main className="flex-1 flex items-center justify-center px-4 py-8 sm:px-5 lg:px-6">
        <div className="w-full max-w-3xl">
          <Badge variant="outline" className="mb-4">
            Ready to ship
          </Badge>
          <h1 className="font-sans text-4xl sm:text-5xl lg:text-6xl xl:text-[68px] font-semibold tracking-tight leading-[1.05]">
            Welcome to Velocity.
            <br />
            <span className="text-muted">
              {message || 'Build faster with pure Go power.'}
            </span>
          </h1>
          <p className="mt-4 max-w-xl text-[15px] lg:text-[17px] leading-relaxed text-muted">
            A batteries-included full-stack framework for Go - routing, auth, db,
            queues, websockets. Twenty-three services, zero yak-shaving.
          </p>

          <div className="mt-6 flex flex-col gap-2.5 sm:flex-row">
            <Link href="/register">
              <Button size="lg" iconRight={ArrowRight} fullWidth>
                Get started
              </Button>
            </Link>
            <a
              href="https://github.com/velocitykode/velocity"
              target="_blank"
              rel="noopener noreferrer"
              className="inline-flex h-12 items-center justify-center gap-2 border border-border-strong px-5 text-[14px] tracking-wide text-fg transition-colors hover:bg-surface-alt"
            >
              <Github className="h-4 w-4" />
              Star on GitHub
            </a>
          </div>

          {/* Resource rows */}
          <div className="mt-10 grid gap-px bg-border sm:grid-cols-2 border border-border">
            <a
              href="https://github.com/velocitykode/velocity"
              target="_blank"
              rel="noopener noreferrer"
              className="group flex items-start gap-3 bg-surface p-4 transition-colors hover:bg-surface-alt"
            >
              <BookOpen className="h-4 w-4 mt-0.5 text-muted" />
              <div className="flex-1">
                <div className="caption-mono-upper text-dim">
                  01 · Documentation
                </div>
                <div className="mt-0.5 text-sm text-fg">
                  Read the full API reference and guides.
                </div>
              </div>
              <ArrowRight className="h-4 w-4 text-dim transition-transform group-hover:translate-x-0.5" />
            </a>
            <a
              href="https://github.com/velocitykode/velocity/tree/main/examples"
              target="_blank"
              rel="noopener noreferrer"
              className="group flex items-start gap-3 bg-surface p-4 transition-colors hover:bg-surface-alt"
            >
              <Github className="h-4 w-4 mt-0.5 text-muted" />
              <div className="flex-1">
                <div className="caption-mono-upper text-dim">
                  02 · Examples
                </div>
                <div className="mt-0.5 text-sm text-fg">
                  Browse real-world starter apps & patterns.
                </div>
              </div>
              <ArrowRight className="h-4 w-4 text-dim transition-transform group-hover:translate-x-0.5" />
            </a>
          </div>
        </div>
      </main>

      {/* Footer */}
      <footer className="border-t border-border px-4 py-3 sm:px-5 lg:px-6">
        <div className="flex flex-col gap-2 sm:flex-row sm:items-center sm:justify-between">
          <p className="caption-mono-upper text-dim">
            © {new Date().getFullYear()} Velocity
          </p>
          <p className="caption-mono text-dim">
            build:pre-1.0 · go/1.26+
          </p>
        </div>
      </footer>
    </div>
  );
}

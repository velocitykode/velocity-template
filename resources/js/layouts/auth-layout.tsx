import { AppLogo, ThemeToggle } from '@/components/ui';
import { Link } from '@inertiajs/react';
import { type PropsWithChildren } from 'react';

interface AuthLayoutProps {
  title: string;
  description: string;
  /** Footer switch-action link ("New to Velocity?" / "Already have an account?") */
  switchPrompt?: string;
  switchHref?: string;
  switchLabel?: string;
}

export default function AuthLayout({
  title,
  description,
  switchPrompt,
  switchHref,
  switchLabel,
  children,
}: PropsWithChildren<AuthLayoutProps>) {
  const year = new Date().getFullYear();

  return (
    <div className="flex min-h-screen flex-col bg-bg lg:flex-row">
      {/* ================================================================
           LEFT - brand panel (always dark surface for consistent silhouette)
         ================================================================ */}
      <aside className="relative hidden overflow-hidden bg-panel text-panel-fg lg:flex lg:w-1/2 lg:flex-col lg:justify-between xl:w-[45%] 2xl:w-[40%]">
        {/* Subtle grid texture */}
        <div className="absolute inset-0 bg-paper-grid" />

        {/* Top row: logo + theme toggle */}
        <div className="relative flex items-start justify-between px-6 py-5 xl:px-8 xl:py-6">
          <AppLogo variant="full" size="xl" href="/" onDarkSurface />
          <ThemeToggle onDarkSurface />
        </div>

        {/* Middle: brand statement */}
        <div className="relative flex flex-1 items-end px-6 pb-6 xl:px-8 xl:pb-8">
          <div className="max-w-md">
            <p className="caption-mono-upper mb-3 text-panel-muted">
              Velocity · v0.22.0
            </p>
            <h2 className="font-sans text-4xl xl:text-5xl 2xl:text-[56px] font-semibold leading-[1.05] tracking-tight text-panel-fg">
              Build Faster.
              <br />
              Ship Sooner.
            </h2>
            <p className="mt-4 max-w-sm caption-mono text-panel-muted">
              Batteries-included Go framework - routing, auth, db, queues, websockets.
              Twenty-three services, zero yak-shaving.
            </p>
          </div>
        </div>

        {/* Bottom: ornament row */}
        <div className="relative border-t border-panel-border px-6 py-4 xl:px-8">
          <div className="flex items-center justify-between caption-mono-upper text-panel-dim">
            <span>© {year} Velocity</span>
            <span>go/1.26+</span>
          </div>
        </div>
      </aside>

      {/* ================================================================
           RIGHT - form panel
         ================================================================ */}
      <main className="flex flex-1 flex-col bg-bg">
        {/* Mobile top bar */}
        <div className="flex h-14 items-center justify-between border-b border-border px-4 lg:hidden">
          <AppLogo variant="icon" size="lg" href="/" />
          <ThemeToggle />
        </div>

        {/* Form container - centered, generous at XXL */}
        <div className="flex flex-1 items-center justify-center px-5 py-8 sm:px-8 lg:px-10 xl:px-16 2xl:px-24">
          <div className="w-full max-w-[420px] xl:max-w-[440px]">
            {/* Headline */}
            <div className="mb-6 lg:mb-8">
              <h1 className="font-sans text-3xl lg:text-[34px] xl:text-[38px] font-semibold tracking-tight text-fg">
                {title}
              </h1>
              <p className="mt-2 text-[14px] leading-relaxed text-muted">
                {description}
              </p>
            </div>

            {children}

            {/* Switch prompt */}
            {switchPrompt && switchHref && switchLabel && (
              <p className="mt-6 caption-mono text-muted">
                {switchPrompt}{' '}
                <Link href={switchHref} className="link-underline text-fg">
                  {switchLabel}
                </Link>
              </p>
            )}
          </div>
        </div>

        {/* Footer - legal */}
        <div className="border-t border-border px-5 py-4 sm:px-8 lg:px-10 xl:px-16 2xl:px-24">
          <div className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
            <p className="caption-mono text-dim">
              By continuing, you agree to our{' '}
              <a href="#" className="link-underline text-muted">
                Terms
              </a>{' '}
              and{' '}
              <a href="#" className="link-underline text-muted">
                Privacy Policy
              </a>
              .
            </p>
            <p className="caption-mono-upper text-dim">
              © {year} Velocity
            </p>
          </div>
        </div>
      </main>
    </div>
  );
}

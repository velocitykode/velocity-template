import { useAppearance } from '@/hooks/use-appearance';

interface ThemeToggleProps {
  /** Render on a dark brand panel (uses light tokens) */
  onDarkSurface?: boolean;
  className?: string;
}

export function ThemeToggle({ onDarkSurface = false, className = '' }: ThemeToggleProps) {
  const { appearance, updateAppearance } = useAppearance();
  const effective: 'light' | 'dark' =
    appearance === 'dark'
      ? 'dark'
      : appearance === 'light'
        ? 'light'
        : typeof window !== 'undefined' &&
            window.matchMedia('(prefers-color-scheme: dark)').matches
          ? 'dark'
          : 'light';

  const tokens = onDarkSurface
    ? {
        border: 'border-panel-border',
        muted: 'text-panel-dim',
        active: 'text-panel-fg',
        sep: 'text-panel-dim',
      }
    : {
        border: 'border-border',
        muted: 'text-dim',
        active: 'text-fg',
        sep: 'text-dim',
      };

  return (
    <button
      type="button"
      onClick={() => updateAppearance(effective === 'dark' ? 'light' : 'dark')}
      aria-label={`Switch to ${effective === 'dark' ? 'light' : 'dark'} theme`}
      aria-pressed={effective === 'light'}
      className={`inline-flex h-9 items-center gap-1.5 border ${tokens.border} px-3 font-mono text-[11px] uppercase tracking-widest rounded-none transition-colors ${className}`}
    >
      <span className={tokens.sep}>[</span>
      <span className={effective === 'dark' ? tokens.active : tokens.muted}>dark</span>
      <span className={tokens.sep}>|</span>
      <span className={effective === 'light' ? tokens.active : tokens.muted}>light</span>
      <span className={tokens.sep}>]</span>
    </button>
  );
}

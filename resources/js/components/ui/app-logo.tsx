import { Link } from '@inertiajs/react';

interface AppLogoProps {
  /** icon = compact (v_), full = wordmark (velocity_), responsive = icon mobile / full lg+ */
  variant?: 'icon' | 'full' | 'responsive';
  /** Size - drives font-size */
  size?: 'sm' | 'md' | 'lg' | 'xl';
  /** Link to homepage */
  href?: string;
  /** Additional classes */
  className?: string;
  /** Force panel-fg coloring (use inside dark brand panels regardless of theme) */
  onDarkSurface?: boolean;
}

const sizeClasses = {
  sm: { icon: 'text-sm', full: 'text-base' },
  md: { icon: 'text-base', full: 'text-lg' },
  lg: { icon: 'text-lg', full: 'text-xl' },
  xl: { icon: 'text-xl', full: 'text-2xl' },
} as const;

export function AppLogo({
  variant = 'responsive',
  size = 'md',
  href = '/',
  className = '',
  onDarkSurface = false,
}: AppLogoProps) {
  const sizes = sizeClasses[size];
  const fg = onDarkSurface ? 'text-panel-fg' : 'text-fg';

  const mark = (isIcon: boolean) => (
    <span
      className={`font-mono font-semibold tracking-tight leading-none ${isIcon ? sizes.icon : sizes.full} ${fg}`}
    >
      {isIcon ? 'v' : 'velocity'}
      <span className="text-phosphor">_</span>
    </span>
  );

  const logoContent =
    variant === 'icon' ? (
      mark(true)
    ) : variant === 'full' ? (
      mark(false)
    ) : (
      <>
        <span className="lg:hidden">{mark(true)}</span>
        <span className="hidden lg:inline">{mark(false)}</span>
      </>
    );

  if (href) {
    return (
      <Link href={href} className={`inline-flex items-center ${className}`}>
        {logoContent}
      </Link>
    );
  }
  return <div className={`inline-flex items-center ${className}`}>{logoContent}</div>;
}

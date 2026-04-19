import { Link } from '@inertiajs/react';

interface AppLogoProps {
  /** Show icon only (mobile) or full logo (desktop) */
  variant?: 'icon' | 'full' | 'responsive';
  /** Logo height class */
  size?: 'sm' | 'md' | 'lg' | 'xl';
  /** Link to homepage */
  href?: string;
  /** Additional classes */
  className?: string;
  /** Force the dark-surface logo (use inside dark brand panels regardless of theme) */
  onDarkSurface?: boolean;
}

const sizeClasses = {
  sm: { icon: 'h-6 w-6', full: 'h-7' },
  md: { icon: 'h-8 w-8', full: 'h-10' },
  lg: { icon: 'h-9 w-9', full: 'h-12' },
  xl: { icon: 'h-10 w-10', full: 'h-16' },
} as const;

export function AppLogo({
  variant = 'responsive',
  size = 'md',
  href = '/',
  className = '',
  onDarkSurface = false,
}: AppLogoProps) {
  const sizes = sizeClasses[size];

  const logoContent = (
    <>
      {/* Icon - shown on mobile when responsive, always when variant='icon' */}
      <img
        src="/icon.png"
        alt="Velocity"
        className={`${sizes.icon} ${
          variant === 'responsive' ? 'lg:hidden' : variant === 'icon' ? '' : 'hidden'
        }`}
      />

      {variant !== 'icon' && (
        <>
          {onDarkSurface ? (
            <img
              src="/logo-dark.png"
              alt="Velocity"
              className={`${sizes.full} ${variant === 'responsive' ? 'hidden lg:block' : ''}`}
            />
          ) : (
            <>
              <img
                src="/logo-light.png"
                alt="Velocity"
                className={`${sizes.full} ${
                  variant === 'responsive' ? 'hidden lg:block dark:hidden' : 'dark:hidden'
                }`}
              />
              <img
                src="/logo-dark.png"
                alt="Velocity"
                className={`${sizes.full} ${
                  variant === 'responsive' ? 'hidden dark:lg:block' : 'hidden dark:block'
                }`}
              />
            </>
          )}
        </>
      )}
    </>
  );

  if (href) {
    return (
      <Link href={href} className={`flex items-center ${className}`}>
        {logoContent}
      </Link>
    );
  }

  return <div className={`flex items-center ${className}`}>{logoContent}</div>;
}

import { type LucideIcon } from 'lucide-react';
import { forwardRef, type ButtonHTMLAttributes, type ReactNode } from 'react';

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  /** Visual variant */
  variant?: 'primary' | 'secondary' | 'ghost' | 'danger';
  /** Button size */
  size?: 'sm' | 'md' | 'lg';
  /** Icon before text */
  iconLeft?: LucideIcon;
  /** Icon after text */
  iconRight?: LucideIcon;
  /** Loading state */
  loading?: boolean;
  /** Full width */
  fullWidth?: boolean;
  children: ReactNode;
}

const variantClasses = {
  primary:
    'bg-accent text-accent-fg hover:opacity-90',
  secondary:
    'border border-border-strong bg-transparent text-fg hover:bg-surface-alt',
  ghost:
    'text-muted hover:text-fg hover:bg-surface-alt',
  danger:
    'bg-error text-white hover:opacity-90',
} as const;

const sizeClasses = {
  sm: 'h-8 px-3.5 text-xs gap-1.5',
  md: 'h-10 px-4 text-sm gap-2',
  lg: 'h-12 px-6 text-[14px] gap-2',
} as const;

const iconSizeClasses = {
  sm: 'h-3.5 w-3.5',
  md: 'h-4 w-4',
  lg: 'h-5 w-5',
} as const;

export const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  (
    {
      variant = 'primary',
      size = 'md',
      iconLeft: IconLeft,
      iconRight: IconRight,
      loading = false,
      fullWidth = false,
      disabled,
      className = '',
      children,
      ...props
    },
    ref,
  ) => {
    return (
      <button
        ref={ref}
        disabled={disabled || loading}
        className={`inline-flex items-center justify-center rounded-none font-medium tracking-wide transition-[background-color,color,opacity] duration-150 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent focus-visible:ring-offset-2 focus-visible:ring-offset-bg disabled:cursor-not-allowed disabled:opacity-50 ${variantClasses[variant]} ${sizeClasses[size]} ${fullWidth ? 'w-full' : ''} ${className}`}
        {...props}
      >
        {IconLeft && <IconLeft className={iconSizeClasses[size]} />}
        {children}
        {(loading || IconRight) && (
          <span
            className={`inline-flex items-center justify-center ${iconSizeClasses[size]}`}
          >
            {loading ? (
              <span className="ascii-spinner" aria-hidden />
            ) : IconRight ? (
              <IconRight className={iconSizeClasses[size]} />
            ) : null}
          </span>
        )}
      </button>
    );
  },
);

Button.displayName = 'Button';

import { type ReactNode } from 'react';

interface BadgeProps {
  children: ReactNode;
  /** Visual variant */
  variant?: 'default' | 'brand' | 'success' | 'warning' | 'danger' | 'outline';
  /** Badge size */
  size?: 'sm' | 'md';
  /** Optional dot indicator */
  dot?: boolean;
  /** Dot color (when dot=true) */
  dotColor?: 'brand' | 'success' | 'danger';
  /** Additional classes */
  className?: string;
}

const variantClasses = {
  default:
    'bg-surface-alt text-muted',
  brand:
    'bg-accent text-accent-fg',
  success:
    'border border-emerald-600/40 bg-emerald-600/10 text-emerald-700 dark:text-emerald-400',
  warning:
    'border border-amber-600/40 bg-amber-600/10 text-amber-700 dark:text-amber-400',
  danger:
    'border border-error/40 bg-error/10 text-error',
  outline:
    'border border-border-strong bg-transparent text-fg',
} as const;

const sizeClasses = {
  sm: 'px-2 py-0.5 text-[10px]',
  md: 'px-2.5 py-1 text-[11px]',
} as const;

const dotColorClasses = {
  brand: 'bg-accent',
  success: 'bg-emerald-500',
  danger: 'bg-error',
} as const;

export function Badge({
  children,
  variant = 'default',
  size = 'sm',
  dot = false,
  dotColor = 'brand',
  className = '',
}: BadgeProps) {
  return (
    <span
      className={`inline-flex items-center gap-1.5 rounded-none font-mono uppercase tracking-widest ${variantClasses[variant]} ${sizeClasses[size]} ${className}`}
    >
      {dot && <span className={`h-1.5 w-1.5 rounded-full ${dotColorClasses[dotColor]}`} />}
      {children}
    </span>
  );
}

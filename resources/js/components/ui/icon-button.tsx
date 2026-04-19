import { type LucideIcon } from 'lucide-react';
import { type ComponentProps, forwardRef } from 'react';

type ButtonOrLinkProps =
  | (ComponentProps<'button'> & { as?: 'button'; href?: never })
  | (ComponentProps<'a'> & { as: 'a'; href: string });

interface IconButtonBaseProps {
  /** Lucide icon component */
  icon: LucideIcon;
  /** Button size */
  size?: 'sm' | 'md' | 'lg';
  /** Visual variant */
  variant?: 'default' | 'danger' | 'ghost';
  /** Accessible label */
  'aria-label': string;
}

type IconButtonProps = IconButtonBaseProps & ButtonOrLinkProps;

const sizeClasses = {
  sm: 'h-8 w-8',
  md: 'h-10 w-10',
  lg: 'h-11 w-11',
} as const;

const iconSizeClasses = {
  sm: 'h-3.5 w-3.5',
  md: 'h-4 w-4',
  lg: 'h-5 w-5',
} as const;

const variantClasses = {
  default:
    'border border-border bg-transparent text-muted hover:border-fg hover:text-fg',
  danger:
    'border border-border bg-transparent text-muted hover:border-error hover:text-error',
  ghost:
    'bg-transparent text-muted hover:text-fg',
} as const;

export const IconButton = forwardRef<HTMLButtonElement | HTMLAnchorElement, IconButtonProps>(
  ({ icon: Icon, size = 'md', variant = 'default', className = '', as, ...props }, ref) => {
    const classes = `flex items-center justify-center rounded-none transition-colors duration-150 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent focus-visible:ring-offset-2 focus-visible:ring-offset-bg ${sizeClasses[size]} ${variantClasses[variant]} ${className}`;

    if (as === 'a') {
      const { href, ...anchorProps } = props as ComponentProps<'a'> & { href: string };
      return (
        <a ref={ref as React.Ref<HTMLAnchorElement>} href={href} className={classes} {...anchorProps}>
          <Icon className={iconSizeClasses[size]} />
        </a>
      );
    }

    return (
      <button ref={ref as React.Ref<HTMLButtonElement>} className={classes} {...(props as ComponentProps<'button'>)}>
        <Icon className={iconSizeClasses[size]} />
      </button>
    );
  },
);

IconButton.displayName = 'IconButton';

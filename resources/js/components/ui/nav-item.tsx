import { Link } from '@inertiajs/react';
import { ExternalLink, type LucideIcon } from 'lucide-react';

interface NavItemProps {
  /** Display name */
  name: string;
  /** Navigation URL */
  href: string;
  /** Icon component */
  icon: LucideIcon;
  /** Is this item active? */
  active?: boolean;
  /** Is this an external link? */
  external?: boolean;
  /** Click handler */
  onClick?: () => void;
  /** Additional classes */
  className?: string;
}

export function NavItem({
  name,
  href,
  icon: Icon,
  active = false,
  external = false,
  onClick,
  className = '',
}: NavItemProps) {
  const baseClasses =
    'group flex items-center gap-3 rounded-none px-3 py-2.5 text-sm transition-colors duration-150 focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-accent focus-visible:ring-offset-2 focus-visible:ring-offset-bg';
  const activeClasses = active
    ? 'border-l-2 border-accent pl-[10px] text-fg font-medium'
    : 'border-l-2 border-transparent pl-[10px] text-muted hover:text-fg';

  if (external) {
    return (
      <a
        href={href}
        target="_blank"
        rel="noopener noreferrer"
        className={`${baseClasses} ${activeClasses} ${className}`}
        onClick={onClick}
      >
        <Icon className="h-4 w-4" />
        <span className="flex-1 truncate">{name}</span>
        <ExternalLink className="h-3.5 w-3.5 opacity-50" />
      </a>
    );
  }

  return (
    <Link
      href={href}
      className={`${baseClasses} ${activeClasses} ${className}`}
      onClick={onClick}
    >
      <Icon className="h-4 w-4" />
      <span className="flex-1 truncate">{name}</span>
    </Link>
  );
}

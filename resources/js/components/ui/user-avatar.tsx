interface UserAvatarProps {
  /** User's full name */
  name?: string;
  /** Avatar size */
  size?: 'sm' | 'md' | 'lg';
  /** Image URL (optional) */
  src?: string;
  /** Additional classes */
  className?: string;
}

const sizeClasses = {
  sm: 'h-8 w-8 text-[11px]',
  md: 'h-10 w-10 text-xs',
  lg: 'h-12 w-12 text-sm',
} as const;

function getInitials(name?: string): string {
  if (!name) return 'U';
  const parts = name.trim().split(' ');
  if (parts.length === 1) return parts[0].charAt(0).toUpperCase();
  return (parts[0].charAt(0) + parts[parts.length - 1].charAt(0)).toUpperCase();
}

export function UserAvatar({ name, size = 'md', src, className = '' }: UserAvatarProps) {
  const initials = getInitials(name);

  if (src) {
    return (
      <img
        src={src}
        alt={name || 'User avatar'}
        className={`rounded-none object-cover ${sizeClasses[size]} ${className}`}
      />
    );
  }

  return (
    <div
      className={`flex items-center justify-center border border-border-strong bg-surface font-mono tracking-widest text-fg ${sizeClasses[size]} ${className}`}
    >
      {initials}
    </div>
  );
}

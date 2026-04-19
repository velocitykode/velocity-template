import { type LucideIcon } from 'lucide-react';

type ColorVariant = 'blue' | 'violet' | 'emerald' | 'amber' | 'rose';

interface StatCardProps {
  /** Stat label */
  label: string;
  /** Stat value */
  value: string | number;
  /** Icon component */
  icon: LucideIcon;
  /** Color theme (kept for API compatibility; rendered monochrome) */
  color?: ColorVariant;
  /** Optional change indicator (e.g., "+2") */
  change?: string;
  /** Optional status text */
  status?: string;
  /** Additional classes */
  className?: string;
}

export function StatCard({
  label,
  value,
  icon: Icon,
  change,
  className = '',
}: StatCardProps) {
  return (
    <div
      className={`border border-border bg-surface p-4 lg:p-5 transition-colors duration-200 hover:border-border-strong ${className}`}
    >
      <div className="flex items-start justify-between">
        <span className="caption-mono-upper text-dim">
          {label}
        </span>
        <Icon className="h-4 w-4 text-dim" />
      </div>
      <div className="mt-4 lg:mt-5 flex items-baseline gap-2">
        <span className="text-2xl lg:text-3xl xl:text-[32px] font-semibold tracking-tight text-fg">
          {value}
        </span>
        {change && (
          <span className="caption-mono text-muted">
            {change}
          </span>
        )}
      </div>
    </div>
  );
}

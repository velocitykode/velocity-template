import { forwardRef, type InputHTMLAttributes } from 'react';

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  /** Error state */
  error?: boolean;
}

export const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ className = '', error = false, ...props }, ref) => {
    return (
      <input
        ref={ref}
        className={`w-full rounded-none border bg-surface px-3.5 text-[14px] text-fg placeholder:text-dim transition-colors duration-150 focus:outline-none disabled:cursor-not-allowed disabled:opacity-50 h-11 lg:h-12 ${
          error
            ? 'border-error text-error placeholder:text-error focus:border-error'
            : 'border-border focus:border-fg'
        } ${className}`}
        {...props}
      />
    );
  },
);

Input.displayName = 'Input';

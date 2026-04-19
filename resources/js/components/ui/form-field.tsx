import { type ReactNode } from 'react';

interface FormFieldProps {
  /** Field label (rendered as floating caption above the input) */
  label?: string;
  /** Error message */
  error?: string;
  /** Helper text */
  hint?: string;
  /** Required indicator */
  required?: boolean;
  /** HTML for attribute */
  htmlFor?: string;
  /** Form field input */
  children: ReactNode;
  /** Additional classes */
  className?: string;
}

export function FormField({
  label,
  error,
  hint,
  required = false,
  htmlFor,
  children,
  className = '',
}: FormFieldProps) {
  return (
    <div className={`space-y-1.5 ${className}`}>
      {label && (
        <label
          htmlFor={htmlFor}
          className="caption-mono-upper block text-muted"
        >
          {label}
          {required && <span className="ml-1 text-error">*</span>}
        </label>
      )}
      {children}
      {error && (
        <p className="caption-mono text-error">
          {error}
        </p>
      )}
      {hint && !error && (
        <p className="caption-mono text-dim">
          {hint}
        </p>
      )}
    </div>
  );
}

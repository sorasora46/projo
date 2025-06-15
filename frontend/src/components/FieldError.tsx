import type { FC } from "react";

interface FieldErrorProps {
  message?: string;
}

const FieldError: FC<FieldErrorProps> = ({ message }) => {
  return (
    <div className="flex-grow">
      <span className="text-error text-sm block">{message ?? ""}</span>
    </div>
  );
};

export default FieldError;

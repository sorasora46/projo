const FieldError = ({ message }: { message?: string }) => (
  <span
    className={`block text-error text-sm h-4 ${message ? "" : "invisible"}`}
  >
    {message}
  </span>
);

export default FieldError;

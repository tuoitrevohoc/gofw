interface GraphQLError extends Error {
  message: string;
  source: {
    errors?: Array<{ message: string }>;
  };
}

export default function getErrorMessage(err: Error) {
  const error = err as unknown as GraphQLError;

  return (
    (error.source && error.source.errors && error.source.errors[0]?.message) ||
    error.message ||
    "An unexpected error occurred"
  );

  return "An unexpected error occurred";
}

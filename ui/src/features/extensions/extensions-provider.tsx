import { registerExtensions } from '@ui/extensions';

export const ExtensionsProvider = ({ children }: { children: React.ReactNode }) => {
  registerExtensions();
  return children;
};

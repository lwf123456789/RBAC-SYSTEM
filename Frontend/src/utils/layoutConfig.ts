export const routesWithoutDefaultLayout = ['/', '/not-found'];

export const shouldUseDefaultLayout = (pathname: string) => {
  return !routesWithoutDefaultLayout.includes(pathname);
};
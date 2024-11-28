export default defineNuxtRouteMiddleware((to) => {
  if (process.server) return; // Ensure it only runs on the client side
  
  const token = localStorage.getItem('token');

  const publicPaths = ['/login'];

  if (!token && !publicPaths.includes(to.path)) {
    return navigateTo('/login');
  }
});
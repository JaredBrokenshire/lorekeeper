export default [
  {
    path: '/login',
    name: 'login-page',
    component: () => import('@/views/auth/Login.vue'),
    meta: { guestOnly: true }
  },
  {
    path: '/register',
    name: 'register-page',
    component: () => import('@/views/auth/Register.vue'),
    meta: { guestOnly: true }
  },
]

export default [
  {
    path: "/characters",
    name: "character-list",
    component: () => import("@/views/characters/CharacterList.vue"),
    meta: {
      requiresAuth: true,
      showHeader: true,
    }
  }
]
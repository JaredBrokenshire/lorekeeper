export default [
  {
    path: "/campaigns",
    name: "campaigns-list",
    component: () => import("@/views/campaigns/CampaignsList.vue"),
    meta: {
      requiresAuth: true,
      showHeader: true,
    }
  }
]
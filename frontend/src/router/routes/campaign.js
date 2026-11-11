import backgroundTexture from "@/assets/images/pages/background-texture.png"

export default [
  {
    path: "/campaigns",
    name: "campaigns-list",
    component: () => import("@/views/campaigns/CampaignsList.vue"),
    meta: {
      requiresAuth: true,
      showHeader: true,
    }
  },
  {
    path: "/campaigns/:id",
    name: "campaign-details",
    component: () => import("@/views/campaigns/CampaignDetails.vue"),
    meta: {
      requiresAuth: true,
      showHeader: true,
      backgroundImage: backgroundTexture,
    }
  }
]
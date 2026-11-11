<template>
  <div v-if="campaign && !loading">
    <h1 class="w-full text-ink border-b-2 border-ember py-4 mb-4">
      {{ campaign.name }}
    </h1>

    <div class="w-full mb-4">
      <h3 class="text-ink">{{ campaign.dm.username }}</h3>
      <p class="text-ink-soft">Dungeon Master</p>
    </div>

    <!-- Tabs -->
    <div class="flex w-full border-b border-ember mb-6">
      <button
        :class="[
          'px-6 py-3 font-semibold transition-colors',
          activeTab === 'lorekeeper'
            ? 'text-ink border-b-2 border-ember'
            : 'text-ink-light hover:text-ink'
        ]"
        @click="activeTab = 'lorekeeper'"
      >
        Lorekeeper
      </button>

      <button
        :class="[
          'px-6 py-3 font-semibold transition-colors',
          activeTab === 'characters'
            ? 'text-ink border-b-2 border-ember'
            : 'text-ink-light hover:text-ink'
        ]"
        @click="activeTab = 'characters'"
      >
        Characters
      </button>
    </div>

    <!-- Lorekeeper Tab -->
    <div v-if="activeTab === 'lorekeeper'">
      <!-- Lorekeeper content will go here -->
    </div>

    <!-- Characters Tab -->
    <div v-if="activeTab === 'characters'">
      <h3 class="text-ink mb-4">
        Active Characters ({{ campaign.characters.length }})
      </h3>

      <div class="h-fit w-full grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
        <character-card
          v-for="character in campaign.characters"
          :key="`campaign-${campaign.id}-details-character-card-${character.id}`"
          :character="character"
        />
      </div>
    </div>
  </div>
</template>

<script>
  import HelperService from "@/services/helper.service";
  import CampaignService from "@/services/campaign.service";
  import CharacterCard from "@/components/characters/CharacterCard.vue";

  export default {
    name: "CampaignDetails",
    components: {
      CharacterCard
    },
    data() {
      return {
        loading: true,
        campaign: null,
        activeTab: "lorekeeper"
      }
    },
    mounted() {
      this.getCampaign()
    },

    methods: {
      async getCampaign() {
        this.loading = true;

        try {
          const res = await CampaignService.get(this.$route.params.id);
          this.campaign = res.data;
        } catch (error) {
          let errorText = "Could not get campaign, please refresh and try again";

          if (error.response && error.response.data.error) {
            errorText = error.response.data.error;
          }

          HelperService.errorToast(this.$toast, errorText);
        } finally {
          this.loading = false;
        }
      }
    }
  }
</script>

<style scoped>
</style>

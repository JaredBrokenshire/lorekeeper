<template>
  <div v-if="campaign && !loading">
    <h1 class="w-full text-ink border-b-2 border-ember py-4 mb-4">
      {{ campaign.name }}
    </h1>

    <div class="w-full flex-between mb-4">
      <div class="flex-grow">
        <h3 class="text-ink">{{ campaign.dm.username }}</h3>
        <p class="text-ink-soft">Dungeon Master</p>
      </div>

      <h5>
        <button
          class="px-4 py-2 bg-gold text-ink rounded transition-all hover:bg-gold-light shadow-transparent shadow hover:shadow-ink-light"
          @click="showUploadSessionNotesModal = true"
        >
          Upload Session Notes
        </button>
      </h5>
    </div>

    <!-- Tabs -->
    <div class="flex w-full border-b border-ember mb-6">
      <button
        v-for="tab in tabs"
        :key="`campaign-${campaign.id}-details-tab-${tab}`"
        :class="[
          'px-6 py-3 font-semibold transition-colors',
          activeTab === tab
            ? 'text-ink border-b-2 border-ember'
            : 'text-ink-light hover:text-ink'
        ]"
        @click="activeTab = tab"
      >
        {{ tab }}
      </button>
    </div>

    <!-- Lorekeeper Tab -->
    <div v-if="activeTab === 'Lorekeeper'">
      <!-- Lorekeeper content will go here -->
    </div>

    <!-- Characters Tab -->
    <div v-if="activeTab === 'Characters'">
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

    <!-- Session Notes Tab -->
    <div v-if="activeTab === 'Session Notes'">
      <table v-if="campaign.session_notes">
        <tr class="hover:bg-transparent">
          <th class="w-[35%]">Name</th>
          <th class="w-[35%]">Author</th>
          <th class="w-[15%]">Date</th>
          <th class="w-[10%]" />
        </tr>
        <tr
          v-for="(file, index) in campaign.session_notes"
          :key="`campaign-${campaign.id}-session-note-${index}`"
          class="group"
        >
          <td>{{ file.name }}</td>
          <td>{{ file.author?.username }}</td>
          <td>{{ new Date(file.created_at).toLocaleDateString() }}</td>
          <td class="cursor-pointer">
            <router-link
              :to="`/campaigns/${campaign.id}/notes/${file.id}`"
              class="text-ember text-center transition-all font-normal group-hover:font-bold"
            >
              <h4>View</h4>
            </router-link>
          </td>
        </tr>
      </table>

      <div
        v-else
        class="relative h-fit flex justify-center md:items-center my-4 md:my-8 w-full h-fit"
      >
        <card
          :animated="false"
          class="w-[90%] xs:w-[75%] sm:w-[60%] lg:w-[40%] h-[80%]"
        >
          <template #content>
            <div class="min-w-full min-h-full relative flex flex-col-reverse justify-between">
              <div
                class="flex flex-col justify-between items-start lg:justify-end gap-4"
              >
                <div class="sm:mx-2 xl:mx-8">
                  <div class="mb-4">
                    <h2 class="hidden sm:block text-gold uppercase">No Session Notes</h2>
                    <h4 class="block sm:hidden text-gold uppercase">No Session Notes</h4>
                  </div>
                  <p class="text-gold-light">
                    There are no session notes for this campaign yet. Upload some now.
                  </p>
                </div>
              </div>

              <div class="min-w-[125px] min-h-[125px] hidden 2xs:flex justify-center items-center relative mx-auto w-[50%] lg:w-full 2xs:mb-8 lg:mb-16">
                <img
                  alt=""
                  class="min-w-[50%] min-h-[50%] xs:max-w-[150%] xs:max-h-[150%] lg:max-w-[200%] lg:max-h-[200%] aspect-square object-contain absolute -top-[2rem] xs:-top-[3rem] sm:-top-[4rem] md:-top-[5rem] transition-all group-hover:scale-110"
                  src="@/assets/images/pages/campaigns/session-notes.png"
                >
              </div>
            </div>
          </template>
        </card>
      </div>
    </div>
    
    <modal
      id="modal-upload-session-notes"
      :visible="showUploadSessionNotesModal"
      title="Upload Session Notes"
      @close="closeModals"
    >
      <upload-session-notes-modal
        @close="closeModals"
        @upload="uploadSessionNotes"
      />
    </modal>
  </div>
</template>

<script>
  import Card from "@/components/Card.vue";
  import Modal from "@/components/Modal.vue";
  import HelperService from "@/services/helper.service";
  import CampaignService from "@/services/campaign.service";
  import CharacterCard from "@/components/characters/CharacterCard.vue";
  import UploadSessionNotesModal from "@/views/campaigns/sections/UploadSessionNotesModal.vue";

  export default {
    name: "CampaignDetails",
    components: {
      Modal,
      UploadSessionNotesModal,
      Card,
      CharacterCard
    },
    data() {
      return {
        loading: true,
        campaign: null,
        activeTab: "Lorekeeper",
        tabs: ["Lorekeeper", "Characters", "Session Notes"],
        showUploadSessionNotesModal: false,
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
      },
      async uploadSessionNotes(file) {
        this.closeModals()
        this.loading = true

        try {
          const res = await CampaignService.uploadSessionNotes(this.campaign.id, file);
          this.campaign = res.data
        } catch (error) {
          let errorText = "Could not get upload session notes, please refresh and try again";

          if (error.response && error.response.data.error) {
            errorText = error.response.data.error;
          }

          HelperService.errorToast(this.$toast, errorText);
        } finally {
          this.loading = false;
        }
      },
      closeModals() {
        this.showUploadSessionNotesModal = false;
      }
    }
  }
</script>

<style scoped>
</style>

<template>
  <div class="z-20 relative min-h-full flex-grow">
    <div class="flex-between">
      <h1 class="text-gold mb-4">Campaigns</h1>
    </div>

    <div class="w-full md:max-w-[50%] xl:max-w-[30%] mb-8">
      <input-wrapper
        v-model="filters.search"
        label="Search"
        name="classes-search"
        placeholder="Search"
        @input="onSearchInput"
      />
    </div>

    <div
      v-if="!loading && totalCount > 0"
      class="w-full grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-3  gap-8 md:gap-16"
    >
      <campaign-card
        v-for="campaign in campaigns"
        :key="`campaign-card-${campaign.id}`"
        :campaign="campaign"
      />
    </div>

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
                  <h2 class="hidden sm:block text-gold uppercase">No Campaigns</h2>
                  <h4 class="block sm:hidden text-gold uppercase">No Campaigns</h4>
                </div>
                <p class="text-gold-light">
                  No campaigns match your filters, try a different search.
                </p>
              </div>
            </div>

            <div class="min-w-[125px] min-h-[125px] hidden 2xs:flex justify-center items-center relative mx-auto w-[50%] lg:w-full 2xs:mb-8 lg:mb-16">
              <img
                alt=""
                class="min-w-[50%] min-h-[50%] xs:max-w-[150%] xs:max-h-[150%] lg:max-w-[200%] lg:max-h-[200%] aspect-square object-contain absolute -top-[2rem] xs:-top-[3rem] sm:-top-[4rem] md:-top-[5rem] transition-all group-hover:scale-110"
                src="@/assets/images/pages/home/campaigns.png"
              >
            </div>
          </div>
        </template>
      </card>
    </div>
  </div>
</template>

<script>
import debounce from "lodash/debounce";
import Card from "@/components/Card.vue";
import HelperService from "@/services/helper.service";
import CampaignService from "@/services/campaign.service";
import InputWrapper from "@/components/forms/InputWrapper.vue"
import CampaignCard from "@/views/campaigns/sections/CampaignCard.vue";

export default {
    name: "CampaignsList",
    components: {Card, InputWrapper, CampaignCard},
    data() {
      return {
        loading: true,
        campaigns: [],
        filters: {
          page: 0,
          page_size: 24,
          search: "",
        },
        totalCount: 0,
        debouncedSearch: null,
      }
    },
    mounted() {
      this.debouncedSearch = debounce(this.getMoreCampaigns, 500);
      this.getCampaigns()
    },
    methods: {
      getMoreCampaigns(page) {
        this.filters.page = page;
        this.getCampaigns()
      },
      async getCampaigns() {
        this.loading = true;

        try {
          const res = await CampaignService.list(this.filters);
          this.campaigns = res.data.data || [];
          this.filters.page = res.data.meta.page;
          this.filters.page_size = res.data.meta.page_size;
          this.totalCount = res.data.meta.total_count
        } catch (error) {
          let errorText = "Could not get campaigns, please refresh and try again";

          if (res && res.data.error) {
            errorText = res.data.error;
          }

          HelperService.errorToast(this.$toast, errorText);
        } finally {
          this.loading = false;
        }
      },
      onSearchInput() {
        this.filters.page = 0;
        this.debouncedSearch(0)
      }
    }
  }
</script>

<style scoped>

</style>
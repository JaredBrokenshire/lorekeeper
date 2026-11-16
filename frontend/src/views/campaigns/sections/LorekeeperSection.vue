<template>
  <div class="w-full h-full">
    <div v-if="!loading">
      <div class="w-full flex justify-between items-end md:items-center mb-4 gap-4">
        <div class="flex-grow h-full md:max-w-[50%]">
          <input-wrapper
            v-model="question"
            label="Search"
            label-styles="text-ink"
            name="campaigns-notes-search"
            placeholder="Ask a question about the adventures of your party..."
          />
        </div>

        <h4>
          <button
            class="px-4 py-2 md:px-6 md:py-3 rounded bg-gold transition-all hover:bg-gold-light text-ink shadow-lg shadow-transparent hover:shadow-gold"
            @click="searchNotes"
          >
            Search
          </button>
        </h4>
      </div>

      <div v-if="questionsAsked.length > 0 && answersReceived.length > 0">
        <div class="flex flex-col-reverse gap-4 w-full">
          <div
            v-for="(questionAsked, index) in questionsAsked"
            :key="`question-${index}`"
            class="w-full flex flex-col gap-4"
          >
            <p class="max-w-[60%] bg-gold-light p-4 rounded-xl">{{ answersReceived[index] }}</p>
            <p class="max-w-[60%] w-full self-end bg-gold-light p-4 rounded-xl">{{ questionAsked }}</p>
          </div>
        </div>
      </div>

      <h3
        v-else
        class="w-fit mx-auto bg-gold-light p-8 rounded-xl mt-8"
      >
        Ask The Lorekeeper for information
      </h3>


    </div>

    <div
      v-else
      class="relative max-w-[50%] mx-auto mt-8"
    >
      <card
        :animated="false"
        class="flex-center w-[50%]"
      >
        <template #content>
          <div class="min-w-full min-h-full relative flex flex-col-reverse xs:flex-row justify-between lg:flex-col-reverse">
            <div
              class="w-full h-full xs:w-[50%] lg:w-full lg:h-[30%] flex flex-col justify-between items-start lg:justify-end gap-4"
            >
              <div class="mx-8 mt-4">
                <div class="">
                  <h2 class="hidden sm:block text-gold uppercase">Pondering the Scrolls</h2>
                  <h4 class="block sm:hidden text-gold uppercase">Pondering the Scrolls</h4>
                </div>
                <p class="text-gold-light">We are searching the tomes for the answers you seek.</p>
              </div>

              <img
                :src="d20"
                alt=""
                class="mx-auto w-32 aspect-square transition-all animate-spin"
              >
            </div>

            <div class="min-w-[125px] min-h-[125px] hidden xs:flex justify-center items-center relative mx-auto w-[50%] lg:w-full lg:mb-12">
              <img
                :alt="`Scrolls Image`"
                :src="scroll"
                class="min-w-[50%] min-h-[50%] max-w-full max-h-full xs:max-w-[150%] xs:max-h-[150%] lg:max-w-[200%] lg:max-h-[200%] aspect-square object-contain absolute -top-[2rem] xs:-top-[3rem] sm:-top-[4rem] md:-top-[5rem] transition-all group-hover:scale-110"
              >
            </div>
          </div>
        </template>
      </card>
    </div>
  </div>
</template>

<script>
  import Card from "@/components/Card.vue";
  import HelperService from "@/services/helper.service";
  import D20 from "@/assets/images/pages/campaigns/d20.png"
  import CampaignService from "@/services/campaign.service";
  import InputWrapper from "@/components/forms/InputWrapper.vue";
  import Scroll from "@/assets/images/pages/campaigns/scroll.png";

  export default {
    name: "LorekeeperSection",
    components: {Card, InputWrapper},
    props: {
      campaignId: {
        type: Number,
        required: true,
      }
    },
    data() {
      return {
        loading: false,
        question: "",
        answer: null,
        searchTitles: [
          "The Oracle has spoken: ",
          "The prophecy has been fulfilled: ",
          "The scrolls bear the knowledge: ",
          "The Lorekeeper wishes to share his knowledge: ",
        ],
        randomIndex: 0,
        d20: D20,
        scroll: Scroll,
        questionsAsked: [],
        answersReceived: [],
      }
    },
    methods: {
      async searchNotes() {
        if (this.question === "") return
        this.loading = true;
        this.answer = null;
        this.randomIndex = Math.floor(Math.random() * this.searchTitles.length);

        this.questionsAsked.push(this.question);
        try {
          const dto = {
            question: this.question,
          }

          const res = await CampaignService.search(this.$props.campaignId, dto);
          this.answer = res.data;
          this.answersReceived.push(this.answer.answer);
        } catch (error) {
          this.questionsAsked.pop()
          let errorText = "Could not search session notes, please refresh and try again";

          if (error.response && error.response.data.error) {
            errorText = error.response.data.error;
          }

          HelperService.errorToast(this.$toast, error, errorText.toString());
        } finally {
          this.loading = false;
        }
      },
    }
  }
</script>

<style scoped>

</style>
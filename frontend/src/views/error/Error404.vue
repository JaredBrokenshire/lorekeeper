<template>
  <div class="relative flex justify-center m-4 sm:m-8 md:items-center h-full">
    <card
      :animated="false"
      class="w-[90%] xs:w-[75%] lg:w-[40%] h-[80%]"
    >
      <template #content>
        <div class="min-w-full min-h-full relative flex flex-col-reverse justify-between">
          <div
            class="flex flex-col justify-between items-start lg:justify-end gap-4"
          >
            <div class="sm:mx-2 xl:mx-8">
              <div class="mb-4">
                <h2 class="hidden sm:block text-gold uppercase">Page Not Found!</h2>
                <h4 class="block sm:hidden text-gold uppercase">Page Not Found!</h4>
              </div>
              <p class="text-gold-light">
                We live in a world of uncertainty. But certainly, the page you were looking for isn’t here.
                <span class="hidden 2xs:inline">
                  Perhaps this halfling has stolen it and hidden it in another place. Try searching for what you were looking for in another realm.
                </span>
              </p>
            </div>

            <router-link
              class="w-full flex-center"
              to="/"
            >
              <button
                class="mt-4 lg:mt-2 bg-ember rounded text-sm md:text-lg px-2 py-1 sm:px-4 sm:py-2 transition-all shadow-lg shadow-transparent text-gold-light hover:bg-ember-light hover:shadow-ink"
              >
                Return Home
              </button>
            </router-link>

          </div>

          <div class="min-w-[125px] min-h-[125px] hidden 2xs:flex justify-center items-center relative mx-auto w-[50%] lg:w-full 2xs:mb-8 lg:mb-16">
            <img
              alt="404 Image"
              class="min-w-[50%] min-h-[50%] xs:max-w-[150%] xs:max-h-[150%] lg:max-w-[200%] lg:max-h-[200%] aspect-square object-contain absolute -top-[2rem] xs:-top-[3rem] sm:-top-[4rem] md:-top-[5rem] transition-all group-hover:scale-110"
              src="@/assets/images/pages/error/404.png"
            >
          </div>
        </div>
      </template>
    </card>
  </div>
</template>

<script>
  import Card from "@/components/Card.vue";

  export default {
    name: "Error404",
    components: {Card}
  }
</script>

<style scoped>

</style>
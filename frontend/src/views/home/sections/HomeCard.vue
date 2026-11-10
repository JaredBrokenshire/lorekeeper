<template>
  <router-link
    :to="$props.route"
    class="flex-grow w-full h-full"
  >
    <card>
      <template #content>
        <div class="min-w-full min-h-full relative flex flex-col-reverse xs:flex-row justify-between lg:flex-col-reverse">
          <div
            class="w-full h-full xs:w-[50%] lg:w-full lg:h-[30%] flex flex-col justify-between items-start lg:justify-end gap-4"
          >
            <div>
              <div class="mb-4">
                <h2 class="hidden sm:block text-gold uppercase">{{ $props.title }}</h2>
                <h4 class="block sm:hidden text-gold uppercase">{{ $props.title }}</h4>
              </div>
              <p class="text-gold-light">{{ $props.bodyText }}</p>
            </div>

            <button
              class="mt-4 lg:mt-2 bg-ember rounded text-sm md:text-lg px-2 py-1 sm:px-4 sm:py-2 transition-all shadow-lg shadow-transparent text-gold-light group-hover:bg-ember-light group-hover:shadow-ink"
            >
              View {{ $props.title }}
            </button>

          </div>

          <div class="min-w-[125px] min-h-[125px] hidden xs:flex justify-center items-center relative mx-auto w-[50%] lg:w-full lg:mb-12">
            <img
              :alt="`${$props.title} Image`"
              :src="$props.imagePath"
              class="min-w-[50%] min-h-[50%] max-w-full max-h-full xs:max-w-[150%] xs:max-h-[150%] lg:max-w-[200%] lg:max-h-[200%] aspect-square object-contain absolute -top-[2rem] xs:-top-[3rem] sm:-top-[4rem] md:-top-[5rem] transition-all group-hover:scale-110"
            >
          </div>
        </div>
      </template>
    </card>
  </router-link>
</template>

<script>
  import Card from "@/components/Card.vue";

  export default {
    name: "HomeCard",
    components: {Card},
    props: {
      route: {
        type: String,
        default: "/"
      },
      title: {
        type: String,
        default: ""
      },
      bodyText: {
        type: String,
        default: ""
      },
      imagePath: {
        type: String,
        default: ""
      }
    }
  }
</script>

<style scoped>

</style>
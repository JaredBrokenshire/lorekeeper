<template>
  <transition name="fade">
    <div
      v-if="visible"
      class="fixed inset-0 z-50 flex justify-center sm:items-center p-8 sm:p-16 backdrop-blur-sm"
    >
      <div
        :class="maxWidth"
        class="flex-grow h-fit"
      >
        <card :animated="false">
          <template #content>
            <div class="lg:px-12 xl:px-2 relative">
              <div class="flex-between flex-row-reverse w-full">

                <button
                  v-if="$props.showClose"
                  class="w-8 aspect-square rounded transition-all group hover:bg-gray-400"
                  @click="$emit('close')"
                >
                  <span class="sr-only">Close</span>
                  <feather-icon
                    class="w-full h-full text-ember transition-all group-hover:text-ember-light"
                    icon="XIcon"
                  />
                </button>

                <h2
                  v-if="$props.title"
                  class="flex-grow text-gold"
                >
                  {{ $props.title }}
                </h2>
              </div>

              <div class="mt-8">
                <slot />
              </div>
            </div>
          </template>
        </card>
      </div>
    </div>
  </transition>
</template>

<script>
  import Card from "@/components/Card.vue";

  export default {
    name: "Modal",
    components: {Card},
    props: {
      visible: {
        type: Boolean,
        default: false,
      },
      title: {
        type: String,
        default: "",
      },
      showClose: {
        type: Boolean,
        default: true,
      },
      size: {
        type: String,
        default: "md",
      }
    },
    computed: {
      maxWidth() {
        switch (this.size) {
        case "sm":
          return "max-w-full md:max-w-[50%] xl:max-w-[25%]"
        case "md":
        default:
          return "max-w-full sm:max-w-[75%] xl:max-w-[35%]"
        case "lg":
          return "max-w-full md:max-w-[75%] xl:max-w-[50%]"
        case "xl":
          return "max-w-full"
        }
      }
    }
  };
</script>

<style scoped>
</style>

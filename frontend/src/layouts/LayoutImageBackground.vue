
<template>
  <div
    :style="backgroundStyle"
    class="min-h-screen bg-cover bg-center bg-fixed"
  >
    <slot />
  </div>
</template>

<script>
  import trees from "@/assets/images/pages/background.jpg"

  export default {
    name: "LayoutImageBackground",

    props: {
      backgroundImage: {
        type: String,
        default: null
      }
    },

    computed: {
      backgroundStyle() {
        return this.backgroundImage
          ? {
            backgroundImage: `url('${this.backgroundImage}')`
          }
          : {
            backgroundImage: `url('${trees}')`
          };
      }
    }
  }
</script>

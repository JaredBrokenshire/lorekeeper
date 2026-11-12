
<template>
  <component
    :is="layout"
    :background-image="backgroundImage"
  >
    <div class="relative min-h-screen max-w-screen flex flex-col">
      <page-header
        v-if="showHeader"
        class="w-screen shadow-lg shadow-ink"
      />

      <div class="w-full min-h-full flex-grow mx-auto flex justify-center p-4 md:p-8">
        <div class="max-w-[1340px] mx-auto relative flex-grow min-h-full">
          <div class="flex flex-grow w-full min-h-full">
            <router-view class="flex flex-col flex-grow" />
          </div>
        </div>
      </div>
    </div>
  </component>
</template>

<script>
  import PageHeader from "@/components/PageHeader.vue";

  const LayoutImageBackground = () => import("@/layouts/LayoutImageBackground.vue");

  export default {
    name: 'App',
    components: {
      PageHeader
    },
    data() {
      return {
        layout: LayoutImageBackground,
        showHeader: !!this.$route.meta.showHeader,
        backgroundImage: this.$route.meta.backgroundImage || null
      }
    },
    watch: {
      $route(to) {
        this.showHeader = !!to.meta.showHeader;
        this.backgroundImage = to.meta.backgroundImage || null;
      }
    }
  }
</script>
